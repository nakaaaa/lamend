package lamend

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type App struct {
	cliConfig    *Cli
	lamendConfig *Config
	AWSClient    *AWSClient
}

func newApp(ctx context.Context, clicfg *Cli) (*App, error) {
	b, err := Read(clicfg.Config)
	if err != nil {
		return nil, err
	}
	cfg, err := UnmarshalConfigYAML(b)
	if err != nil {
		return nil, err
	}
	awscli, err := initAWS(ctx)
	if err != nil {
		return nil, err
	}
	return &App{
		cliConfig:    clicfg,
		lamendConfig: cfg,
		AWSClient:    awscli,
	}, nil
}

func Start(ctx context.Context) (int, error) {
	sub, clicfg, err := Parse(os.Args[1:])
	if err != nil {
		return 1, err
	}
	app, err := newApp(ctx, clicfg)
	if err != nil {
		return 1, err
	}

	switch sub {
	case "lambda":
		if err := runLambda(ctx, app); err != nil {
			return 1, err
		}
	}
	return 0, nil
}

func runLambda(ctx context.Context, app *App) error {
	params := &lambda.UpdateFunctionCodeInput{
		FunctionName: aws.String(app.lamendConfig.FunctionName),
		ImageUri:     aws.String(app.lamendConfig.ImageURI),
	}
	_, err := app.AWSClient.Lambda.UpdateFunctionCode(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
