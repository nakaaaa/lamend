package lamend

import (
	"strings"

	"github.com/alecthomas/kong"
)

type Cli struct {
	Config string `help:"Config file" default:"lamend.yml"`

	Lambda *lambdaOption `cmd:"" help:"restart lambda"`
}

type lambdaOption struct{}

func Parse(args []string) (string, *Cli, error) {
	var cli Cli
	parser, err := kong.New(&cli)
	if err != nil {
		return "", nil, err
	}
	c, err := parser.Parse(args)
	if err != nil {
		return "", nil, err
	}
	sub := strings.Fields(c.Command())[0]

	return sub, &cli, nil
}
