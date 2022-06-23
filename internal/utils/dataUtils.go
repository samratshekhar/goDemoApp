package utils

import (
	"goDemoApp/deployments/logging"
	"goDemoApp/internal/config"
	"strconv"
)

func ConvertStringToInt(input string, defaultVal int) int {
	log := logging.GetLogger()
	i, err := strconv.Atoi(config.GetConfig().HTTPServerConfig.IdleTimeoutSeconds)
	if err != nil {
		log.Errorf("Error in converting string to int for val %s, with error: %s", input, err.Error())
		return defaultVal
	}
	return i
}
