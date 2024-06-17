# Error Codes Documentation

| STT | Error Code                    | Error Number | Description                                                              |
| --- | ----------------------------- | ------------ | ------------------------------------------------------------------------ |
| 1   | **ErrCodeInternalServer**     | 1000         | Indicates an unexpected internal server error.                           |
| 2   | **ErrCodeInvalidRequest**     | 1001         | Indicates the request parameters are invalid.                            |
| 3   | **ErrCodeUnauthorized**       | 1002         | Indicates the user is not authenticated.                                 |
| 4   | **ErrCodeForbidden**          | 1003         | Indicates the user does not have permission to access the resource.      |
| 5   | **ErrCodeNotFound**           | 1004         | Indicates the requested resource was not found.                          |
| 6   | **ErrCodeConflict**           | 1005         | Indicates there is a conflict with the current state of the resource.    |
| 7   | **ErrCodeTooManyRequests**    | 1006         | Indicates the user has sent too many requests in a given amount of time. |
| 8   | **ErrCodeServiceUnavailable** | 1007         | Indicates the service is currently unavailable or crashed.               |
| 9   | **ErrCodeHeaderNotExit**      | 1008         | Indicates the headers do not exist.                                      |
| 10  | **ErrorNotRead**              | 1010         | Indicates the request body was not read.                                 |

| STT | Error Code               | Error Number | Description                                                |
| --- | ------------------------ | ------------ | ---------------------------------------------------------- |
| 11  | **ErrCodeDBConnection**  | 2000         | Indicates a database connection error.                     |
| 12  | **ErrCodeDBQuery**       | 2001         | Indicates an error occurred while querying the database.   |
| 13  | **ErrCodeDBTransaction** | 2002         | Indicates an error occurred during a database transaction. |

| STT | Error Code                     | Error Number | Description                                           |
| --- | ------------------------------ | ------------ | ----------------------------------------------------- |
| 14  | **ErrCodeCacheConnection**     | 2100         | Indicates a cache connection error.                   |
| 15  | **ErrCodeCacheQuery**          | 2101         | Indicates an error occurred while querying the cache. |
| 16  | **ErrCodeCacheInvalidRequest** | 2102         | Indicates the request parameters are invalid.         |

| STT | Error Code                               | Error Number | Description                                                         |
| --- | ---------------------------------------- | ------------ | ------------------------------------------------------------------- |
| 17  | **ErrCodeValidation**                    | 3000         | Indicates there was a validation error with the request parameters. |
| 18  | **ErrCodeMissingField**                  | 3001         | Indicates a required field is missing.                              |
| 19  | **ErrCodeInvalidFormat**                 | 3002         | Indicates a field has an invalid format.                            |
| 20  | **ErrCodeContentType**                   | 3003         | Indicates the content type is invalid.                              |
| 21  | **ErrCookieInvalid**                     | 3004         | Indicates the cookie is invalid.                                    |
| 22  | **ErrorBodySizeTooLarge**                | 3005         | Indicates the request body size is too large.                       |
| 23  | **ErrPotentiallyDangerousInputDetected** | 3006         | Indicates dangerous input detected.                                 |

| STT | Error Code                  | Error Number | Description                                                 |
| --- | --------------------------- | ------------ | ----------------------------------------------------------- |
| 24  | **ErrCodeAuthTokenExpired** | 4000         | Indicates the authentication token has expired.             |
| 25  | **ErrCodeAuthTokenInvalid** | 4001         | Indicates the authentication token is invalid.              |
| 26  | **ErrCodePermissionDenied** | 4002         | Indicates the user does not have the necessary permissions. |
| 27  | **ErrCodeLoginFailed**      | 4003         | Indicates the login attempt failed.                         |

| STT | Error Code                   | Error Number | Description                                      |
| --- | ---------------------------- | ------------ | ------------------------------------------------ |
| 28  | **ErrCodeResourceExhausted** | 5000         | Indicates the resource has been exhausted.       |
| 29  | **ErrCodeResourceNotFound**  | 5001         | Indicates the resource was not found.            |
| 30  | **ErrCodeResourceConflict**  | 5002         | Indicates there is a conflict with the resource. |
| 31  | **ErrCodeResourceLocked**    | 5003         | Indicates the resource is locked.                |
| 32  | **ErrCSRFTokenInvalid**      | 5004         | Indicates the CSRF token is invalid.             |
| 33  | **ErrIpBlackList**           | 5005         | Indicates the IP is blacklisted.                 |
| 34  | **ErrPathTraversal**         | 5006         | Indicates a path traversal attack.               |

| STT | Error Code                            | Error Number | Description                                              |
| --- | ------------------------------------- | ------------ | -------------------------------------------------------- |
| 35  | **ErrCodeExternalService**            | 6000         | Indicates an error occurred with an external service.    |
| 36  | **ErrCodeExternalTimeout**            | 6001         | Indicates an external service request timed out.         |
| 37  | **ErrCodeExternalServiceUnavailable** | 6002         | Indicates the external service is currently unavailable. |

| STT | Error Code                    | Error Number | Description                            |
| --- | ----------------------------- | ------------ | -------------------------------------- |
| 38  | **ErrCodeNetworkUnavailable** | 7000         | Indicates the network is unavailable.  |
| 39  | **ErrCodeNetworkTimeout**     | 7001         | Indicates a network request timed out. |
| 40  | **ErrCodeNetworkError**       | 7002         | Indicates a general network error.     |

