package aws_config

import (
	"context"
	"strings"
	"time"

	"chat-platform-api.com/chat-platform-api/src/tool/logging_tool"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

var (
	s3Client *S3Client
)

type S3Client struct {
	client   *s3.Client
	uploader *manager.Uploader
}

type UploadInput struct {
	Bucket      string
	Key         string
	Body        string
	ContentType string
}

func GetS3Client() *S3Client {
	if s3Client == nil {
		s3Cfg := s3.NewFromConfig(*cfg)
		s3Client = &S3Client{
			client:   s3Cfg,
			uploader: manager.NewUploader(s3Cfg),
		}
	}
	return s3Client
}

/*
s3 Upload
503 maximum number of attempts 오류로 인한 재시도 로직 추가
참조: https://repost.aws/ko/knowledge-center/http-5xx-errors-s3
*/
func (client *S3Client) UploadS3V2(input UploadInput, retryCount int) *manager.UploadOutput {
	if retryCount > 5 {
		retryCount = 5
	}

	reader := strings.NewReader(input.Body)
	uploadParams := &s3.PutObjectInput{
		Bucket:      aws.String(input.Bucket),
		Key:         aws.String(input.Key),
		Body:        reader,
		ContentType: aws.String(input.ContentType),
	}

	result, err := client.uploader.Upload(context.TODO(), uploadParams)

	if err != nil &&
		strings.Contains(err.Error(), "S3: PutObject, exceeded maximum number of attempts, 1, https response") &&
		retryCount > 0 {
		time.Sleep(time.Millisecond * (10 * time.Duration(10-retryCount)))
		return client.UploadS3V2(input, retryCount-1)
	} else if err != nil {
		logging_tool.PrintErrorLog("UploadS3V2", err.Error())
		return nil
	}

	return result
}
