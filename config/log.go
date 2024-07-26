package config

import (
	"log"
	"os"
)

//InitLogz - ...
func InitLogz() {
	// create logfile
	c := NewConfig()
	//n := "logs/logfile.log"
	n := c.DIRlogs + "/" + c.ARQlog
	f, err := os.OpenFile(n, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		os.Exit(1)
	}
	log.SetOutput(f)

	// flags to log timestamp and lineno
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
