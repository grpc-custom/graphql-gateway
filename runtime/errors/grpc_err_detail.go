package errors

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type (
	BadRequest struct {
		Type            string            `json:"type"`
		FieldViolations []*FieldViolation `json:"fieldViolations"`
	}
	FieldViolation struct {
		Field       string `json:"field"`
		Description string `json:"description"`
	}
	DebugInfo struct {
		Type         string   `json:"type"`
		StackEntries []string `json:"stackEntries"`
	}
	Help struct {
		Type  string  `json:"type"`
		Links []*Link `json:"links"`
	}
	Link struct {
		Description string `json:"description"`
		URL         string `json:"url"`
	}
	LocalizedMessage struct {
		Locale  string `json:"locale"`
		Message string `json:"message"`
	}
	PreconditionFailure struct {
		Violations []*PreconditionFailureViolations `json:"violations"`
	}
	PreconditionFailureViolations struct {
		Type        string `json:"type"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	QuotaFailure struct {
		Violations []*QuotaFailureViolation `json:"violations"`
	}
	QuotaFailureViolation struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
	}
	RequestInfo struct {
		RequestID   string `json:"requestId"`
		ServingData string `json:"servingData"`
	}
	ResourceInfo struct {
		ResourceType string `json:"resourceType"`
		ResourceName string `json:"resourceName"`
		Owner        string `json:"owner"`
		Description  string `json:"description"`
	}
	RetryInfo struct {
		RetryDelay time.Duration `json:"retryDelay"`
	}
)

func newBadRequest(detail *errdetails.BadRequest) *BadRequest {
	ret := make([]*FieldViolation, len(detail.FieldViolations))
	for i, fieldViolation := range detail.FieldViolations {
		ret[i] = &FieldViolation{
			Field:       fieldViolation.Field,
			Description: fieldViolation.Description,
		}
	}
	return &BadRequest{
		Type:            "BAD_REQUEST",
		FieldViolations: ret,
	}
}

func newDebugInfo(detail *errdetails.DebugInfo) *DebugInfo {
	return &DebugInfo{
		Type:         "DEBUG_INFO",
		StackEntries: detail.StackEntries,
	}
}

func newHelp(detail *errdetails.Help) *Help {
	links := make([]*Link, len(detail.Links))
	for i, link := range detail.Links {
		links[i] = &Link{
			Description: link.Description,
			URL:         link.Url,
		}
	}
	return &Help{
		Links: links,
	}
}

func newLocalizedMessage(detail *errdetails.LocalizedMessage) *LocalizedMessage {
	return &LocalizedMessage{
		Locale:  detail.Locale,
		Message: detail.Message,
	}
}

func newPreconditionFailure(detail *errdetails.PreconditionFailure) *PreconditionFailure {
	violations := make([]*PreconditionFailureViolations, len(detail.Violations))
	for i, violation := range detail.Violations {
		violations[i] = &PreconditionFailureViolations{
			Type:        violation.Type,
			Subject:     violation.Subject,
			Description: violation.Description,
		}
	}
	return &PreconditionFailure{
		Violations: violations,
	}
}

func newQuotaFailure(detail *errdetails.QuotaFailure) *QuotaFailure {
	violations := make([]*QuotaFailureViolation, len(detail.Violations))
	for i, violation := range detail.Violations {
		violations[i] = &QuotaFailureViolation{
			Subject:     violation.Subject,
			Description: violation.Description,
		}
	}
	return &QuotaFailure{
		Violations: violations,
	}
}

func newRequestInfo(detail *errdetails.RequestInfo) *RequestInfo {
	return &RequestInfo{
		RequestID:   detail.RequestId,
		ServingData: detail.ServingData,
	}
}

func newResourceInfo(detail *errdetails.ResourceInfo) *ResourceInfo {
	return &ResourceInfo{
		ResourceType: detail.ResourceType,
		ResourceName: detail.ResourceName,
		Owner:        detail.Owner,
		Description:  detail.Description,
	}
}

func newRetryInfo(detail *errdetails.RetryInfo) (*RetryInfo, error) {
	retry, err := ptypes.Duration(detail.RetryDelay)
	return &RetryInfo{
		RetryDelay: retry,
	}, err
}
