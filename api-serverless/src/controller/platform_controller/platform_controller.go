package platform_controller

import (
	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
	"github.com/aws/aws-lambda-go/events"
)

func RegisterPlatform(request common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : platform register
	return events.APIGatewayProxyResponse{}
}

func IssueToken(request common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : platform access token issue
	return events.APIGatewayProxyResponse{}
}

func UpdatePlatform(request common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : platform update
	return events.APIGatewayProxyResponse{}
}

func WithdrawPlatform(request common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : platform withdraw
	return events.APIGatewayProxyResponse{}
}

func GetOne(request common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : get platform data
	return events.APIGatewayProxyResponse{}
}
