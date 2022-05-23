/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/767829413/webhook/cmd/webhook-apiserver/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	cmd := app.NewAPIServerCommand()
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
