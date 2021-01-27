package identity

import (
	"regexp"
)

type AWSAccessKeyID string

func (A AWSAccessKeyID) IsValid() bool {
	if len(string(A)) != 20 {
		return false
	}
	re := regexp.MustCompile("[A-Z0-9]{20}")
	return re.MatchString(string(A))
}

type AWSSecretAccessKey string

func (A AWSSecretAccessKey) IsValid() bool {
	if len(string(A)) != 40 {
		return false
	}
	re := regexp.MustCompile("[a-zA-Z0-9/+=]{40}")
	return re.MatchString(string(A))
}
