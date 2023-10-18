package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwtgo "github.com/dgrijalva/jwt-go"
	_error "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/spf13/viper"
)

var (
	ErrTokenExpired = "Token is expired"
)

type RealmAccess struct {
	Roles []string
}

func GetValue(r *http.Request, attributes []string) (map[string]interface{}, error) {
	var secretKey = viper.GetString("KEYCLOAK_SECRET_KEY")

	reqToken := strings.ReplaceAll(r.Header.Get("Authorization"), "Bearer ", "")

	key, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(secretKey))
	if err != nil {
		return nil, _error.ErrInvalidConfig
	}

	token, err := jwtgo.Parse(reqToken, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodRSA); !ok {
			err = fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, _error.ErrInvalidToken
		}
		return key, nil
	})
	if err != nil {
		return nil, _error.ErrInvalidToken
	}

	if data, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		result := make(map[string]interface{})

		for _, v := range attributes {
			result[v] = data[v]
		}

		return result, nil
	}

	return nil, _error.ErrInvalidToken
}

func GetRealmAccess(attribute interface{}) (*RealmAccess, error) {
	realmAccessMap, ok := attribute.(map[string]interface{})
	if !ok {
		return nil, errors.New("FAILED TO GET REALM_ACCESS")
	}

	var realmAccess RealmAccess
	if roles, ok := realmAccessMap["roles"].([]interface{}); ok {
		for _, role := range roles {
			if roleName, ok := role.(string); ok {
				realmAccess.Roles = append(realmAccess.Roles, roleName)
			}
		}
	}

	return &realmAccess, nil
}
