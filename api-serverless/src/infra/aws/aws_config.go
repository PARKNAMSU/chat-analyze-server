package aws_config

import (
	"context"

	"chat-platform-api.com/chat-platform-api/src/tool/logging_tool"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	Region             = "ap-southeast-2"
	cfg    *aws.Config = getAWSConfig()
)

func getAWSConfig() *aws.Config {
	config, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRetryer(func() aws.Retryer {
			return retry.AddWithMaxAttempts(retry.NewStandard(), 1)
		}),
		config.WithRegion(Region),
	)

	if err != nil {
		logging_tool.PrintErrorLog("getAWSConfig", err.Error())
		return nil
	}

	return &config
}
