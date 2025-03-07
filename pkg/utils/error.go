package utils

import (
	"fmt"
)

type ErrorHandler func(error) error

func handlePanic(r any, errHandler ErrorHandler) error {
	if err, ok := r.(error); ok {
		return errHandler(err)
	}
	return errHandler(fmt.Errorf("panic occurred: %v", r))
}

func tryCatchHandler(fn func(), errHandler ErrorHandler) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = handlePanic(r, errHandler)
		}
	}()
	fn()
	return
}

func TryCatch[T any](fn func() T, errHandler ErrorHandler) (res T, err error) {
	err = tryCatchHandler(func() {
		res = fn()
	}, errHandler)
	return res, err
}

func TryCatchVoid(fn func(), errHandler ErrorHandler) error {
	return tryCatchHandler(fn, errHandler)
}

func Raise[T any](res T, err error) T {
	if err != nil {
		panic(err)
	}
	return res
}

func RaiseTwo[T any, U any](res T, res2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return res, res2
}

func RaiseVoid(err error) {
	if err != nil {
		panic(err)
	}
}

func RaiseVoidByError(err error, customError error) {
	if err != nil {
		panic(customError)
	}
}

func RaiseVoidByErrorHandler(err error, errHandler ErrorHandler) {
	if err != nil {
		panic(errHandler(err))
	}
}

func DefaultErrorHandler(err error) error {
	println("Error: ", err.Error())
	return err
}

func Throw(err error) {
	panic(err)
}
