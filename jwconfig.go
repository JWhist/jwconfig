package jwconfig

import (
	"os"
	"reflect"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

func LoadFile[T any](path string) T {
	var t T
	err := Load(path, &t)
	if err != nil {
		panic(err)
	}
	return t
}

func Load(configFile string, obj interface{}) error {
	content, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	if reflect.TypeOf(obj).Kind() == reflect.Pointer {
		typ := reflect.ValueOf(obj).Elem().Type()
		if typ.Kind() == reflect.Struct {
			if err := defaults.Set(obj); err != nil {
				return err
			}
		}
	}

	return yaml.Unmarshal(content, obj)
}
