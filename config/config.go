package config

import (
	"os"
)

// AllConfig - ...
var AllConfig *Config

// Config - ...
type Config struct {
	APPenv         string
	APIport        string
	DIRlogs        string
	ARQlog         string
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
func NewConfig() *Config {
	AllConfig = &Config{
		APPenv:  getConfigValue("APP_ENV", "dev"),
		APIport: getConfigValue("API_PORT", "8888"),
		DIRlogs: getConfigValue("DIR_LOGS", "logs"),
		ARQlog:  getConfigValue("ARQ_LOG", "logfile.log"),
		//ORACLEhost:     getConfigValue("ORACLE_HOST", "192.168.1.1"),
		ORACLEhost: getConfigValue("ORACLE_HOST", "192.168.37.34"),
		ORACLEport: getConfigValue("ORACLE_PORT", "1521"),
		//ORACLEuser:     getConfigValue("ORACLE_USER", "user1"),
		ORACLEuser: getConfigValue("ORACLE_USER", "api_oracle_hom"),
		//ORACLEpassword: getConfigValue("ORACLE_PASSWORD", "user1"),
		ORACLEpassword: getConfigValue("ORACLE_PASSWORD", "api_oracle_hom"),
		//ORACLEservice:  getConfigValue("ORACLE_SERVICE", "orcl"),
		ORACLEservice: getConfigValue("ORACLE_SERVICE", "isapfhom"),
		ORACLElibdir:  getConfigValue("ORACLE_LIB_DIR", "/opt/oracle/instantclient_19_9/"),
		TESTEhost:     getConfigValue("TESTE_HOST", "192.168.1.1"),
		TESTEport:     getConfigValue("TESTE_PORT", "1521"),
		TESTEuser:     getConfigValue("TESTE_USER", "user2"),
		TESTEpassword: getConfigValue("TESTE_PASSWORD", "user2"),
		TESTEservice:  getConfigValue("TESTE_SERVICE", "orcl_teste"),
		TESTElibdir:   getConfigValue("TESTE_LIB_DIR", "/opt/oracle/instantclient_19_9/"),
		APIabc01:      getConfigValue("URL_API_ABC01", ""),
	}
	return AllConfig
}
