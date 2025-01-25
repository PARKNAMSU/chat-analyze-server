package api_variable

const (
	STATUS_OK       = 200
	STATUS_CREATED  = 201
	STATUS_ACCEPTED = 202

	STATUS_BAD_REQUEST        = 400
	STATUS_UNAUTHORIZED       = 401
	StatusPaymentRequired     = 402
	STATUS_FORBIDDEN          = 403
	STATUS_NOTFOUND           = 404
	STATUS_METHOD_NOT_ALLOWED = 405

	STATUS_INTERNAL_SERVER_ERROR = 500
)

const (
	RESPONSE_INVALID_PATH = "Invalid url path"
	INVALID_ROUTER        = "Invalid Router"
	INVALID_API_KEY       = "Invalid API Key"
	NOT_FOUND             = "Not Found"
	INVALID_AUTHORIZATION = "unAuthorization"
	PERMISSION_DENIED     = "Permission Denied"
	INTERNAL_SERVER_ERROR = "Internal Server Error"
)
