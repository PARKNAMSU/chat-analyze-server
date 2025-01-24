package common_model

import (
	"chat-platform-api.com/chat-platform-api/src/tool/logging_tool"
	"github.com/aws/aws-lambda-go/events"
)

type GlobalParameter = map[string]any

type CustomAPIRequest struct {
	events.APIGatewayProxyRequest
	globalParameter GlobalParameter
}

func (c *CustomAPIRequest) GetHeader(key string) string {
	value, ok := c.Headers[key]
	if !ok {
		logging_tool.PrintErrorLog("GetHeader", "Not found Header [key]:["+key+"]")
		value = ""
	}
	return value
}

func (c *CustomAPIRequest) GetParameter(key string) any {
	if c.globalParameter == nil {
		c.globalParameter = make(map[string]any)
		return nil
	}
	data, ok := c.globalParameter[key]
	if !ok {
		return nil
	}
	return data
}

func (c *CustomAPIRequest) SetParameter(key string, data any) {
	if c.globalParameter == nil {
		c.globalParameter = make(map[string]any)
	}
	c.globalParameter[key] = data
}

func (c *CustomAPIRequest) SetParameters(data map[string]any) {
	if c.globalParameter == nil {
		c.globalParameter = make(map[string]any)
	}
	for key, value := range data {
		c.globalParameter[key] = value
	}
}
