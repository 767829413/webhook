package main

import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/github"
)

func main() {
	payload := flag.String("route", "/webhooks", "webhook path")
	filePath := flag.String("path", "/root", "git pull path")
	port := flag.String("port", "3000", "listen port")
	secret := flag.String("secret", "", "github webhook secret")
	flag.Parse()

	hook, _ := github.New(github.Options.Secret(*secret))
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST(*payload, func(c *gin.Context) {
		payload, err := hook.Parse(c.Request, github.PushEvent)
		if err != nil {
			fmt.Println("error is: " + err.Error())
		}
		payload, ok := payload.(github.PushPayload)
		if ok {
			cmd := fmt.Sprintf("cd %s && git pull", *filePath)
			res, _ := exec.Command("bash", "-c", cmd).CombinedOutput()
			fmt.Println("bash res: " + string(res))
		}
	})
	router.Run(":" + *port)
}
