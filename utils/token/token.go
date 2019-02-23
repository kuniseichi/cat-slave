package token

import (
	"fmt"
	"time"

	"cat-slave/utils/err"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader   = &err.Errno{Code: 10001, Message: "The length of the `Authorization` header is zero."}
	ErrTokenExpired    = &err.Errno{Code: 10002, Message: "The token has expired."}
	ErrValidateFailed  = &err.Errno{Code: 10003, Message: "The token validate failed."}
	ErrMissingClaims   = &err.Errno{Code: 10004, Message: "claims missing."}
	ErrTokenSignFailed = &err.Errno{Code: 10005, Message: "The token sign failed."}
)

// sign token
func Sign(claims jwt.MapClaims) (string, *err.Errno) {
	claims["createTime"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// secret :=
	tokenString, err := token.SignedString([]byte(viper.GetString("jwt_secret")))
	if err != nil {
		return "", ErrTokenSignFailed
	}
	return tokenString, nil
}

func ParseRequest(c *gin.Context) (map[string]interface{}, *err.Errno) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return nil, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t)
}

func Parse(tokenString string) (map[string]interface{}, *err.Errno) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, validateToken())
	if err != nil {
		return nil, ErrValidateFailed
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrValidateFailed
}

func validateToken() func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(viper.GetString("jwt_secret")), nil
	}
}
