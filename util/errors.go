package util

type InvalidObjectError struct{}

func (e InvalidObjectError) Error() string {
	return "Invalid object"
}
