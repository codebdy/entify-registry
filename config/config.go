package config

import "github.com/spf13/viper"

var theConfig config

type config struct {
	v *viper.Viper
}

func Init() {

}

func GetString(key string) string {
	return ""
}

func GetBoolean(key string) bool {
	return false
}

func SetValue(key string, value interface{}) {

}
