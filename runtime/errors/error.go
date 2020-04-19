package errors

import (
	"errors"

	"github.com/graphql-go/graphql/gqlerrors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidArguments = errors.New("runtime: invalid arguments")
	ErrWrongType        = errors.New("runtime: wrong type for Get method")
)

var (
	_ error                   = (*Error)(nil)
	_ gqlerrors.ExtendedError = (*Error)(nil)
)

func NewError(status *status.Status) error {
	return &Error{
		code:    status.Code(),
		message: status.Message(),
		details: status.Details(),
	}
}

type Error struct {
	code    codes.Code
	message string
	details []interface{}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Extensions() map[string]interface{} {
	if e.code == codes.OK {
		return nil
	}
	details := make([]interface{}, 0, len(e.details))
	for _, detail := range e.details {
		switch t := detail.(type) {
		case *errdetails.BadRequest:
			details = append(details, newBadRequest(t))
		case *errdetails.DebugInfo:
			details = append(details, newDebugInfo(t))
		case *errdetails.Help:
			details = append(details, newHelp(t))
		case *errdetails.LocalizedMessage:
			details = append(details, newLocalizedMessage(t))
		case *errdetails.PreconditionFailure:
			details = append(details, newPreconditionFailure(t))
		case *errdetails.QuotaFailure:
			details = append(details, newQuotaFailure(t))
		case *errdetails.RequestInfo:
			details = append(details, newRequestInfo(t))
		case *errdetails.ResourceInfo:
			details = append(details, newResourceInfo(t))
		case *errdetails.RetryInfo:
			retryInfo, _ := newRetryInfo(t)
			details = append(details, retryInfo)
		}
	}
	return map[string]interface{}{
		"code":    e.code,
		"details": details,
	}
}

func ToGraphQLError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	return NewError(st)
}
