package redfox

// New create a new exception
func New[T comparable](errCode T, message string) Exception[T] {
	return &exception[T]{
		errCode:  errCode,
		message:  message,
		metadata: make(map[string]interface{}),
	}
}

// NewWithBase create a new exception with base error
// base error is the error that caused the exception
func NewWithBase[T comparable](errCode T, message string, base error) Exception[T] {
	return New(errCode, message).WithBase(base)
}

// NewAPI create a new exception with status code
// status code is the http status code
func NewForAPI[T comparable](errCode T, message string, statusCode int) Exception[T] {
	return New(errCode, message).WithStatusCode(statusCode)
}

// NewManager create a new exception manager
func NewManager[T comparable]() *Manager[T] {
	return &Manager[T]{
		cache: make(map[T]Exception[T]),
	}
}

// NewManagerStr create a new exception manager for string error codes
func NewManagerStr() *Manager[string] {
	return NewManager[string]()
}

// NewManagerInt create a new exception manager for int error codes
func NewManagerInt() *Manager[int] {
	return NewManager[int]()
}
