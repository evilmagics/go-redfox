package redfox

import "github.com/evilmagics/go-redfox/lib"

// Exception is a generic interface for handling and transforming error information
// with support for error codes, messages, metadata, and additional context.
// The generic type T allows for flexible error code representation while maintaining type safety.
//
// [Recommended] Always clone the exception before modifying it to avoid unintended side effects.
type Exception[T comparable] interface {
	// Error returns a string representation of the exception, typically combining the error code and message.
	Error() string

	// StatusCode returns the HTTP status code associated with the exception.
	StatusCode() int

	// Message returns the detailed error message of the exception.
	Message() string

	// DisplayMessage returns a user-friendly error message suitable for display to end-users.
	DisplayMessage() string

	// ErrType returns the type or category of the exception.
	ErrType() string

	// ErrCode returns the unique error code identifying the specific exception.
	ErrCode() T

	// Metadata returns additional contextual information about the exception as a map.
	Metadata() map[string]interface{}

	// StackTrace returns the call stack trace at the point where the exception was created.
	StackTrace() []string

	// Reason returns any underlying reason or cause associated with the exception.
	Reason() interface{}

	// Base returns the original error that triggered this exception, if any.
	Base() error

	// WithErrType sets a new error type for the exception and returns a modified copy.
	WithErrType(errType string) Exception[T]

	// WithDisplayMessage sets a new display message for the exception and returns a modified copy.
	WithDisplayMessage(message string) Exception[T]

	// WithMetadata adds or updates metadata for the exception and returns a modified copy.
	WithMetadata(data map[string]interface{}) Exception[T]

	// WithReason sets a new reason for the exception and returns a modified copy.
	WithReason(reason interface{}) Exception[T]

	// WithBase sets a new base error for the exception and returns a modified copy.
	WithBase(err error) Exception[T]

	// WithStatusCode sets a new HTTP status code for the exception and returns a modified copy.
	WithStatusCode(statusCode int) Exception[T]

	// Clone creates an exact copy of the current exception.
	Clone() Exception[T]

	// C simple way to clone an exception
	C() Exception[T]

	// View converts the exception to a serializable ExceptionView representation.
	View() ExceptionView[T]
}

// ExceptionView is a serializable representation of an Exception, designed for easy JSON marshaling.
// It contains all the key details of an exception, including error code, message, type, and additional context.
// The generic type T allows for flexible error code representation while maintaining type safety.
type ExceptionView[T comparable] struct {
	Base           string                 `json:"base,omitempty"`
	Message        string                 `json:"message"`
	DisplayMessage string                 `json:"displayMessage,omitempty"`
	ErrType        string                 `json:"errorType,omitempty"`
	ErrCode        T                      `json:"errorCode"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	StackTrace     []string               `json:"stackTrace,omitempty"`
	Reason         interface{}            `json:"reason,omitempty"`
	StatusCode     int                    `json:"statusCode,omitempty"`
}

// exception is a generic struct implementing the Exception interface, representing a detailed error
// with comprehensive error information including base error, message, error type, code, stack trace,
// reason, metadata, and HTTP status code. The generic type T allows for flexible error code representation.
type exception[T comparable] struct {
	base           error
	message        string
	displayMessage string
	errType        string
	errCode        T
	stackTrace     []string
	reason         interface{}
	metadata       map[string]interface{}
	statusCode     int
}

func (e exception[T]) Base() error {
	return e.base
}
func (e exception[T]) Error() string {
	return lib.Stringify(e.errCode) + ": " + e.message
}
func (e exception[T]) StatusCode() int {
	return e.statusCode
}
func (e exception[T]) Message() string {
	return e.message
}
func (e exception[T]) DisplayMessage() string {
	return e.displayMessage
}
func (e exception[T]) ErrType() string {
	return e.errType
}
func (e exception[T]) ErrCode() T {
	return e.errCode
}
func (e exception[T]) Metadata() map[string]interface{} {
	return e.metadata
}
func (e exception[T]) StackTrace() []string {
	return e.stackTrace
}
func (e exception[T]) Reason() interface{} {
	return e.reason
}
func (e *exception[T]) WithErrType(errType string) Exception[T] {
	e.errType = errType
	return e
}
func (e *exception[T]) WithDisplayMessage(message string) Exception[T] {
	e.displayMessage = message
	return e
}
func (e *exception[T]) WithMetadata(data map[string]interface{}) Exception[T] {
	e.metadata = data
	return e
}
func (e *exception[T]) WithReason(reason interface{}) Exception[T] {
	e.reason = reason
	return e
}
func (e *exception[T]) WithBase(err error) Exception[T] {
	e.base = err
	return e
}
func (e *exception[T]) WithStatusCode(statusCode int) Exception[T] {
	e.statusCode = statusCode
	return e
}
func (e *exception[T]) Clone() Exception[T] {
	return &exception[T]{
		base:           e.base,
		message:        e.message,
		displayMessage: e.displayMessage,
		errType:        e.errType,
		errCode:        e.errCode,
		stackTrace:     e.stackTrace,
		reason:         e.reason,
		metadata:       e.metadata,
		statusCode:     e.statusCode,
	}
}

func (e *exception[T]) C() Exception[T] {
	return e.Clone()
}

func (e exception[T]) View() ExceptionView[T] {
	b := ""
	if e.base != nil {
		b = e.base.Error()
	}

	return ExceptionView[T]{
		Base:           b,
		Message:        e.message,
		DisplayMessage: e.displayMessage,
		ErrType:        e.errType,
		ErrCode:        e.errCode,
		Metadata:       e.metadata,
		StackTrace:     e.stackTrace,
		Reason:         e.reason,
		StatusCode:     e.statusCode,
	}
}
