package curruser

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Claims Used to determine auth token passed in header.
type Claims struct {
	jwt.MapClaims
	UUID           string
	Email          string
	StandardClaims *jwt.StandardClaims
}

//GetCurrUser Gets current from access token.
func GetCurrUser(w http.ResponseWriter, r *http.Request) string {
	header := strings.TrimSpace(r.Header.Get("x-access-token"))

	claims := Claims{}

	header = strings.Replace(header, "Bearer ", "", -1)
	_, err := jwt.ParseWithClaims(header, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		panic(err)
	}
	return claims.UUID
}
