package utils

import(
	"os"
)

//GetEnv if the env variable exists gives it otherwise returns the default value
func GetEnv(key, defaultvalue string)string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultvalue
	}
	return value
}