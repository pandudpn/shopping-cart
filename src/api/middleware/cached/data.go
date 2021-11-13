package cached

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (rm *RedisMiddleware) CachedData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var (
			urlPath     = r.URL.Path
			queryString = r.URL.Query()
			key         = urlPath
			ctx         = r.Context()
			data        = make(map[string]interface{})
		)
		key = strings.ReplaceAll(key, "/", "")

		if limit, found := queryString["limit"]; found {
			key = fmt.Sprintf("%s-%s", key, limit[0])
		} else {
			key = fmt.Sprintf("%s-20", key)
		}

		if page, found := queryString["page"]; found {
			key = fmt.Sprintf("%s-%s", key, page[0])
		} else {
			key = fmt.Sprintf("%s-1", key)
		}

		if search, found := queryString["search"]; found {
			if !reflect.ValueOf(search).IsZero() {
				key = fmt.Sprintf("%s-%s", key, search[0])
			}
		}

		res := rm.RedisDb.Get(ctx, key)
		if res.Err() != nil {
			logger.Log.Error(res.Err())

			next.ServeHTTP(rw, r)
			return
		}

		err := json.Unmarshal([]byte(res.Val()), &data)
		if err != nil {
			logger.Log.Error(err)

			next.ServeHTTP(rw, r)
			return
		}

		if len(data) > 0 {
			responseSuccess(rw, data)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
