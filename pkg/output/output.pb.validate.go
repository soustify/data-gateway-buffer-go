// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: output.proto

package output

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CountResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CountResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CountResponseMultiError, or
// nil if none found.
func (m *CountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Count

	if len(errors) > 0 {
		return CountResponseMultiError(errors)
	}

	return nil
}

// CountResponseMultiError is an error wrapping multiple validation errors
// returned by CountResponse.ValidateAll() if the designated constraints
// aren't met.
type CountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CountResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CountResponseMultiError) AllErrors() []error { return m }

// CountResponseValidationError is the validation error returned by
// CountResponse.Validate if the designated constraints aren't met.
type CountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CountResponseValidationError) ErrorName() string { return "CountResponseValidationError" }

// Error satisfies the builtin error interface
func (e CountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CountResponseValidationError{}

// Validate checks the field values on StatusDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StatusDataResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatusDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StatusDataResponseMultiError, or nil if none found.
func (m *StatusDataResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StatusDataResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Message

	if len(errors) > 0 {
		return StatusDataResponseMultiError(errors)
	}

	return nil
}

// StatusDataResponseMultiError is an error wrapping multiple validation errors
// returned by StatusDataResponse.ValidateAll() if the designated constraints
// aren't met.
type StatusDataResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatusDataResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatusDataResponseMultiError) AllErrors() []error { return m }

// StatusDataResponseValidationError is the validation error returned by
// StatusDataResponse.Validate if the designated constraints aren't met.
type StatusDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusDataResponseValidationError) ErrorName() string {
	return "StatusDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StatusDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatusDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusDataResponseValidationError{}

// Validate checks the field values on PersistenceDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PersistenceDataResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PersistenceDataResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PersistenceDataResponseMultiError, or nil if none found.
func (m *PersistenceDataResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PersistenceDataResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return PersistenceDataResponseMultiError(errors)
	}

	return nil
}

// PersistenceDataResponseMultiError is an error wrapping multiple validation
// errors returned by PersistenceDataResponse.ValidateAll() if the designated
// constraints aren't met.
type PersistenceDataResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PersistenceDataResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PersistenceDataResponseMultiError) AllErrors() []error { return m }

// PersistenceDataResponseValidationError is the validation error returned by
// PersistenceDataResponse.Validate if the designated constraints aren't met.
type PersistenceDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PersistenceDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PersistenceDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PersistenceDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PersistenceDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PersistenceDataResponseValidationError) ErrorName() string {
	return "PersistenceDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PersistenceDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPersistenceDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PersistenceDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PersistenceDataResponseValidationError{}
