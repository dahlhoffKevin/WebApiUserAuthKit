package WebApiAuthKit

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dahlhoffKevin/WebApiAuthKit/authKitHttpErrorHandler"
)

func checkBearerTokenIntegritiy(bearerToken string) *authKitHttpErrorHandler.HTTPError {
	//vallidate bearertoken in db
	fmt.Printf("Bearer-Token to check: %v", bearerToken)
	return authKitHttpErrorHandler.New(http.StatusNotImplemented, "not yet implemented")
}

func getBearerTokenFromRequestHeader(r *http.Request) (string, *authKitHttpErrorHandler.HTTPError) {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return "", authKitHttpErrorHandler.New(http.StatusUnauthorized, "could not validate bearer token")
	}

	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return "", authKitHttpErrorHandler.New(http.StatusUnauthorized, "could not validate bearer token format")
	}

	return splitToken[1], nil
}

func AuthMiddlewareBearer(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[Request] %v -> %v ", r.Method, r.URL.Path)

		bearerToken, err := getBearerTokenFromRequestHeader(r)
		if err != nil {
			authKitHttpErrorHandler.Write(w, err)
			return
		}
		if bearerToken == "" {
			authKitHttpErrorHandler.Write(w, authKitHttpErrorHandler.New(http.StatusUnauthorized, "could not validate bearer token"))
			return
		}

		errAuth := checkBearerTokenIntegritiy(bearerToken)
		if errAuth != nil {
			authKitHttpErrorHandler.Write(w, errAuth)
			fmt.Printf("[unauthorized: %v]\n", errAuth.Error())
			return
		}

		next(w, r)
		fmt.Printf("took: %vms [authenticated]\n", time.Since(start).Milliseconds())
	}
}
