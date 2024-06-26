// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gotocompany/siren/provider/v1beta1/provider.proto

package sirenproviderv1beta1

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

// Validate checks the field values on SyncRuntimeConfigRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SyncRuntimeConfigRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SyncRuntimeConfigRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SyncRuntimeConfigRequestMultiError, or nil if none found.
func (m *SyncRuntimeConfigRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SyncRuntimeConfigRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NamespaceId

	// no validation rules for NamespaceUrn

	if all {
		switch v := interface{}(m.GetProvider()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SyncRuntimeConfigRequestValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SyncRuntimeConfigRequestValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProvider()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SyncRuntimeConfigRequestValidationError{
				field:  "Provider",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Labels

	if len(errors) > 0 {
		return SyncRuntimeConfigRequestMultiError(errors)
	}

	return nil
}

// SyncRuntimeConfigRequestMultiError is an error wrapping multiple validation
// errors returned by SyncRuntimeConfigRequest.ValidateAll() if the designated
// constraints aren't met.
type SyncRuntimeConfigRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SyncRuntimeConfigRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SyncRuntimeConfigRequestMultiError) AllErrors() []error { return m }

// SyncRuntimeConfigRequestValidationError is the validation error returned by
// SyncRuntimeConfigRequest.Validate if the designated constraints aren't met.
type SyncRuntimeConfigRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SyncRuntimeConfigRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SyncRuntimeConfigRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SyncRuntimeConfigRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SyncRuntimeConfigRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SyncRuntimeConfigRequestValidationError) ErrorName() string {
	return "SyncRuntimeConfigRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SyncRuntimeConfigRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSyncRuntimeConfigRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SyncRuntimeConfigRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SyncRuntimeConfigRequestValidationError{}

// Validate checks the field values on SyncRuntimeConfigResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SyncRuntimeConfigResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SyncRuntimeConfigResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SyncRuntimeConfigResponseMultiError, or nil if none found.
func (m *SyncRuntimeConfigResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SyncRuntimeConfigResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Labels

	if len(errors) > 0 {
		return SyncRuntimeConfigResponseMultiError(errors)
	}

	return nil
}

// SyncRuntimeConfigResponseMultiError is an error wrapping multiple validation
// errors returned by SyncRuntimeConfigResponse.ValidateAll() if the
// designated constraints aren't met.
type SyncRuntimeConfigResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SyncRuntimeConfigResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SyncRuntimeConfigResponseMultiError) AllErrors() []error { return m }

// SyncRuntimeConfigResponseValidationError is the validation error returned by
// SyncRuntimeConfigResponse.Validate if the designated constraints aren't met.
type SyncRuntimeConfigResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SyncRuntimeConfigResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SyncRuntimeConfigResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SyncRuntimeConfigResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SyncRuntimeConfigResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SyncRuntimeConfigResponseValidationError) ErrorName() string {
	return "SyncRuntimeConfigResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SyncRuntimeConfigResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSyncRuntimeConfigResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SyncRuntimeConfigResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SyncRuntimeConfigResponseValidationError{}

// Validate checks the field values on UpsertRuleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpsertRuleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpsertRuleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpsertRuleRequestMultiError, or nil if none found.
func (m *UpsertRuleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpsertRuleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNamespace()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpsertRuleRequestValidationError{
					field:  "Namespace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpsertRuleRequestValidationError{
					field:  "Namespace",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNamespace()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpsertRuleRequestValidationError{
				field:  "Namespace",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetProvider()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpsertRuleRequestValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpsertRuleRequestValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProvider()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpsertRuleRequestValidationError{
				field:  "Provider",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRule()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpsertRuleRequestValidationError{
					field:  "Rule",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpsertRuleRequestValidationError{
					field:  "Rule",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRule()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpsertRuleRequestValidationError{
				field:  "Rule",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTemplate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpsertRuleRequestValidationError{
					field:  "Template",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpsertRuleRequestValidationError{
					field:  "Template",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTemplate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpsertRuleRequestValidationError{
				field:  "Template",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpsertRuleRequestMultiError(errors)
	}

	return nil
}

// UpsertRuleRequestMultiError is an error wrapping multiple validation errors
// returned by UpsertRuleRequest.ValidateAll() if the designated constraints
// aren't met.
type UpsertRuleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpsertRuleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpsertRuleRequestMultiError) AllErrors() []error { return m }

// UpsertRuleRequestValidationError is the validation error returned by
// UpsertRuleRequest.Validate if the designated constraints aren't met.
type UpsertRuleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertRuleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertRuleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertRuleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertRuleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertRuleRequestValidationError) ErrorName() string {
	return "UpsertRuleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpsertRuleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertRuleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertRuleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertRuleRequestValidationError{}

// Validate checks the field values on UpsertRuleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpsertRuleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpsertRuleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpsertRuleResponseMultiError, or nil if none found.
func (m *UpsertRuleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpsertRuleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpsertRuleResponseMultiError(errors)
	}

	return nil
}

// UpsertRuleResponseMultiError is an error wrapping multiple validation errors
// returned by UpsertRuleResponse.ValidateAll() if the designated constraints
// aren't met.
type UpsertRuleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpsertRuleResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpsertRuleResponseMultiError) AllErrors() []error { return m }

// UpsertRuleResponseValidationError is the validation error returned by
// UpsertRuleResponse.Validate if the designated constraints aren't met.
type UpsertRuleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpsertRuleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpsertRuleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpsertRuleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpsertRuleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpsertRuleResponseValidationError) ErrorName() string {
	return "UpsertRuleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpsertRuleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpsertRuleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpsertRuleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpsertRuleResponseValidationError{}

// Validate checks the field values on SetConfigRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SetConfigRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetConfigRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetConfigRequestMultiError, or nil if none found.
func (m *SetConfigRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SetConfigRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ConfigRaw

	if len(errors) > 0 {
		return SetConfigRequestMultiError(errors)
	}

	return nil
}

// SetConfigRequestMultiError is an error wrapping multiple validation errors
// returned by SetConfigRequest.ValidateAll() if the designated constraints
// aren't met.
type SetConfigRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetConfigRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetConfigRequestMultiError) AllErrors() []error { return m }

// SetConfigRequestValidationError is the validation error returned by
// SetConfigRequest.Validate if the designated constraints aren't met.
type SetConfigRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetConfigRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetConfigRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetConfigRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetConfigRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetConfigRequestValidationError) ErrorName() string { return "SetConfigRequestValidationError" }

// Error satisfies the builtin error interface
func (e SetConfigRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetConfigRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetConfigRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetConfigRequestValidationError{}

// Validate checks the field values on SetConfigResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SetConfigResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetConfigResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetConfigResponseMultiError, or nil if none found.
func (m *SetConfigResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SetConfigResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SetConfigResponseMultiError(errors)
	}

	return nil
}

// SetConfigResponseMultiError is an error wrapping multiple validation errors
// returned by SetConfigResponse.ValidateAll() if the designated constraints
// aren't met.
type SetConfigResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetConfigResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetConfigResponseMultiError) AllErrors() []error { return m }

// SetConfigResponseValidationError is the validation error returned by
// SetConfigResponse.Validate if the designated constraints aren't met.
type SetConfigResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetConfigResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetConfigResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetConfigResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetConfigResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetConfigResponseValidationError) ErrorName() string {
	return "SetConfigResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SetConfigResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetConfigResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetConfigResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetConfigResponseValidationError{}

// Validate checks the field values on TransformToAlertsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TransformToAlertsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TransformToAlertsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TransformToAlertsRequestMultiError, or nil if none found.
func (m *TransformToAlertsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TransformToAlertsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProviderId

	// no validation rules for NamespaceId

	if all {
		switch v := interface{}(m.GetBody()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TransformToAlertsRequestValidationError{
					field:  "Body",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TransformToAlertsRequestValidationError{
					field:  "Body",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBody()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TransformToAlertsRequestValidationError{
				field:  "Body",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TransformToAlertsRequestMultiError(errors)
	}

	return nil
}

// TransformToAlertsRequestMultiError is an error wrapping multiple validation
// errors returned by TransformToAlertsRequest.ValidateAll() if the designated
// constraints aren't met.
type TransformToAlertsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransformToAlertsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransformToAlertsRequestMultiError) AllErrors() []error { return m }

// TransformToAlertsRequestValidationError is the validation error returned by
// TransformToAlertsRequest.Validate if the designated constraints aren't met.
type TransformToAlertsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransformToAlertsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransformToAlertsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransformToAlertsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransformToAlertsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransformToAlertsRequestValidationError) ErrorName() string {
	return "TransformToAlertsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TransformToAlertsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransformToAlertsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransformToAlertsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransformToAlertsRequestValidationError{}

// Validate checks the field values on TransformToAlertsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TransformToAlertsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TransformToAlertsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TransformToAlertsResponseMultiError, or nil if none found.
func (m *TransformToAlertsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TransformToAlertsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAlerts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TransformToAlertsResponseValidationError{
						field:  fmt.Sprintf("Alerts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TransformToAlertsResponseValidationError{
						field:  fmt.Sprintf("Alerts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TransformToAlertsResponseValidationError{
					field:  fmt.Sprintf("Alerts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for FiringNum

	if len(errors) > 0 {
		return TransformToAlertsResponseMultiError(errors)
	}

	return nil
}

// TransformToAlertsResponseMultiError is an error wrapping multiple validation
// errors returned by TransformToAlertsResponse.ValidateAll() if the
// designated constraints aren't met.
type TransformToAlertsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransformToAlertsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransformToAlertsResponseMultiError) AllErrors() []error { return m }

// TransformToAlertsResponseValidationError is the validation error returned by
// TransformToAlertsResponse.Validate if the designated constraints aren't met.
type TransformToAlertsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransformToAlertsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransformToAlertsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransformToAlertsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransformToAlertsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransformToAlertsResponseValidationError) ErrorName() string {
	return "TransformToAlertsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e TransformToAlertsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransformToAlertsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransformToAlertsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransformToAlertsResponseValidationError{}
