package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/andersondelgado/prueba_go_graphql/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"strings"
)

type CustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func EncodeToken(claims CustomClaims) (string, error) {
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokensJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	appKey := config.Environment.Key
	key := []byte(appKey)

	tokenStr, er := tokensJWT.SignedString(key)
	if er != nil {
		fmt.Println("error: ", er)
		return "", er
	}

	return tokenStr, nil
}

func DecodeToken(strToken string) (string, error) {
	appKey := config.Environment.Key
	tokenObj, err := jwt.ParseWithClaims(strToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(appKey), nil
	})
	if claims, ok := tokenObj.Claims.(*CustomClaims); ok && tokenObj.Valid {
		var id = strconv.FormatUint(uint64(claims.ID), 10)
		return id, nil
	} else {
		return "", err
	}
}

func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func EncryptedIv(key string, iv string, plainText string) string {

	enc0 := base64.URLEncoding.EncodeToString([]byte(plainText))
	keyEnc := base64.URLEncoding.EncodeToString([]byte(key))
	iVEnc := base64.URLEncoding.EncodeToString([]byte(iv))

	str := key + "|" + enc0
	enc := base64.URLEncoding.EncodeToString([]byte(str))
	encF0 := strings.Split(enc, "==")
	keyEncF := strings.Split(keyEnc, "==")
	iVEncF := strings.ReplaceAll(iVEnc, "==", "")
	encf := encF0[0] + "-" + keyEncF[0] + "." + iVEncF
	return encf
}

func DecryptedIv(key string, iv string, encripted string) (string, error) {
	encriptedFF := strings.Split(encripted, "-")
	encriptedFF2 := strings.Split(encripted, ".")
	encriptedF := encriptedFF[0] + "=="

	encriptedF1 := encriptedFF2[1] + "=="

	dec00, _ := base64.URLEncoding.DecodeString(encriptedFF[1] + "==")

	dec0, _ := base64.URLEncoding.DecodeString(encriptedF)
	dec0x, _ := base64.URLEncoding.DecodeString(encriptedF1)
	keyc := string(dec00)
	split := strings.Split(string(dec0), "|")
	splitE := strings.Split(string(split[1]), ".")
	dec1, _ := base64.URLEncoding.DecodeString(splitE[0])
	if keyc != key && iv != string(dec0x) {
		return "", errors.New("invalid key")
	}
	return string(dec1), nil
}
