package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"gorm.io/gorm/clause"

	contractsauth "github.com/goravel/framework/contracts/auth"
	"github.com/goravel/framework/contracts/cache"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/support/carbon"
	"github.com/goravel/framework/support/database"
)

const ctxKey = "GoravelAuth"

type Claims struct {
	Key string `json:"key"`
	jwt.RegisteredClaims
}

type Guard struct {
	Claims *Claims
	Token  string
}

type Guards map[string]*Guard

type Auth struct {
	cache  cache.Cache
	config config.Config
	ctx    http.Context
	guard  string
	orm    orm.Orm
}

func NewAuth(guard string, cache cache.Cache, config config.Config, ctx http.Context, orm orm.Orm) *Auth {
	return &Auth{
		cache:  cache,
		config: config,
		ctx:    ctx,
		guard:  guard,
		orm:    orm,
	}
}

func (a *Auth) Guard(name string) contractsauth.Auth {
	return NewAuth(name, a.cache, a.config, a.ctx, a.orm)
}

// User need parse token first.
func (a *Auth) User(user any) error {
	auth, ok := a.ctx.Value(ctxKey).(Guards)
	if !ok || auth[a.guard] == nil {
		return errors.AuthParseTokenFirst
	}
	if auth[a.guard].Claims == nil {
		return errors.AuthParseTokenFirst
	}
	if auth[a.guard].Claims.Key == "" {
		return errors.AuthInvalidKey
	}
	if auth[a.guard].Token == "" {
		return errors.AuthTokenExpired
	}
	if err := a.orm.Query().FindOrFail(user, clause.Eq{Column: clause.PrimaryColumn, Value: auth[a.guard].Claims.Key}); err != nil {
		return err
	}

	return nil
}

func (a *Auth) ID() (string, error) {
	auth, ok := a.ctx.Value(ctxKey).(Guards)
	if !ok || auth[a.guard] == nil {
		return "", errors.AuthParseTokenFirst
	}
	if auth[a.guard].Token == "" {
		return "", errors.AuthTokenExpired
	}

	return auth[a.guard].Claims.Key, nil
}

func (a *Auth) Parse(token string) (*contractsauth.Payload, error) {
	token = strings.ReplaceAll(token, "Bearer ", "")
	if a.cache == nil {
		return nil, errors.CacheSupportRequired.SetModule(errors.ModuleAuth)
	}
	if a.tokenIsDisabled(token) {
		return nil, errors.AuthTokenDisabled
	}

	jwtSecret := a.config.GetString("jwt.secret")
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	}, jwt.WithTimeFunc(func() time.Time {
		return carbon.Now().StdTime()
	}))
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) && tokenClaims != nil {
			claims, ok := tokenClaims.Claims.(*Claims)
			if !ok {
				return nil, errors.AuthInvalidClaims
			}

			a.makeAuthContext(claims, "")

			return &contractsauth.Payload{
				Guard:    claims.Subject,
				Key:      claims.Key,
				ExpireAt: claims.ExpiresAt.Local(),
				IssuedAt: claims.IssuedAt.Local(),
			}, errors.AuthTokenExpired
		}

		return nil, errors.AuthInvalidToken
	}
	if tokenClaims == nil || !tokenClaims.Valid {
		return nil, errors.AuthInvalidToken
	}

	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok {
		return nil, errors.AuthInvalidClaims
	}

	a.makeAuthContext(claims, token)

	return &contractsauth.Payload{
		Guard:    claims.Subject,
		Key:      claims.Key,
		ExpireAt: claims.ExpiresAt.Time,
		IssuedAt: claims.IssuedAt.Time,
	}, nil
}

func (a *Auth) Login(user any) (token string, err error) {
	id := database.GetID(user)
	if id == nil {
		return "", errors.AuthNoPrimaryKeyField
	}

	return a.LoginUsingID(id)
}

func (a *Auth) LoginUsingID(id any) (token string, err error) {
	jwtSecret := a.config.GetString("jwt.secret")
	if jwtSecret == "" {
		return "", errors.AuthEmptySecret
	}

	nowTime := carbon.Now()
	expireTime := nowTime.AddMinutes(a.getTtl()).StdTime()
	key := cast.ToString(id)
	if key == "" {
		return "", errors.AuthInvalidKey
	}
	claims := Claims{
		key,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime.StdTime()),
			Subject:   a.guard,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	a.makeAuthContext(&claims, token)

	return
}

// Refresh need parse token first.
func (a *Auth) Refresh() (token string, err error) {
	auth, ok := a.ctx.Value(ctxKey).(Guards)
	if !ok || auth[a.guard] == nil {
		return "", errors.AuthParseTokenFirst
	}
	if auth[a.guard].Claims == nil {
		return "", errors.AuthParseTokenFirst
	}

	nowTime := carbon.Now()
	refreshTtl := a.config.GetInt("jwt.refresh_ttl")
	if refreshTtl == 0 {
		// 100 years
		refreshTtl = 60 * 24 * 365 * 100
	}

	expireTime := carbon.FromStdTime(auth[a.guard].Claims.ExpiresAt.Time).AddMinutes(refreshTtl)
	if nowTime.Gt(expireTime) {
		return "", errors.AuthRefreshTimeExceeded
	}

	return a.LoginUsingID(auth[a.guard].Claims.Key)
}

func (a *Auth) Logout() error {
	auth, ok := a.ctx.Value(ctxKey).(Guards)
	if !ok || auth[a.guard] == nil || auth[a.guard].Token == "" {
		return nil
	}

	if a.cache == nil {
		return errors.CacheSupportRequired.SetModule(errors.ModuleAuth)
	}

	if err := a.cache.Put(getDisabledCacheKey(auth[a.guard].Token),
		true,
		time.Duration(a.getTtl())*time.Minute,
	); err != nil {
		return err
	}

	delete(auth, a.guard)
	a.ctx.WithValue(ctxKey, auth)

	return nil
}

func (a *Auth) getTtl() int {
	var ttl int
	guardTtl := a.config.Get(fmt.Sprintf("auth.guards.%s.ttl", a.guard))
	if guardTtl == nil {
		ttl = a.config.GetInt("jwt.ttl")
	} else {
		ttl = cast.ToInt(guardTtl)
	}

	if ttl == 0 {
		// 100 years
		ttl = 60 * 24 * 365 * 100
	}

	return ttl
}

func (a *Auth) makeAuthContext(claims *Claims, token string) {
	guards, ok := a.ctx.Value(ctxKey).(Guards)
	if !ok {
		guards = make(Guards)
	}
	guards[a.guard] = &Guard{claims, token}
	a.ctx.WithValue(ctxKey, guards)
}

func (a *Auth) tokenIsDisabled(token string) bool {
	return a.cache.GetBool(getDisabledCacheKey(token), false)
}

func getDisabledCacheKey(token string) string {
	return "jwt:disabled:" + token
}
