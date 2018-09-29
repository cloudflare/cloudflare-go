package cloudflare

// Error messages
const (
	errEmptyCredentials     = "invalid credentials: key & email must not be empty"
	errMakeRequestError     = "error from makeRequest"
	errUnmarshalError       = "error unmarshalling the JSON response"
	errRequestNotSuccessful = "error reported by API"
)
