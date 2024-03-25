package env

import (
	"reflect"
	"strings"
)

func LoadEnvironmentVariables() EnvironmentVariables {
	vars := EnvironmentVariables{}
	err := loadType(vars, "")
	if err != nil {
		// TODO: Panic
	}
	return vars
}

func loadType(container interface{}, prefix string) error {
	t := reflect.TypeOf(container)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if len(prefix) > 0 {
			prefix = prefix + "_"
		}
		key := strings.ToUpper(prefix + field.Name)
		println(key)
		if field.Type.Kind() == reflect.Struct {

		}
	}

	return nil
}

var Environment = LoadEnvironmentVariables()
