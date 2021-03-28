package exception

type runtimeException struct {
	errType string
}

func (this runtimeException) Error() string {
	return this.errType
}

func RuntimeException(et string) error {
	return runtimeException{errType: et}
}
