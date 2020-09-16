package util

import (
	"crypto/md5"
	"io/ioutil"
	"os/user"
	"runtime"
	"strings"
)



func Check(err error) {
	if err != nil {
		panic(err)
	}
}


func UserInfo() (string, string) {
	user, _ := user.Current()
	os := runtime.GOOS
	bytes := []byte(user.Name)
	if len(bytes) == 0 {
		return user.Username, strings.Title(strings.ToLower(os))
	} else {
		return user.Name, strings.Title(strings.ToLower(os))
	}
}

func SumCheck(path string) ([16]byte){
	data, err := ioutil.ReadFile(path)
	Check(err)
	return md5.Sum(data)
}
