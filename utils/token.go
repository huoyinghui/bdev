package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type Claims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

func EncodeB64(src []byte) (retour string) {
	return base64.StdEncoding.EncodeToString(src)
}

func DecodeB64(message string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(message)
}

func SecSecret(uid, salt string) string {
	return ToMd5(fmt.Sprintf("%v:+-*/:%v", salt, uid))
}

func ToMd5(text string) (decode string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(text))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

/*
before CreateToken:
	secret = SecSecret(uid, salt)
then:
	CreateToken(uid, secret, expireToken)
*/
func CreateToken(uid, secret string, expireToken int64) (string, error) {
	if expireToken <= 0 {
		expireToken = time.Now().Add(time.Hour * 1).Unix()
	}
	// 1. Create payload
	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    uid,
		},
	}

	// 2. Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Signs the token with a secret.
	signedToken, err := token.SignedString([]byte(secret))

	return signedToken, err
}

/*
before AuthToken:
	uid, err := GetUid(signedToken)
	if err != nil {
		return "", err
	}
	secret = SecSecret(uid, secret)
then AuthToken(signedToken, secret)
*/
func AuthToken(signedToken, secret string) (string, error) {
	// 1.jwt decode by token & secret
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", fmt.Errorf("AuthToken src.err:%v sercret:%v signedToken:%v", err, secret, signedToken)
	}
	// 2. token valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Uid, nil
	}
	return "", err
}

/*
jwt:
header: part[0]
payload: part[1]
sign: part[2]
*/
func GetUid(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", jwt.NewValidationError("token contains an invalid number of segments", jwt.ValidationErrorMalformed)
	}

	payload := parts[1]
	data, err := DecodeB64(payload)
	if err != nil {
		return "", err
	}
	ob := Claims{}
	err = json.Unmarshal(data, &ob)
	if err != nil {
		return "", err
	}

	return ob.Uid, nil
}
