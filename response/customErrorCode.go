package response

const (
	//* General errors
	// ErrCodeInternalServer indicates an unexpected internal server error.
	ErrCodeInternalServer = 1000

	// ErrCodeInvalidRequest indicates the request parameters are invalid.
	ErrCodeInvalidRequest = 1001

	// ErrCodeUnauthorized indicates the user is not authenticated.
	ErrCodeUnauthorized = 1002

	// ErrCodeForbidden indicates the user does not have permission to access the resource.
	ErrCodeForbidden = 1003

	// ErrCodeNotFound indicates the requested resource was not found.
	ErrCodeNotFound = 1004

	// ErrCodeConflict indicates there is a conflict with the current state of the resource.
	ErrCodeConflict = 1005

	// ErrCodeTooManyRequests indicates the user has sent too many requests in a given amount of time.
	ErrCodeTooManyRequests = 1006

	// ErrCodeServiceUnavailable indicates the service is currently unavailable or crash
	ErrCodeServiceUnavailable = 1007

	// ErrCodeHeaderNotExit indicates the headers not exits
	ErrCodeHeaderNotExit = 1008

	// ErrorNotRead indicates the request body not read
	ErrorNotRead = 1010

	//* Database errors
	// ErrCodeDBConnection indicates a database connection error.
	ErrCodeDBConnection = 2000

	// ErrCodeDBQuery indicates an error occurred while querying the database.
	ErrCodeDBQuery = 2001

	// ErrCodeDBTransaction indicates an error occurred during a database transaction.
	ErrCodeDBTransaction = 2002

	//* Cache errors
	// ErrCodeCacheConnection indicates a cache connection error.
	ErrCodeCacheConnection = 2100

	// ErrCodeCacheQuery indicates an error occurred while querying the cache.
	ErrCodeCacheQuery = 2101

	// ErrCodeCacheInvalidRequest indicates the request parameters are invalid.
	ErrCodeCacheInvalidRequest = 2102

	//* Validation errors
	// ErrCodeValidation indicates there was a validation error with the request parameters.
	ErrCodeValidation = 3000

	// ErrCodeMissingField indicates a required field is missing.
	ErrCodeMissingField = 3001

	// ErrCodeInvalidFormat indicates a field has an invalid format.
	ErrCodeInvalidFormat = 3002

	// ErrCodeContentType indicates the content type is invalid.
	ErrCodeContentType = 3003

	// ErrCookieInvalid indicates the cookie is invalid.
	ErrCookieInvalid = 3004

	// ErrorBodySizeTooLarge indicates the request body size is too large.
	ErrorBodySizeTooLarge = 3005

	// ErrPotentiallyDangerousInputDetected dangerous input detected indicates the request body size is too large.
	ErrPotentiallyDangerousInputDetected = 3006

	//* Authentication and Authorization errors
	// ErrCodeAuthTokenExpired indicates the authentication token has expired.
	ErrCodeAuthTokenExpired = 4000

	// ErrCodeAuthTokenInvalid indicates the authentication token is invalid.
	ErrCodeAuthTokenInvalid = 4001

	// ErrCodePermissionDenied indicates the user does not have the necessary permissions.
	ErrCodePermissionDenied = 4002

	// ErrCodeLoginFailed indicates the login attempt failed.
	ErrCodeLoginFailed = 4003

	//* Resource errors
	// ErrCodeResourceExhausted indicates the resource has been exhausted.
	ErrCodeResourceExhausted = 5000

	// ErrCodeResourceNotFound indicates the resource was not found.
	ErrCodeResourceNotFound = 5001

	// ErrCodeResourceConflict indicates there is a conflict with the resource.
	ErrCodeResourceConflict = 5002

	// ErrCodeResourceLocked indicates the resource is locked.
	ErrCodeResourceLocked = 5003

	// ErrCSRFTokenInvalid indicates the CSRF token is invalid.
	ErrCSRFTokenInvalid = 5004

	// ErrIpBlackList indicates the ip is black list
	ErrIpBlackList = 5005

	// ErrPathTraversal indicates the path traversal attack
	ErrPathTraversal = 5006

	//* External service errors
	// ErrCodeExternalService indicates an error occurred with an external service.
	ErrCodeExternalService = 6000

	// ErrCodeExternalTimeout indicates an external service request timed out.
	ErrCodeExternalTimeout = 6001

	// ErrCodeExternalServiceUnavailable indicates the external service is currently unavailable.
	ErrCodeExternalServiceUnavailable = 6002

	//* Network errors
	// ErrCodeNetworkUnavailable indicates the network is unavailable.
	ErrCodeNetworkUnavailable = 7000

	// ErrCodeNetworkTimeout indicates a network request timed out.
	ErrCodeNetworkTimeout = 7001

	// ErrCodeNetworkError indicates a general network error.
	ErrCodeNetworkError = 7002

	//* File errors
	// ErrCodeFileNotFound indicates the specified file was not found.
	ErrCodeFileNotFound = 8000

	// ErrCodeFilePermissionDenied indicates permission to access the file was denied.
	ErrCodeFilePermissionDenied = 8001

	// ErrCodeFileUploadFailed indicates the file upload failed.
	ErrCodeFileUploadFailed = 8002

	// ErrCodeFileTooLarge indicates the file size exceeds the allowed limit.
	ErrCodeFileTooLarge = 8003

	//* User input errors
	// ErrCodeInvalidInput indicates the user input is invalid.
	ErrCodeInvalidInput = 9000

	// ErrCodeInputTooLong indicates the user input is too long.
	ErrCodeInputTooLong = 9001

	// ErrCodeInputTooShort indicates the user input is too short.
	ErrCodeInputTooShort = 9002

	//* Payment errors
	// ErrCodePaymentFailed indicates the payment process failed.
	ErrCodePaymentFailed = 10000

	// ErrCodeInsufficientFunds indicates there are insufficient funds for the transaction.
	ErrCodeInsufficientFunds = 10001

	// ErrCodePaymentGatewayError indicates an error occurred with the payment gateway.
	ErrCodePaymentGatewayError = 10002

	//* System errors
	// ErrCodeSystemOverload indicates the system is overloaded.
	ErrCodeSystemOverload = 11000

	// ErrCodeSystemMaintenance indicates the system is under maintenance.
	ErrCodeSystemMaintenance = 11001

	// ErrCodeSystemError indicates a general system error.
	ErrCodeSystemError = 11002
)
