package lamend

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

const defaultRegion = "ap-northeast-1"

type AWSClient struct {
	Lambda ILambda
}

func initAWS(ctx context.Context) (*AWSClient, error) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = defaultRegion
	}
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}

	return &AWSClient{
		Lambda: NewLambda(&cfg),
	}, nil
}

type ILambda interface {
	UpdateFunctionCode(ctx context.Context, params *lambda.UpdateFunctionCodeInput, optFns ...func(*lambda.Options)) (*lambda.UpdateFunctionCodeOutput, error)
}

type Lambda struct {
	cfg *aws.Config
}

func NewLambda(cfg *aws.Config) ILambda {
	return &Lambda{
		cfg: cfg,
	}
}

func (l *Lambda) UpdateFunctionCode(ctx context.Context, params *lambda.UpdateFunctionCodeInput, optFns ...func(*lambda.Options)) (*lambda.UpdateFunctionCodeOutput, error) {
	if params == nil {
		params = &lambda.UpdateFunctionCodeInput{}
	}

	slog.Info(fmt.Sprintf("[INFO] FunctionName: %s", *params.FunctionName))
	slog.Info(fmt.Sprintf("[INFO] ImageURI: %s", *params.ImageUri))
	resp, err := lambda.NewFromConfig(*l.cfg).UpdateFunctionCode(ctx, params, optFns...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
