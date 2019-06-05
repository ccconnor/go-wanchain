// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

const (

	// ErrCodeBillExpirationException for service response error code
	// "BillExpirationException".
	//
	// The requested report expired. Update the date interval and try again.
	ErrCodeBillExpirationException = "BillExpirationException"

	// ErrCodeDataUnavailableException for service response error code
	// "DataUnavailableException".
	//
	// The requested data is unavailable.
	ErrCodeDataUnavailableException = "DataUnavailableException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The pagination token is invalid. Try again without a pagination token.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// You made too many calls in a short period of time. Try again later.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeRequestChangedException for service response error code
	// "RequestChangedException".
	//
	// Your request parameters changed between pages. Try again with the old parameters
	// or without a pagination token.
	ErrCodeRequestChangedException = "RequestChangedException"
)
