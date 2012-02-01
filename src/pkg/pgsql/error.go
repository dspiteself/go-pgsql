// Copyright 2010 Alexander Neumann. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pgsql

import (
	"fmt"
)

// PGError contains detailed error information received from a PostgreSQL backend.
//
// Many go-pgsql functions return an os.PGError value. In case of a backend error,
// a type assertion as shown below gives you a *pgsql.PGError with all details:
//
//	...
//	_, err := rs.FetchNext()
//	if err != nil {
//		if pgerr, ok := err.(*pgsql.PGError); ok {
//			// Do something with pgerr
//		}
//	}
//	...
type PGError struct {
	severity         string
	code             string
	message          string
	detail           string
	hint             string
	position         string
	internalPosition string
	internalQuery    string
	where            string
	file             string
	line             string
	routine          string
}

func (e *PGError) Severity() string {
	return e.severity
}

func (e *PGError) Code() string {
	return e.code
}

func (e *PGError) Message() string {
	return e.message
}

func (e *PGError) Detail() string {
	return e.detail
}

func (e *PGError) Hint() string {
	return e.hint
}

func (e *PGError) Position() string {
	return e.position
}

func (e *PGError) InternalPosition() string {
	return e.internalPosition
}

func (e *PGError) InternalQuery() string {
	return e.internalQuery
}

func (e *PGError) Where() string {
	return e.where
}

func (e *PGError) File() string {
	return e.file
}

func (e *PGError) Line() string {
	return e.line
}

func (e *PGError) Routine() string {
	return e.routine
}

func (e *PGError) String() string {
	return fmt.Sprintf(
		`Severity: %s
		Code: %s
		Message: %s
		Detail: %s
		Hint: %s
		Position: %s
		Internal Position: %s
		Internal Query: %s
		Where: %s
		File: %s
		Line: %s
		Routine: %s`,
		e.severity, e.code, e.message, e.detail, e.hint, e.position,
		e.internalPosition, e.internalQuery, e.where, e.file, e.line, e.routine)
}
func (e *PGError) Error() string {
	return e.String()
}