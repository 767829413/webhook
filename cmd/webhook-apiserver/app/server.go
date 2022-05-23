package app

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/767829413/webhook/cmd/webhook-apiserver/app/options"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/github"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	cmd := &cobra.Command{
		Use:  "webhook-apiserver",
		Long: `This is the callback service used to receive events from the github repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(s)
		},
	}
	s.Flags(cmd)
	return cmd
}

func Run(opt *options.ServerRunOptions) error {
	hook, _ := github.New(github.Options.Secret(opt.Secret))
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.POST(opt.Route, func(c *gin.Context) {
		payload, err := hook.Parse(c.Request, github.PushEvent)
		if err != nil {
			fmt.Println("error is: " + err.Error())
		}
		payload, ok := payload.(github.PushPayload)
		if ok {
			cmd := fmt.Sprintf("cd %s && git pull", opt.Path)
			res, _ := exec.Command("bash", "-c", cmd).CombinedOutput()
			fmt.Println("bash res: " + string(res))
		}
	})
	var eg errgroup.Group
	eg.Go(func() error {
		if err := router.Run(":" + opt.Port); err != nil {
			fmt.Printf("Fail to listening: "+":"+opt.Port+", error info: %s", err.Error())
		}
		return nil
	})
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := eg.Wait(); err != nil {
		fmt.Printf("Waiting for a concurrent service failure: %s", err.Error())
	}

	return nil
}
