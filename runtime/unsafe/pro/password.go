package pro

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	_ "unsafe" // for go:linkname
)

//go:linkname mdEncode github.com/coulsonzero/gopkg/pro.MD5Encode
func mdEncode(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

//go:linkname hashPassword github.com/coulsonzero/gopkg/pro.HashPassword
// hashPassword 加密
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//go:linkname checkPasswordHash github.com/coulsonzero/gopkg/pro.CheckPasswordHash
// checkPasswordHash 解密
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
