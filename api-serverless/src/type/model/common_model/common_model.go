package common_model

import "github.com/aws/aws-lambda-go/events"

type GlobalParameter = map[string]any

type CustomAPIRequest struct {
	events.APIGatewayProxyRequest
	GlobalParameter GlobalParameter
}

type Middleware func (request *CustomAPIRequest) error