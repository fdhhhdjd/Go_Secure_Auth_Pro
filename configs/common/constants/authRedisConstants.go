package constants

import "time"

const (
	SpamKey                 = "spam_user"
	SpamKeyLogin            = "spam_user_login"
	SpamKeyLinkVerification = "spam_user_link_verification"
	SpamKeyForget           = "spam_user_forget"
)
const (
	RequestThreshold                 = 5
	RequestThresholdLinkVerification = 3
	RequestThresholdForget           = 2
)

const (
	InitialBlock   = 5 * time.Minute
	ExtendedBlock  = 30 * time.Minute
	ExpireDuration = 30 * time.Second
)