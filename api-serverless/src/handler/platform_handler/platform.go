package main

import (
	"strings"

	"chat-platform-api.com/chat-platform-api/src/middleware"
	"chat-platform-api.com/chat-platform-api/src/middleware/api_middleware"
	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
	api_variable "chat-platform-api.com/chat-platform-api/src/variable/api_variable"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	urls := strings.Split(request.Path, "/")

	clientRequest := common_model.CustomAPIRequest{
		APIGatewayProxyRequest: request,
	}

	err, code := middleware.ExecMiddlewares(
		&clientRequest,
		api_middleware.GetAPIMiddleware(),
	)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: code,
		}, nil
	}

	var response events.APIGatewayProxyResponse

	switch urls[2] {
	case "register":
	case "issueToken":
	case "update":
	case "withdraw":
	case "getData":
	case "getStat":
	default:
		response = events.APIGatewayProxyResponse{
			Body:       api_variable.RESPONSE_INVALID_PATH,
			StatusCode: api_variable.STATUS_BAD_REQUEST,
		}
	}

	return response, nil
}

func main() {
	lambda.Start(handler)
}
