package cached

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

type CacheKey int

const (
	CtxUserId        = CacheKey(12)
	bearerToken      = "Bearer"
	AuthorizationNil = "91"
	InvalidToken     = "92"
	SessionExpired   = "93"
	SystemError      = "99"
)

var errorMessage = map[string]string{
	AuthorizationNil: "The request is missing a valid Token",
	InvalidToken:     "Invalid Token",
	SessionExpired:   "Session has expired",
	SystemError:      "Something went wrong, please try again later, thank you",
}

type SessionMiddleware struct {
	RedisDb dbc.RDbc
}

func (sm *SessionMiddleware) CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var err error
		authorization := r.Header.Get("authorization")
		ctx := r.Context()

		// check apakah authorization pada request header, ada atau tidak.
		if reflect.ValueOf(authorization).IsZero() {
			err = errors.New("authorization nil")

			responseError(rw, http.StatusUnauthorized, AuthorizationNil, getErrorMessage(AuthorizationNil), err)
			return
		}
		// split authorization menjadi array string
		// [0] => token_type (Bearer)
		// [1] => access_token (uuid)
		token := strings.Split(authorization, " ")
		if token[0] != bearerToken {
			err = errors.New("token type invalid")

			responseError(rw, http.StatusUnauthorized, InvalidToken, getErrorMessage(InvalidToken), err)
			return
		}

		res := sm.RedisDb.Get(ctx, token[1])
		// anggap saja jika terjadi error pada redis pencarian
		// di anggap sesi telah berakhir / sessi tidak ditemukan
		// jika tetap ingin memastikan, please uncomment below
		//
		// if res.Err() != nil && res.Err() == redis.Nil {
		if res.Err() != nil {
			logger.Log.Errorf("error redis %v", res.Err())
			responseError(rw, http.StatusUnauthorized, SessionExpired, getErrorMessage(SessionExpired), err)
			return
		}

		sessToken := res.Val()
		user := &model.User{}

		err = json.Unmarshal([]byte(sessToken), &user)
		if err != nil {
			logger.Log.Errorf("error unmarshal token %v", err)
			responseError(rw, http.StatusInternalServerError, SystemError, getErrorMessage(SystemError), err)
			return
		}

		ctx = context.WithValue(ctx, CtxUserId, user.Id)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func getErrorMessage(code string) string {
	return errorMessage[code]
}
