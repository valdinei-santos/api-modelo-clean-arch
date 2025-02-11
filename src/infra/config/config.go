package config

import (
	"fmt"
	"os"
)

// AllConfig - ...
var AllConfig *config

// Config - ...
type config struct {
	APPenv    string
	APIport   string
	LogOutput string
	LogLevel  string
	//DIRlogs        string
	//ARQlog         string
	ORACLEhost     string
	ORACLEport     string
	ORACLEuser     string
	ORACLEpassword string
	ORACLEservice  string
	ORACLElibdir   string
	TESTEhost      string
	TESTEport      string
	TESTEuser      string
	TESTEpassword  string
	TESTEservice   string
	TESTElibdir    string
	APIabc01       string
}

func getConfigValue(envName string, defaultValue string) string {
	if value, ok := os.LookupEnv(envName); ok {
		return value
	}
	return defaultValue
}

// NewConfig - ...
// func NewConfig() *Config {
// init config
func init() {
	AllConfig = &config{
		APPenv:         getConfigValue("APP_ENV", "dev"),
		APIport:        getConfigValue("API_PORT", "8888"),
		LogOutput:      getConfigValue("LOG_OUTPUT", "logs/logfile.log"),
		LogLevel:       getConfigValue("LOG_LEVEL", "info"),
		ORACLEhost:     getConfigValue("ORACLE_HOST", "192.168.37.34"),
		ORACLEport:     getConfigValue("ORACLE_PORT", "1521"),
		ORACLEuser:     getConfigValue("ORACLE_USER", "user1"),
		ORACLEpassword: getConfigValue("ORACLE_PASSWORD", "user1"),
		ORACLEservice:  getConfigValue("ORACLE_SERVICE", "orcl"),
		ORACLElibdir:   getConfigValue("ORACLE_LIB_DIR", "/opt/oracle/instantclient_19"),
		TESTEhost:      getConfigValue("TESTE_HOST", "192.168.1.1"),
		TESTEport:      getConfigValue("TESTE_PORT", "1521"),
		TESTEuser:      getConfigValue("TESTE_USER", "user2"),
		TESTEpassword:  getConfigValue("TESTE_PASSWORD", "user2"),
		TESTEservice:   getConfigValue("TESTE_SERVICE", "orcl_teste"),
		TESTElibdir:    getConfigValue("TESTE_LIB_DIR", "/opt/oracle/instantclient_19"),
		APIabc01:       getConfigValue("URL_API_ABC01", ""),
	}
	fmt.Println("Carregou configuração...")
	//return AllConfig
}
