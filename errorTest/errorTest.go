package main

import "errors"

var (
	ErrNotFount   = errors.New("Resource Not Found")
	ErrBadRequest = errors.New("Bad Request")
	ErrorMap      = make(map[int]error)
)

type errorMessage struct {
	code    int
	message error
}

func GetError(code int) error {
	func() {
		ErrorMap[500] = ErrBadRequest
		ErrorMap[404] = ErrNotFount
	}()

	if code < 404 {
		panic(errors.New("Invalid code"))
	}
	return ErrorMap[code]
}
