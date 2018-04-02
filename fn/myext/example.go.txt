package myext

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
)

func init() {
	server.RegisterExtension(&myExt{})
}

type myExt struct {
}

func (e *myExt) Name() string {
	return "github.com/iwangxiaodong/container-config/fn/myext"
}

func (e *myExt) Setup(s fnext.ExtServer) error {
	s.AddAPIMiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("My ext - ", time.Now())

			ctx := r.Context()
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader == "" {
				server.WriteError(ctx, w, http.StatusUnauthorized, errors.New("No Authorization header, access denied"))
				return
			}

			ahSplit := strings.Split(authorizationHeader, " ")
			if len(ahSplit) != 2 {
				server.WriteError(ctx, w, http.StatusUnauthorized, errors.New("Invalid authorization header, access denied"))
				return
			}

			if ahSplit[1] != "token13572468" {
				server.WriteError(ctx, w, http.StatusUnauthorized, errors.New("Invalid authorization token, access denied"))
				return
			}

			fmt.Println("My ext ServeHTTP - ", time.Now())

			next.ServeHTTP(w, r)
		})
	})
	return nil
}
