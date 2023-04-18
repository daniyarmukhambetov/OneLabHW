package transport

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"hw1/pkg"
	"net/http"
	"os"
)

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			tknStr := request.Header.Get("Authorization")
			claims := &pkg.Claims{}
			tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
				return os.Getenv("JWT_SECRET"), nil
			})
			if err != nil {
				ctx := context.WithValue(context.Background(), "jwt_error", err.Error())
				next.ServeHTTP(writer, request.WithContext(ctx))
				return
			}
			if !tkn.Valid {
				ctx := context.WithValue(context.Background(), "username", "")
				next.ServeHTTP(writer, request.WithContext(ctx))
				return
			}
			ctx := context.WithValue(request.Context(), "username", claims.Username)
			next.ServeHTTP(writer, request.WithContext(ctx))
		},
	)
}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tknStr := r.Header.Get("Authorization")
		claims := &pkg.Claims{}
		jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return os.Getenv("JWT_SECRET"), nil
		})
		ctx := context.WithValue(r.Context(), "user", claims.Username)
		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		//
		// note: context.Context values are nested, so any previously set
		// values will be accessible as well, and the new `"user"` key
		// will be accessible from this point forward.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
