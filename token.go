package main

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/golang-jwt/jwt"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func CreateFsafeToken(FsafeClientKey, FsafeSecretKey string) string {
	timestr := Time2Str(Now(), YDM)
	return FsafeSecretKey + "::" + GetMD5Hash(FsafeClientKey+"::"+FsafeSecretKey+timestr)
}
func CreateHifptAuthToken(AuthClientKey, AuthSecretKey string) string {
	timestr := Time2Str(Now(), YDM)
	return AuthClientKey + "::" + GetMD5Hash(AuthClientKey+"::"+AuthSecretKey+timestr)
}
func CreateEcomToken(EcomClientKey, EcomSecretKey string) string {
	timestr := Time2Str(Now(), YDM)
	return GetMD5Hash(EcomClientKey + "::" + EcomSecretKey + timestr)
}
func CreateHiCustomerToken(CustomerClientKey, CustomerSecretKey string) string {
	timestr := Time2Str(Now(), YDM)
	return GetMD5Hash(CustomerClientKey + "::" + CustomerSecretKey + timestr)
}

func CreateTokenJwt(info map[string]interface{}) (string, error) {
	mySigningKey := []byte("AllYourBase")
	claims := jwt.MapClaims{}
	claims = info
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}
