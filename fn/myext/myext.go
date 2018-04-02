package myext

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
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
	now := time.Now()
	s.AddAPIMiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("My ext AddAPIMiddlewareFunc - ", now)

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
			
			idToken := ahSplit[1]
			
			var jwtSecret = os.Getenv("ALL_JWT_SECRET")
			// var ownJWTSecret = os.Getenv("OWN_JWT_SECRET")
			// var myJWTSecret = "92d75b10-358d-11e8-a6fc-0a580a340088"
			
			fmt.Println("idToken - " + idToken+"|jwtSecret - " + jwtSecret)
			
			jwtToken, err := jwt.Parse(idToken, func(jt *jwt.Token) (interface{}, error) {
				if _, ok := jt.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", jt.Header["alg"])
				}
				return []byte(jwtSecret), nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, err.Error())
				return
			}

			if !jwtToken.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, errors.New("Invalid authorization token, access denied").Error())
				return
			}
			var claims jwt.MapClaims
			var ok bool
			if claims, ok = jwtToken.Claims.(jwt.MapClaims); !ok {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, errors.New("Invalid authorization token, invalid claims, access denied").Error())
				return
			}
			// ok, so finally we're good
			username := claims["username"].(string)
			fmt.Println("claims username - " + username)

			appNameV := ctx.Value(fnext.AppNameKey)
			if appNameV != nil {
				appName := appNameV.(string)
				fmt.Println("appName - " + appName)
			}

			fmt.Println("My ext ServeHTTP - ", time.Now())

			next.ServeHTTP(w, r)
		})
	})

	s.AddRootMiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			fmt.Println("My ext AddRootMiddlewareFunc - ", now)

			next.ServeHTTP(w, r)
		})
	})
	return nil
}
