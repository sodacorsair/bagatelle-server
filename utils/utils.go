package utils

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func ResponseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func CryptPwd(unique, pwd string) string {
	salt := "@@$%" + unique + "!^&*#"
	tmpByte, _ := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32)
	return fmt.Sprintf("%x", tmpByte)
}

func TimeFormat(time time.Time) string {
	timeStr := time.String()
	timeStr = strings.Replace(timeStr, "T", " ", -1)
    re := regexp.MustCompile(`^(.*)\+`)
    sub := re.FindSubmatch([]byte(timeStr))
    return string(sub[1])
}