package WebApiAuthKit

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dahlhoffKevin/WebApiAuthKit/errorhandler"
)

func checkBearerTokenIntegritiy(bearerToken string) *errorhandler.HTTPError {
	//vallidate bearertoken in db
	fmt.Printf("Bearer-Token to check: %v", bearerToken)
	return errorhandler.New(http.StatusNotImplemented, "not yet implemented")
}

func getBearerTokenFromRequestHeader(r *http.Request) (string, *errorhandler.HTTPError) {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return "", errorhandler.New(http.StatusUnauthorized, "could not validate bearer token")
	}

	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return "", errorhandler.New(http.StatusUnauthorized, "could not validate bearer token format")
	}

	return splitToken[1], nil
}

func AuthMiddlewareBearer(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[Request] %v -> %v ", r.Method, r.URL.Path)

		bearerToken, err := getBearerTokenFromRequestHeader(r)
		if err != nil {
			errorhandler.Write(w, err)
			return
		}
		if bearerToken == "" {
			errorhandler.Write(w, errorhandler.New(http.StatusUnauthorized, "could not validate bearer token"))
			return
		}

		errAuth := checkBearerTokenIntegritiy(bearerToken)
		if errAuth != nil {
			errorhandler.Write(w, errAuth)
			fmt.Printf("[unauthorized: %v]\n", errAuth.Error())
			return
		}

		next(w, r)
		fmt.Printf("took: %vms [authenticated]\n", time.Since(start).Milliseconds())
	}
}
