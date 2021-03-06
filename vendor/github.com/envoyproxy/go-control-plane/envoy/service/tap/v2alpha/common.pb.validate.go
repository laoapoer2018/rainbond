// Code generated by protoc-gen-validate
// source: envoy/service/tap/v2alpha/common.proto
// DO NOT EDIT!!!

package envoy_service_tap_v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on TapConfig with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TapConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMatchConfig() == nil {
		return TapConfigValidationError{
			Field:  "MatchConfig",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMatchConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TapConfigValidationError{
				Field:  "MatchConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if m.GetOutputConfig() == nil {
		return TapConfigValidationError{
			Field:  "OutputConfig",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetOutputConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TapConfigValidationError{
				Field:  "OutputConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// TapConfigValidationError is the validation error returned by
// TapConfig.Validate if the designated constraints aren't met.
type TapConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e TapConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTapConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = TapConfigValidationError{}

// Validate checks the field values on MatchPredicate with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MatchPredicate) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Rule.(type) {

	case *MatchPredicate_OrMatch:

		if v, ok := interface{}(m.GetOrMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					Field:  "OrMatch",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *MatchPredicate_AndMatch:

		if v, ok := interface{}(m.GetAndMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					Field:  "AndMatch",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *MatchPredicate_NotMatch:

		if v, ok := interface{}(m.GetNotMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					Field:  "NotMatch",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *MatchPredicate_AnyMatch:

		if m.GetAnyMatch() != true {
			return MatchPredicateValidationError{
				Field:  "AnyMatch",
				Reason: "value must equal true",
			}
		}

	case *MatchPredicate_HttpRequestMatch:

		if v, ok := interface{}(m.GetHttpRequestMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					Field:  "HttpRequestMatch",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *MatchPredicate_HttpResponseMatch:

		if v, ok := interface{}(m.GetHttpResponseMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					Field:  "HttpResponseMatch",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return MatchPredicateValidationError{
			Field:  "Rule",
			Reason: "value is required",
		}

	}

	return nil
}

// MatchPredicateValidationError is the validation error returned by
// MatchPredicate.Validate if the designated constraints aren't met.
type MatchPredicateValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e MatchPredicateValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = MatchPredicateValidationError{}

// Validate checks the field values on HttpRequestMatch with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HttpRequestMatch) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpRequestMatchValidationError{
					Field:  fmt.Sprintf("Headers[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpRequestMatchValidationError is the validation error returned by
// HttpRequestMatch.Validate if the designated constraints aren't met.
type HttpRequestMatchValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpRequestMatchValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpRequestMatch.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpRequestMatchValidationError{}

// Validate checks the field values on HttpResponseMatch with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HttpResponseMatch) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpResponseMatchValidationError{
					Field:  fmt.Sprintf("Headers[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpResponseMatchValidationError is the validation error returned by
// HttpResponseMatch.Validate if the designated constraints aren't met.
type HttpResponseMatchValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpResponseMatchValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpResponseMatch.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpResponseMatchValidationError{}

// Validate checks the field values on OutputConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OutputConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetSinks()) != 1 {
		return OutputConfigValidationError{
			Field:  "Sinks",
			Reason: "value must contain exactly 1 item(s)",
		}
	}

	for idx, item := range m.GetSinks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OutputConfigValidationError{
					Field:  fmt.Sprintf("Sinks[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// OutputConfigValidationError is the validation error returned by
// OutputConfig.Validate if the designated constraints aren't met.
type OutputConfigValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e OutputConfigValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutputConfig.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = OutputConfigValidationError{}

// Validate checks the field values on OutputSink with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OutputSink) Validate() error {
	if m == nil {
		return nil
	}

	switch m.OutputSinkType.(type) {

	case *OutputSink_StreamingAdmin:

		if v, ok := interface{}(m.GetStreamingAdmin()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OutputSinkValidationError{
					Field:  "StreamingAdmin",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return OutputSinkValidationError{
			Field:  "OutputSinkType",
			Reason: "value is required",
		}

	}

	return nil
}

// OutputSinkValidationError is the validation error returned by
// OutputSink.Validate if the designated constraints aren't met.
type OutputSinkValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e OutputSinkValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutputSink.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = OutputSinkValidationError{}

// Validate checks the field values on StreamingAdminSink with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamingAdminSink) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StreamingAdminSinkValidationError is the validation error returned by
// StreamingAdminSink.Validate if the designated constraints aren't met.
type StreamingAdminSinkValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StreamingAdminSinkValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamingAdminSink.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StreamingAdminSinkValidationError{}

// Validate checks the field values on MatchPredicate_MatchSet with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MatchPredicate_MatchSet) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRules()) < 2 {
		return MatchPredicate_MatchSetValidationError{
			Field:  "Rules",
			Reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicate_MatchSetValidationError{
					Field:  fmt.Sprintf("Rules[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// MatchPredicate_MatchSetValidationError is the validation error returned by
// MatchPredicate_MatchSet.Validate if the designated constraints aren't met.
type MatchPredicate_MatchSetValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e MatchPredicate_MatchSetValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate_MatchSet.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = MatchPredicate_MatchSetValidationError{}
