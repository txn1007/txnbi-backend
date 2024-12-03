package encry

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeByMd5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
