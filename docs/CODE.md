# Error Codes Documentation

## **General Errors**

- **ErrCodeInternalServer (1000)**: Indicates an unexpected internal server error.
- **ErrCodeInvalidRequest (1001)**: Indicates the request parameters are invalid.
- **ErrCodeUnauthorized (1002)**: Indicates the user is not authenticated.
- **ErrCodeForbidden (1003)**: Indicates the user does not have permission to access the resource.
- **ErrCodeNotFound (1004)**: Indicates the requested resource was not found.
- **ErrCodeConflict (1005)**: Indicates there is a conflict with the current state of the resource.
- **ErrCodeTooManyRequests (1006)**: Indicates the user has sent too many requests in a given amount of time.
- **ErrCodeServiceUnavailable (1007)**: Indicates the service is currently unavailable or crashed.
- **ErrCodeHeaderNotExit (1008)**: Indicates the headers do not exist.
- **ErrorNotRead (1010)**: Indicates the request body was not read.

## **Database Errors**

- **ErrCodeDBConnection (2000)**: Indicates a database connection error.
- **ErrCodeDBQuery (2001)**: Indicates an error occurred while querying the database.
- **ErrCodeDBTransaction (2002)**: Indicates an error occurred during a database transaction.

## **Cache Errors**

- **ErrCodeCacheConnection (2100)**: Indicates a cache connection error.
- **ErrCodeCacheQuery (2101)**: Indicates an error occurred while querying the cache.
- **ErrCodeCacheInvalidRequest (2102)**: Indicates the request parameters are invalid.

## **Validation Errors**

- **ErrCodeValidation (3000)**: Indicates there was a validation error with the request parameters.
- **ErrCodeMissingField (3001)**: Indicates a required field is missing.
- **ErrCodeInvalidFormat (3002)**: Indicates a field has an invalid format.
- **ErrCodeContentType (3003)**: Indicates the content type is invalid.
- **ErrCookieInvalid (3004)**: Indicates the cookie is invalid.
- **ErrorBodySizeTooLarge (3005)**: Indicates the request body size is too large.
- **ErrPotentiallyDangerousInputDetected (3006)**: Indicates dangerous input detected.

## **Authentication and Authorization Errors**

- **ErrCodeAuthTokenExpired (4000)**: Indicates the authentication token has expired.
- **ErrCodeAuthTokenInvalid (4001)**: Indicates the authentication token is invalid.
- **ErrCodePermissionDenied (4002)**: Indicates the user does not have the necessary permissions.
- **ErrCodeLoginFailed (4003)**: Indicates the login attempt failed.

## **Resource Errors**

- **ErrCodeResourceExhausted (5000)**: Indicates the resource has been exhausted.
- **ErrCodeResourceNotFound (5001)**: Indicates the resource was not found.
- **ErrCodeResourceConflict (5002)**: Indicates there is a conflict with the resource.
- **ErrCodeResourceLocked (5003)**: Indicates the resource is locked.
- **ErrCSRFTokenInvalid (5004)**: Indicates the CSRF token is invalid.
- **ErrIpBlackList (5005)**: Indicates the IP is blacklisted.
- **ErrPathTraversal (5006)**: Indicates a path traversal attack.

## **External Service Errors**

- **ErrCodeExternalService (6000)**: Indicates an error occurred with an external service.
- **ErrCodeExternalTimeout (6001)**: Indicates an external service request timed out.
- **ErrCodeExternalServiceUnavailable (6002)**: Indicates the external service is currently unavailable.

## **Network Errors**

- **ErrCodeNetworkUnavailable (7000)**: Indicates the network is unavailable.
- **ErrCodeNetworkTimeout (7001)**: Indicates a network request timed out.
- **ErrCodeNetworkError (7002)**: Indicates a general network error.

## **File Errors**

- **ErrCodeFileNotFound (8000)**: Indicates the specified file was not found.
- **ErrCodeFilePermissionDenied (8001)**: Indicates permission to access the file was denied.
- **ErrCodeFileUploadFailed (8002)**: Indicates the file upload failed.
- **ErrCodeFileTooLarge (8003)**: Indicates the file size exceeds the allowed limit.

## **User Input Errors**

- **ErrCodeInvalidInput (9000)**: Indicates the user input is invalid.
- **ErrCodeInputTooLong (9001)**: Indicates the user input is too long.
- **ErrCodeInputTooShort (9002)**: Indicates the user input is too short.

## **Payment Errors**

- **ErrCodePaymentFailed (10000)**: Indicates the payment process failed.
- **ErrCodeInsufficientFunds (10001)**: Indicates there are insufficient funds for the transaction.
- **ErrCodePaymentGatewayError (10002)**: Indicates an error occurred with the payment gateway.

## **System Errors**

- **ErrCodeSystemOverload (11000)**: Indicates the system is overloaded.
- **ErrCodeSystemMaintenance (11001)**: Indicates the system is under maintenance.
- **ErrCodeSystemError (11002)**: Indicates a general system error.

## **User Table Errors**

- **ErrUserNotExit (12000)**: Indicates the account does not exist in the database.
- **ErrUserDuplicateEmail (12001)**: Indicates the email already exists in the users table.
- **ErrUserNotExitEmail (12002)**: Indicates the email does not exist in the database.
- **ErrorUserNotExitUsername (12003)**: Indicates the username does not exist in the database.
- **ErrorUserPhoneNotExit (12004)**: Indicates the phone does not exist in the database.
- **ErrUserNotActive (12005)**: Indicates the account is not active.
- **ErrTwoFactorDisabled (12006)**: Indicates two-factor authentication is disabled.
- **ErrTwoFactorInvalid (12007)**: Indicates two-factor authentication is invalid.
- **ErrTwoFactorUnauthorized (12008)**: Indicates two-factor authentication is unauthorized.
- **ErrorUsernameInvalid (12009)**: Indicates the username is invalid.
- **ErrorUserPhoneInvalid (12010)**: Indicates the phone number is invalid.
- **ErrorUserEmailInvalid (12011)**: Indicates the email is invalid.

## **Device Table Errors**

- **ErrCodeDeviceNotExit (12002)**: Indicates the device does not exist.

## **Verification Table Errors**

- **ErrorVerificationCodeNotExit (14000)**: Indicates the verification code does not exist.
- **ErrorVerificationCodeExpired (14001)**: Indicates the verification code has expired.
- **ErrorVerificationCodeInvalid (14002)**: Indicates the verification code is invalid.
- **ErrorVerificationCodeDuplicate (14003)**: Indicates the verification code is duplicate.

## **Password Table Errors**

- **ErrorPasswordNotExit (15000)**: Indicates the password does not exist.
- **ErrorPasswordNotMatch (15001)**: Indicates the password does not match.
- **ErrorPasswordNotUpdate (15002)**: Indicates the password was not updated.
- **ErrorEncryptPassword (15003)**: Indicates the password was not encrypted.
- **ErrorPassWeak (15004)**: Indicates the password is weak.
- **ErrorPasswordIsOld (15005)**: Indicates the password is old.

## **OTP Table Errors**

- **ErrorOTPNotExit (16000)**: Indicates the OTP does not exist.
- **ErrorOTPExpired (16001)**: Indicates the OTP has expired.
- **ErrorOTPInvalid (16002)**: Indicates the OTP is invalid.
