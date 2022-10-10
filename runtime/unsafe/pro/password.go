package runtime

import (
	"crypto/md5"
	"encoding/hex"
	_ "unsafe" // for go:linkname
)

//go:linkname mdEncode github.com/coulsonzero/gopkg/pro.MD5Encode
func mdEncode(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
