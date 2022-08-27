package carry

type Error struct {
	cause error
	message string
}

//Error act as error interface
func (e *Error) Error()string {
	if e.cause != nil {
		return fmt.Sprintf("%s cause:%v", e.message, e.cause)
	}
	return e.message
}

// Unwrap give the cause of this error
func (e *Error) Unwrap() error {
	return e.cause
}

// WithCause wrap err With *Error
func WithCause(err error) *Error {
	return &Error{cause: err}
}

type Wrapper func(err *Error) *Error

func WithCode(code int) Wrapper {
	return func (err *Error) *Error {
		
	}
}

func WithLocation(depth int,skip int) Wrapper{
	return func (err *Error) *Error {
		if err.HasLocation() {
			return ?
		}
		// capture new location : depth+skip+1
		return ?
	}
}

func ToRpcResponse(ctx context.Context, resp Response, fn interface{}) Wrapper {
	// save a log
	//1. where(who)
	//2. what(why)
	//3. when

	//translate to RpcResponse
}

func New(wrappers ...Wrapper) *Error{

}