package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/goravel/framework/contracts/testing/docker"
	"github.com/goravel/framework/support/process"
)

type Redis struct {
	port        int
	containerID string
	image       *docker.Image
}

func NewRedis() *Redis {
	return &Redis{
		image: &docker.Image{
			Repository:   "redis",
			Tag:          "latest",
			ExposedPorts: []string{"6379"},
		},
	}
}

func (receiver *Redis) Build() error {
	command, exposedPorts := ImageToCommand(receiver.image)
	containerID, err := process.Run(command)
	if err != nil {
		return fmt.Errorf("init Redis docker error: %v", err)
	}
	if containerID == "" {
		return fmt.Errorf("no container id return when creating Redis docker")
	}

	receiver.containerID = containerID
	receiver.port = ExposedPort(exposedPorts, 6379)

	if _, err := receiver.connect(); err != nil {
		return fmt.Errorf("connect Redis docker error: %v", err)
	}

	return nil
}

func (receiver *Redis) Config() RedisConfig {
	return RedisConfig{
		Port: receiver.port,
	}
}

func (receiver *Redis) Shutdown() error {
	if _, err := process.Run(fmt.Sprintf("docker stop %s", receiver.containerID)); err != nil {
		return fmt.Errorf("stop Redis docker error: %v", err)
	}

	return nil
}

func (receiver *Redis) connect() (*redis.Client, error) {
	var (
		client *redis.Client
		err    error
	)
	for i := 0; i < 60; i++ {
		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("localhost:%d", receiver.port),
			Password: "",
			DB:       0,
		})

		if _, err = client.Ping(context.Background()).Result(); err == nil {
			break
		}

		time.Sleep(2 * time.Second)
	}

	return client, err
}

type RedisConfig struct {
	Port int
}
