package utils

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type EnvTypes struct {
	Env      string `key:"env" default:"prod"`
	HttpPort int    `key:"HTTP_PORT" default:"8080"`
}

func InitEnv() {
	EnvData = &EnvTypes{}
	refEnv := reflect.ValueOf(EnvData).Elem()
	refType := refEnv.Type()
	for i := 0; i < refType.NumField(); i++ {
		fType := refType.Field(i)
		fVal := refEnv.Field(i)
		tags := fType.Tag
		key := tags.Get("key")
		if key == "" {
			continue
		}
		_, required := tags.Lookup("required")

		val := fVal.Addr().Interface()
		loadEnv(val, key, tags.Get("default"), required)
	}

}

var EnvData *EnvTypes

func GetEnvs() *EnvTypes {
	return EnvData
}

var envFunc = os.Getenv

// loadEnv reads and parse the value corresponds to key.
func loadEnv(i interface{}, key, fallback string, required bool) {
	v := envFunc(key)
	if v == "" {
		if required {
			panic(fmt.Errorf("%v is required env", key))
		}
		v = fallback
	}
	switch val := i.(type) {
	case *string:
		*val = v
	case *int:
		*val, _ = strconv.Atoi(v)
	}
}
