package main

import (
	_ "go_server/internal/config/api"

	"go_server/pkg/util/log"
	// "net/http"
	// _ "net/http/pprof"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("Application crashed with error: %v", r)
		}
	}()
	select {}
}
