package middleware

import (
	"net/http"

	"github.com/devheaven-platform/auth-service/pkg/utils/transport"
	"github.com/go-chi/jwtauth"
)

// Authenticator is used as a helper function to check
// if a valid jwt token is provided with the request.
// If no token or an invalid token is provided an 401
// error will be returned. Otherwise the next function
// is called.
func Authenticator(next http.Handler) http.Handler {
	t := transport.BaseHTTPTransport{}
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		token, _, err := jwtauth.FromContext(req.Context())

		if err != nil || token == nil || !token.Valid {
			t.RespondError(res, "Your not authorized to access this resource", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(res, req)
	})
}
