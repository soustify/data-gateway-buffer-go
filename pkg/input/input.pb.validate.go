// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: input.proto

package input

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

// Validate checks the field values on PaginationRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PaginationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PaginationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PaginationRequestMultiError, or nil if none found.
func (m *PaginationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PaginationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Limit

	for idx, item := range m.GetFilter() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PaginationRequestValidationError{
						field:  fmt.Sprintf("Filter[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PaginationRequestValidationError{
						field:  fmt.Sprintf("Filter[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PaginationRequestValidationError{
					field:  fmt.Sprintf("Filter[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetOrderBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PaginationRequestValidationError{
					field:  "OrderBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PaginationRequestValidationError{
					field:  "OrderBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrderBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PaginationRequestValidationError{
				field:  "OrderBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PaginationRequestMultiError(errors)
	}

	return nil
}

// PaginationRequestMultiError is an error wrapping multiple validation errors
// returned by PaginationRequest.ValidateAll() if the designated constraints
// aren't met.
type PaginationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginationRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginationRequestMultiError) AllErrors() []error { return m }

// PaginationRequestValidationError is the validation error returned by
// PaginationRequest.Validate if the designated constraints aren't met.
type PaginationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationRequestValidationError) ErrorName() string {
	return "PaginationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PaginationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationRequestValidationError{}

// Validate checks the field values on UUIDRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UUIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UUIDRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UUIDRequestMultiError, or
// nil if none found.
func (m *UUIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UUIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	if len(errors) > 0 {
		return UUIDRequestMultiError(errors)
	}

	return nil
}

// UUIDRequestMultiError is an error wrapping multiple validation errors
// returned by UUIDRequest.ValidateAll() if the designated constraints aren't met.
type UUIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UUIDRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UUIDRequestMultiError) AllErrors() []error { return m }

// UUIDRequestValidationError is the validation error returned by
// UUIDRequest.Validate if the designated constraints aren't met.
type UUIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UUIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UUIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UUIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UUIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UUIDRequestValidationError) ErrorName() string { return "UUIDRequestValidationError" }

// Error satisfies the builtin error interface
func (e UUIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUUIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UUIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UUIDRequestValidationError{}

// Validate checks the field values on FilteredRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FilteredRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FilteredRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FilteredRequestMultiError, or nil if none found.
func (m *FilteredRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FilteredRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FilterContent

	// no validation rules for Operator

	for idx, item := range m.GetFilteredFields() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FilteredRequestValidationError{
						field:  fmt.Sprintf("FilteredFields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FilteredRequestValidationError{
						field:  fmt.Sprintf("FilteredFields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilteredRequestValidationError{
					field:  fmt.Sprintf("FilteredFields[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FilteredRequestMultiError(errors)
	}

	return nil
}

// FilteredRequestMultiError is an error wrapping multiple validation errors
// returned by FilteredRequest.ValidateAll() if the designated constraints
// aren't met.
type FilteredRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FilteredRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FilteredRequestMultiError) AllErrors() []error { return m }

// FilteredRequestValidationError is the validation error returned by
// FilteredRequest.Validate if the designated constraints aren't met.
type FilteredRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilteredRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilteredRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilteredRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilteredRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilteredRequestValidationError) ErrorName() string { return "FilteredRequestValidationError" }

// Error satisfies the builtin error interface
func (e FilteredRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilteredRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilteredRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilteredRequestValidationError{}

// Validate checks the field values on FilterFieldsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FilterFieldsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FilterFieldsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FilterFieldsRequestMultiError, or nil if none found.
func (m *FilterFieldsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FilterFieldsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Field

	// no validation rules for ComparisonKind

	if len(errors) > 0 {
		return FilterFieldsRequestMultiError(errors)
	}

	return nil
}

// FilterFieldsRequestMultiError is an error wrapping multiple validation
// errors returned by FilterFieldsRequest.ValidateAll() if the designated
// constraints aren't met.
type FilterFieldsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FilterFieldsRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FilterFieldsRequestMultiError) AllErrors() []error { return m }

// FilterFieldsRequestValidationError is the validation error returned by
// FilterFieldsRequest.Validate if the designated constraints aren't met.
type FilterFieldsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterFieldsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterFieldsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterFieldsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterFieldsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterFieldsRequestValidationError) ErrorName() string {
	return "FilterFieldsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FilterFieldsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterFieldsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterFieldsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterFieldsRequestValidationError{}

// Validate checks the field values on OrderByRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderByRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderByRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderByRequestMultiError,
// or nil if none found.
func (m *OrderByRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderByRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Field

	// no validation rules for Order

	if len(errors) > 0 {
		return OrderByRequestMultiError(errors)
	}

	return nil
}

// OrderByRequestMultiError is an error wrapping multiple validation errors
// returned by OrderByRequest.ValidateAll() if the designated constraints
// aren't met.
type OrderByRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderByRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderByRequestMultiError) AllErrors() []error { return m }

// OrderByRequestValidationError is the validation error returned by
// OrderByRequest.Validate if the designated constraints aren't met.
type OrderByRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderByRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderByRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderByRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderByRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderByRequestValidationError) ErrorName() string { return "OrderByRequestValidationError" }

// Error satisfies the builtin error interface
func (e OrderByRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderByRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderByRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderByRequestValidationError{}
