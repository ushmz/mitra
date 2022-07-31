package model

import "errors"

var (
	// ErrQueryBuildFailure shows that something goes wrong while building SQL query
	ErrQueryBuildFailure = errors.New("Failed to build SQL query")

	// ErrDatabaseExecutionFailere shows that the DB operation failed
	ErrDatabaseExecutionFailere = errors.New("Failed to execute DB operation")

	// ErrNilReceiver means the method is called with Nil receiver
	ErrNilReceiver = errors.New("Called with Nil receiver")

	// ErrBadRequest means the function is called with no required parameters
	ErrBadRequest = errors.New("No required parameter")

	// ErrNotFound means the requested resource is not found
	ErrNotFound = errors.New("Resource not found")

	// ErrInternal means the errors which details does not matter
	ErrInternal = errors.New("Somethig went wrong")
)
