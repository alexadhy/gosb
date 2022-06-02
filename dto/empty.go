package dto

// EmptyResponse is a universal response for any endpoints that didn't return anything
// but may return Error
type EmptyResponse struct {
	Error
}
