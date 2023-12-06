package lamend

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

const defaultRegion = "ap-northeast-1"

func initAWS(ctx context.Context) (*aws.Config, error) {
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

	return &cfg, nil
}