| STT | Error Code                      | Error Number | Description                                         |
| --- | ------------------------------- | ------------ | --------------------------------------------------- |
| 41  | **ErrCodeFileNotFound**         | 8000         | Indicates the specified file was not found.         |
| 42  | **ErrCodeFilePermissionDenied** | 8001         | Indicates permission to access the file was denied. |
| 43  | **ErrCodeFileUploadFailed**     | 8002         | Indicates the file upload failed.                   |
| 44  | **ErrCodeFileTooLarge**         | 8003         | Indicates the file size exceeds the allowed limit.  |

| STT | Error Code               | Error Number | Description                            |
| --- | ------------------------ | ------------ | -------------------------------------- |
| 45  | **ErrCodeInvalidInput**  | 9000         | Indicates the user input is invalid.   |
| 46  | **ErrCodeInputTooLong**  | 9001         | Indicates the user input is too long.  |
| 47  | **ErrCodeInputTooShort** | 9002         | Indicates the user input is too short. |

| STT | Error Code                     | Error Number | Description                                                 |
| --- | ------------------------------ | ------------ | ----------------------------------------------------------- |
| 48  | **ErrCodePaymentFailed**       | 10000        | Indicates the payment process failed.                       |
| 49  | **ErrCodeInsufficientFunds**   | 10001        | Indicates there are insufficient funds for the transaction. |
| 50  | **ErrCodePaymentGatewayError** | 10002        | Indicates an error occurred with the payment gateway.       |

| STT | Error Code                   | Error Number | Description                                |
| --- | ---------------------------- | ------------ | ------------------------------------------ |
| 51  | **ErrCodeSystemOverload**    | 11000        | Indicates the system is overloaded.        |
| 52  | **ErrCodeSystemMaintenance** | 11001        | Indicates the system is under maintenance. |
| 53  | **ErrCodeSystemError**       | 11002        | Indicates a general system error.          |

| STT | Error Code                            | Error Number | Description                                              |
| --- | ------------------------------------- | ------------ | -------------------------------------------------------- |
| 54  | **ErrUserNotExit**                    | 12000        | Indicates the account does not exist in the database.    |
| 55  | **ErrUserDuplicateEmail**             | 12001        | Indicates the email already exists in the users table.   |
| 56  | **ErrUserNotExitEmail**               | 12002        | Indicates the email does not exist in the database.      |
| 57  | **ErrorUserNotExitUsername**          | 12003        | Indicates the username does not exist in the database.   |
| 58  | **ErrorUserPhoneNotExit**             | 12004        | Indicates the phone does not exist in the database.      |
| 59  | **ErrUserNotActive**                  | 12005        | Indicates the account is not active.                     |
| 60  | **ErrTwoFactorDisabled**              | 12006        | Indicates two-factor authentication is disabled.         |
| 61  | **ErrTwoFactorInvalid**               | 12007        | Indicates two-factor authentication is invalid.          |
| 62  | **ErrTwoFactorUnauthorized**          | 12008        | Indicates two-factor authentication is unauthorized.     |
| 63  | **ErrorUsernameInvalid**              | 12009        | Indicates the username is invalid.                       |
| 64  | **ErrorUserPhoneInvalid**             | 12010        | Indicates the phone number is invalid.                   |
| 65  | **ErrCodeExternalServiceUnavailable** | 6002         | Indicates the external service is currently unavailable. |
| 66  | **ErrorUserEmailInvalid**             | 12011        | Indicates the email is invalid.                          |

| STT | Error Code               | Error Number | Description                          |
| --- | ------------------------ | ------------ | ------------------------------------ |
| 67  | **ErrCodeDeviceNotExit** | 13000        | Indicates the device does not exist. |

| STT | Error Code                         | Error Number | Description                                     |
| --- | ---------------------------------- | ------------ | ----------------------------------------------- |
| 68  | **ErrorVerificationCodeNotExit**   | 14000        | Indicates the verification code does not exist. |
| 69  | **ErrorVerificationCodeExpired**   | 14001        | Indicates the verification code has expired.    |
| 70  | **ErrorVerificationCodeInvalid**   | 14002        | Indicates the verification code is invalid.     |
| 71  | **ErrorVerificationCodeDuplicate** | 14003        | Indicates the verification code is duplicate.   |

| STT | Error Code                 | Error Number | Description                               |
| --- | -------------------------- | ------------ | ----------------------------------------- |
| 72  | **ErrorPasswordNotExit**   | 15000        | Indicates the password does not exist.    |
| 73  | **ErrorPasswordNotMatch**  | 15001        | Indicates the password does not match.    |
| 74  | **ErrorPasswordNotUpdate** | 15002        | Indicates the password was not updated.   |
| 75  | **ErrorEncryptPassword**   | 15003        | Indicates the password was not encrypted. |
| 76  | **ErrorPassWeak**          | 15004        | Indicates the password is weak.           |
| 77  | **ErrorPasswordIsOld**     | 15005        | Indicates the password is old.            |

| STT | Error Code          | Error Number | Description                       |
| --- | ------------------- | ------------ | --------------------------------- |
| 78  | **ErrorOTPNotExit** | 16000        | Indicates the OTP does not exist. |
| 79  | **ErrorOTPExpired** | 16001        | Indicates the OTP has expired.    |
| 80  | **ErrorOTPInvalid** | 16002        | Indicates the OTP is invalid.     |
