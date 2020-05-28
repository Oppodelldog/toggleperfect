package config

import (
	"log"
	"os"
	"strconv"
)

var EnableDeviceUI bool
var EnableRemoteUI bool

func init() {
	EnableDeviceUI = getEnvBool("TP_ENABLE_DEVICE_UI", true)
	EnableRemoteUI = getEnvBool("TP_ENABLE_REMOTE_UI", false)
}

func getEnvBool(envVarName string, defaultValue bool) bool {
	if v, ok := os.LookupEnv(envVarName); ok {
		b, err := strconv.ParseBool(v)
		if err != nil {
			log.Printf("could not parse boolean from env var '%s, got %v", envVarName, v)
		}
		return b
	}

	return defaultValue
}
