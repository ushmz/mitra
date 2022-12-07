package store

import "errors"

var (
	// ErrQueryBuildFailure shows that something went wrong while building SQL query
	ErrQueryBuildFailure = errors.New("Failed to build SQL query")

	// ErrDatabaseExecutionFailere shows that the DB operation failed
	ErrDatabaseExecutionFailere = errors.New("Failed to execute DB operation")

	// ErrNilReceiver means the method is called with Nil receiver
	ErrNilReceiver = errors.New("Called with Nil receiver")

	// ErrUIDAlreadyExists means it failed to signup because of duplicated UID
	ErrUIDAlreadyExists = errors.New("UID already exists")
)
