package main

import (
	_ "go_server/internal/config/api"
	"go_server/pkg/util/log"
	"testing"
)

func TestMain(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("Application crashed with error: %v", r)
		}
	}()
	select {}
}
