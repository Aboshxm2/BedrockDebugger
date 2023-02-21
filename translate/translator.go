package translate

import (
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
)

var messages *ini.Section

func LoadLang(file string) {
	load, err := ini.Load(file)
	if err != nil {
		return
	}

	messages = load.Section("")
}

func Translate(key string, params []string) string {
	key = key[1:]

	if messages.HasKey(key) {
		msg := messages.Key(key).String()

		for i, v := range params {
			msg = strings.ReplaceAll(msg, "{%"+strconv.Itoa(i)+"}", v)
		}

		return msg
	}

	return key
}
