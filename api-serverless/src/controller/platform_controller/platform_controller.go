package platform_controller

import (
	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
	"github.com/aws/aws-lambda-go/events"
)

func RegisterPlatformController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : platform register
	return events.APIGatewayProxyResponse{}
}

func IssueTokenController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : platform access token issue
	return events.APIGatewayProxyResponse{}
}

func UpdatePlatformController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : platform update
	return events.APIGatewayProxyResponse{}
}

func WithdrawPlatformController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : platform withdraw
	return events.APIGatewayProxyResponse{}
}

func GetOneController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : get platform data
	return events.APIGatewayProxyResponse{}
}
