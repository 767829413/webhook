package options

import "github.com/spf13/cobra"

type ServerRunOptions struct {
	Route  string
	Path   string
	Port   string
	Secret string
}

func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{}
	return &s
}

func (s *ServerRunOptions) Flags(cmd *cobra.Command) {
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd.Flags().StringVar(&s.Route, "route", "/webhooks", "webhook path")
	cmd.Flags().StringVar(&s.Path, "path", "/root", "git pull path")
	cmd.Flags().StringVar(&s.Port, "port", "3000", "listen port")
	cmd.Flags().StringVar(&s.Secret, "secret", "", "github webhook secret")
}
