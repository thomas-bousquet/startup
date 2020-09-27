package middlewares

import (
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	. "github.com/thomas-bousquet/startup/errors"
	"github.com/thomas-bousquet/startup/repositories"
	errorHandler "github.com/thomas-bousquet/startup/utils/error-handler"
	"github.com/thomas-bousquet/startup/utils/jwt"
	"net/http"
	"strings"
)

type AuthenticationMiddleware struct {
	jwt jwt.JWT
	userRepository repositories.UserRepository
}

func NewAuthenticationMiddleware(jwt jwt.JWT, userRepository repositories.UserRepository) AuthenticationMiddleware {
	return AuthenticationMiddleware{
		jwt: jwt,
		userRepository: userRepository,
	}
}

func (m AuthenticationMiddleware) ExecuteWithRole(role string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authorizationHeaderParts := strings.Fields(r.Header.Get("Authorization"))

			if len(authorizationHeaderParts) < 2 || strings.ToLower(authorizationHeaderParts[0]) != "bearer" {
				authorizationError := NewAuthorizationError("Authorization header is not valid")
				errorHandler.WriteJSONErrorResponse(w, authorizationError)
				return
			}

			authorizationToken := authorizationHeaderParts[1]

			token, err := m.jwt.ParseToken(authorizationToken)

			if err != nil {
				errorHandler.WriteJSONErrorResponse(w, err)
				return
			}

			claims := token.Claims.(*jwtGo.StandardClaims)
			userId := claims.Subject

			user, err := m.userRepository.FindUserWithRole(userId, role)

			if err != nil {
				unexpectedError := NewUnexpectedError()
				errorHandler.WriteJSONErrorResponse(w, unexpectedError)
				return
			}

			if user == nil {
				authorizationError := NewAuthorizationError("Authorization header is not valid")
				errorHandler.WriteJSONErrorResponse(w, authorizationError)
				return
			}

			context.Set(r, "user_id", &claims.Subject)
			next.ServeHTTP(w, r)
		})
	}

}
