// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gotocompany/siren/v1/siren.proto

package sirenv1

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

// Validate checks the field values on Subscription with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Subscription) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Subscription with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubscriptionMultiError, or
// nil if none found.
func (m *Subscription) ValidateAll() error {
	return m.validate(true)
}

func (m *Subscription) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if !_Subscription_Urn_Pattern.MatchString(m.GetUrn()) {
		err := SubscriptionValidationError{
			field:  "Urn",
			reason: "value does not match regex pattern \"^[A-Za-z0-9_-]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Namespace

	for idx, item := range m.GetReceiversRelation() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SubscriptionValidationError{
						field:  fmt.Sprintf("ReceiversRelation[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SubscriptionValidationError{
						field:  fmt.Sprintf("ReceiversRelation[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubscriptionValidationError{
					field:  fmt.Sprintf("ReceiversRelation[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Match

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubscriptionValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubscriptionValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubscriptionValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubscriptionValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubscriptionValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubscriptionValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CreatedBy

	// no validation rules for UpdatedBy

	if len(errors) > 0 {
		return SubscriptionMultiError(errors)
	}

	return nil
}

// SubscriptionMultiError is an error wrapping multiple validation errors
// returned by Subscription.ValidateAll() if the designated constraints aren't met.
type SubscriptionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubscriptionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubscriptionMultiError) AllErrors() []error { return m }

// SubscriptionValidationError is the validation error returned by
// Subscription.Validate if the designated constraints aren't met.
type SubscriptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscriptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscriptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscriptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscriptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscriptionValidationError) ErrorName() string { return "SubscriptionValidationError" }

// Error satisfies the builtin error interface
func (e SubscriptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscription.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscriptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscriptionValidationError{}

var _Subscription_Urn_Pattern = regexp.MustCompile("^[A-Za-z0-9_-]+$")

// Validate checks the field values on ListSubscriptionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListSubscriptionsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSubscriptionsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSubscriptionsRequestMultiError, or nil if none found.
func (m *ListSubscriptionsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSubscriptionsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NamespaceId

	// no validation rules for Match

	// no validation rules for NotificationMatch

	// no validation rules for SilenceId

	// no validation rules for Metadata

	// no validation rules for ReceiverId

	// no validation rules for SubscriptionReceiverLabels

	if len(errors) > 0 {
		return ListSubscriptionsRequestMultiError(errors)
	}

	return nil
}

// ListSubscriptionsRequestMultiError is an error wrapping multiple validation
// errors returned by ListSubscriptionsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListSubscriptionsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSubscriptionsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSubscriptionsRequestMultiError) AllErrors() []error { return m }

// ListSubscriptionsRequestValidationError is the validation error returned by
// ListSubscriptionsRequest.Validate if the designated constraints aren't met.
type ListSubscriptionsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSubscriptionsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSubscriptionsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSubscriptionsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSubscriptionsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSubscriptionsRequestValidationError) ErrorName() string {
	return "ListSubscriptionsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListSubscriptionsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSubscriptionsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSubscriptionsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSubscriptionsRequestValidationError{}

// Validate checks the field values on ListSubscriptionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListSubscriptionsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSubscriptionsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSubscriptionsResponseMultiError, or nil if none found.
func (m *ListSubscriptionsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSubscriptionsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSubscriptions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListSubscriptionsResponseValidationError{
						field:  fmt.Sprintf("Subscriptions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListSubscriptionsResponseValidationError{
						field:  fmt.Sprintf("Subscriptions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSubscriptionsResponseValidationError{
					field:  fmt.Sprintf("Subscriptions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListSubscriptionsResponseMultiError(errors)
	}

	return nil
}

// ListSubscriptionsResponseMultiError is an error wrapping multiple validation
// errors returned by ListSubscriptionsResponse.ValidateAll() if the
// designated constraints aren't met.
type ListSubscriptionsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSubscriptionsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSubscriptionsResponseMultiError) AllErrors() []error { return m }

// ListSubscriptionsResponseValidationError is the validation error returned by
// ListSubscriptionsResponse.Validate if the designated constraints aren't met.
type ListSubscriptionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSubscriptionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSubscriptionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSubscriptionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSubscriptionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSubscriptionsResponseValidationError) ErrorName() string {
	return "ListSubscriptionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListSubscriptionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSubscriptionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSubscriptionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSubscriptionsResponseValidationError{}

// Validate checks the field values on SubscriptionReceiverRelation with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubscriptionReceiverRelation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubscriptionReceiverRelation with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubscriptionReceiverRelationMultiError, or nil if none found.
func (m *SubscriptionReceiverRelation) ValidateAll() error {
	return m.validate(true)
}

func (m *SubscriptionReceiverRelation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SubscriptionId

	// no validation rules for ReceiverId

	// no validation rules for Labels

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubscriptionReceiverRelationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubscriptionReceiverRelationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionReceiverRelationValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubscriptionReceiverRelationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubscriptionReceiverRelationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionReceiverRelationValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SubscriptionReceiverRelationMultiError(errors)
	}

	return nil
}

// SubscriptionReceiverRelationMultiError is an error wrapping multiple
// validation errors returned by SubscriptionReceiverRelation.ValidateAll() if
// the designated constraints aren't met.
type SubscriptionReceiverRelationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubscriptionReceiverRelationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubscriptionReceiverRelationMultiError) AllErrors() []error { return m }

// SubscriptionReceiverRelationValidationError is the validation error returned
// by SubscriptionReceiverRelation.Validate if the designated constraints
// aren't met.
type SubscriptionReceiverRelationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscriptionReceiverRelationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscriptionReceiverRelationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscriptionReceiverRelationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscriptionReceiverRelationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscriptionReceiverRelationValidationError) ErrorName() string {
	return "SubscriptionReceiverRelationValidationError"
}

// Error satisfies the builtin error interface
func (e SubscriptionReceiverRelationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscriptionReceiverRelation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscriptionReceiverRelationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscriptionReceiverRelationValidationError{}
