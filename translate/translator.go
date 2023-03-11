package translate

import (
	"strings"

	"gopkg.in/ini.v1"
)

var messages = map[string]string{}

func LoadLang(file string) {
	load, err := ini.Load(file)
	if err != nil {
		return
	}

	for _, key := range load.Section("").Keys() {
		messages[key.Name()] = key.String()
	}
}

func Translate(key string, params []string) string {
	key = key[1:]

	if _, ok := messages[key]; ok {
		msg := messages[key]

		for _, v := range params {
			msg = strings.Replace(msg, "%s", v, 1)
		}

		return msg
	}

	return key
}
