package syslog

import (
	"io"
	"log"
	"os"
)

var (
	Debug *log.Logger
)

func init() {
	debug_file, err := os.OpenFile("syslog/debug.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0766)

	if err != nil {
		log.Fatal(err)
	}

	Debug = log.New(io.MultiWriter(debug_file, os.Stdout), "DEBUG: ", log.Ltime)

}
