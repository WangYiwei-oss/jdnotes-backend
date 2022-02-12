package helper

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func Md5Str(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}

func Md5Data(data map[string]string) string {
	str := strings.Builder{}
	for k, v := range data {
		str.WriteString(k)
		str.WriteString(v)
	}
	return Md5Str(str.String())
}

func CmIsEq(cm1 map[string]string, cm2 map[string]string) bool {
	return Md5Data(cm1) == Md5Data(cm2)
}
