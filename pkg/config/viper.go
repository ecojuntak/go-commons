package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

func NewEnvConfig() *viper.Viper {
	fang := viper.New()

	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()
	return fang
}

func NewConfig(configName string, configPath string, prefix string) *viper.Viper {
	fang := viper.New()

	if len(strings.TrimSpace(prefix)) > 0 {
		fang.SetEnvPrefix(prefix)
	}
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()

	fang.SetConfigName(configName)
	fang.AddConfigPath(".")
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		fang.AddConfigPath(configPath)
	}
	_ = fang.ReadInConfig()

	return fang
}

func GetStringDefault(viper *viper.Viper, key string, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func GetIntDefault(viper *viper.Viper, key string, defaultValue int) int {
	viper.SetDefault(key, defaultValue)
	return viper.GetInt(key)
}

func GetArrayString(viper *viper.Viper, key string) []string {
	return GetStringSplit(viper, key, ",")
}

func GetStringSplit(viper *viper.Viper, key string, separator string) []string {
	return strings.Split(viper.GetString(key), separator)
}

func GetBoolDefault(viper *viper.Viper, key string, defaultValue bool) bool {
	viper.SetDefault(key, defaultValue)
	return viper.GetBool(key)
}
