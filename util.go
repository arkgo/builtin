package builtin

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

//密码加密格式
func Hash(str string) string {
	return Sha1(str)
}

//sha1加密
func Sha1(str string) string {
	sha1Ctx := sha1.New()
	sha1Ctx.Write([]byte(str))
	cipherStr := sha1Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//sha1加密文件
func Sha1File(file string) string {
	if f, e := os.Open(file); e == nil {
		defer f.Close()

		h := sha1.New()
		if _, e := io.Copy(h, f); e == nil {
			return fmt.Sprintf("%x", h.Sum(nil))
		}
	}
	return ""
}

//md5加密
func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//md5加密文件
func Md5File(file string) string {
	if f, e := os.Open(file); e == nil {
		defer f.Close()

		h := md5.New()
		if _, e := io.Copy(h, f); e == nil {
			return fmt.Sprintf("%x", h.Sum(nil))
		}
	}
	return ""
}
