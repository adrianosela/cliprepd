package lib

// in an effort to keep the client's error handling consistent and for ease of
// testing, all error messages should be listed below and reused where possible
const (
	// input validation errors
	clientErrURLEmpty            = "url cannot be empty"
	clientErrAuthEmpty           = "auth credentials cannot be empty"
	clientErrObjectEmpty         = "object cannot be empty"
	clientErrObjectTypeEmpty     = "object type cannot be empty"
	clientErrViolationEmpty      = "violation cannot be empty"
	clientErrReputationNil       = "reputation cannot be nil"
	clientErrViolationRequestNil = "violation request cannot be nil"
	clientErrMarshal             = "could not marshal payload"
	// http client errors
	clientErrBuildRequest = "could not build http request"
	clientErrSendRequest  = "could not send http request"
	// http response payload errors
	clientErrReadResponse = "could not read response body"
	clientErrNon200       = "non 200 status code received"
	clientErrUnmarshal    = "could not unmarshal response body"
)
