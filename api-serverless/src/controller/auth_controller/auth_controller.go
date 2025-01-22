package auth_controller

import (
	"chat-platform-api.com/chat-platform-api/src/type/model/common_model"
	"github.com/aws/aws-lambda-go/events"
)

func SignUpController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : sign up
	return events.APIGatewayProxyResponse{}
}

func SignInController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : sign in
	return events.APIGatewayProxyResponse{}
}

func RevokeController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : token revoke
	return events.APIGatewayProxyResponse{}
}

func SendEmailController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : auth email send
	return events.APIGatewayProxyResponse{}
}

func AuthenticationController(request *common_model.CustomAPIRequest) events.APIGatewayProxyResponse {
	// todo : email auth
	return events.APIGatewayProxyResponse{}
}