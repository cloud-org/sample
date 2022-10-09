package main

import "net/http"

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

func NewStatusError(err error) StatusError {
	return StatusError{
		Code: http.StatusBadRequest,
		Err:  err,
	}
}
