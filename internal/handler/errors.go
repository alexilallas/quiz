package handler

type ErrorScanAnswer struct{}

func (ErrorScanAnswer) Error() string {
	return "failed to scan answer"
}
