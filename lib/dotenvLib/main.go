package dotenvLib

import (
	"booking/constants"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	once sync.Once
)

func loadEnv() {

	viper.SetConfigName("config")   // Name of the config file (without extension)
	viper.SetConfigType("yaml")     // Config file type
	viper.AddConfigPath("./config") // Path to look for the config file in the config directory

	err := viper.ReadInConfig()

	constants.HandleErrorByPanic(err)

	envVariables, _ := json.Marshal(viper.Get("env"))

	fmt.Println(string(envVariables))
}

func GetEnv(keys ...string) interface{} {
	once.Do(loadEnv)

	return getYAMLValue(keys...)

}

func getYAMLValue(keys ...string) string {
	if len(keys) == 0 {
		return ""
	}

	v := viper.Get(keys[0])

	if v == nil {
		return "" // Key not found
	}

	// Recursively dive into nested levels
	for _, key := range keys[1:] {
		if subMap, ok := v.(map[string]interface{}); ok {
			v = subMap[key]
		} else {
			return "" // Key not found in nested structure
		}
	}

	// Convert the final value to string if possible
	if strValue, ok := v.(string); ok {
		return strValue
	}

	// Handle other types if needed
	return fmt.Sprintf("%v", v)
}
