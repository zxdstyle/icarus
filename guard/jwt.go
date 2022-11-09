package guard

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"github.com/zxdstyle/icarus/server/requests"
	"strings"
	"time"
)

type jwtGuard struct {
	secret []byte
	expire time.Duration
}

const (
	HeaderAuthorization = "Authorization"
	jwtGuardContextKey  = "guard-user"
)

var (
	ErrMissingToken = errors.New("missing or malformed JWT")
)

func NewJwtGuard(secret []byte, expire time.Duration) Guard {
	return &jwtGuard{
		secret: secret,
		expire: expire,
	}
}

func (j jwtGuard) Check(req requests.Request) error {
	token := req.GetHeader(HeaderAuthorization)
	token = strings.Replace(token, "Bearer ", "", 1)
	if len(token) == 0 {
		return ErrMissingToken
	}
	token = strings.Replace(token, "Bearer ", "", 1)

	t, e := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return j.secret, nil
	})
	if e != nil {
		return e
	}

	if !t.Valid {
		return fmt.Errorf("invalid token")
	}

	uid := t.Claims.(jwt.MapClaims)["user_id"]

	req.Context(jwtGuardContextKey, cast.ToUint(uid))
	return nil
}

func (j jwtGuard) LoginUsingID(id uint) (any, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(j.expire).Unix()
	return token.SignedString(j.secret)
}

func (j jwtGuard) ID(req requests.Request) uint {
	return cast.ToUint(req.Value(jwtGuardContextKey))
}
