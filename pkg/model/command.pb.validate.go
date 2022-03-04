// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/model/command.proto

package model

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

// Validate checks the field values on Command with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Command) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Command with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CommandMultiError, or nil if none found.
func (m *Command) ValidateAll() error {
	return m.validate(true)
}

func (m *Command) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := CommandValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPipedId()) < 1 {
		err := CommandValidationError{
			field:  "PipedId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ApplicationId

	// no validation rules for DeploymentId

	// no validation rules for StageId

	// no validation rules for Commander

	// no validation rules for ProjectId

	// no validation rules for Status

	// no validation rules for Metadata

	if m.GetHandledAt() < 0 {
		err := CommandValidationError{
			field:  "HandledAt",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Command_Type_name[int32(m.GetType())]; !ok {
		err := CommandValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSyncApplication()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "SyncApplication",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "SyncApplication",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSyncApplication()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommandValidationError{
				field:  "SyncApplication",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateApplicationConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "UpdateApplicationConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "UpdateApplicationConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateApplicationConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommandValidationError{
				field:  "UpdateApplicationConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCancelDeployment()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "CancelDeployment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "CancelDeployment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCancelDeployment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommandValidationError{
				field:  "CancelDeployment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetApproveStage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "ApproveStage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "ApproveStage",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetApproveStage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommandValidationError{
				field:  "ApproveStage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBuildPlanPreview()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "BuildPlanPreview",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "BuildPlanPreview",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBuildPlanPreview()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommandValidationError{
				field:  "BuildPlanPreview",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetChainSyncApplication()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "ChainSyncApplication",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommandValidationError{
					field:  "ChainSyncApplication",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetChainSyncApplication()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommandValidationError{
				field:  "ChainSyncApplication",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetCreatedAt() <= 0 {
		err := CommandValidationError{
			field:  "CreatedAt",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUpdatedAt() <= 0 {
		err := CommandValidationError{
			field:  "UpdatedAt",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CommandMultiError(errors)
	}

	return nil
}

// CommandMultiError is an error wrapping multiple validation errors returned
// by Command.ValidateAll() if the designated constraints aren't met.
type CommandMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommandMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommandMultiError) AllErrors() []error { return m }

// CommandValidationError is the validation error returned by Command.Validate
// if the designated constraints aren't met.
type CommandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommandValidationError) ErrorName() string { return "CommandValidationError" }

// Error satisfies the builtin error interface
func (e CommandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommandValidationError{}

// Validate checks the field values on Command_SyncApplication with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Command_SyncApplication) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Command_SyncApplication with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Command_SyncApplicationMultiError, or nil if none found.
func (m *Command_SyncApplication) ValidateAll() error {
	return m.validate(true)
}

func (m *Command_SyncApplication) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetApplicationId()) < 1 {
		err := Command_SyncApplicationValidationError{
			field:  "ApplicationId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for SyncStrategy

	if len(errors) > 0 {
		return Command_SyncApplicationMultiError(errors)
	}

	return nil
}

// Command_SyncApplicationMultiError is an error wrapping multiple validation
// errors returned by Command_SyncApplication.ValidateAll() if the designated
// constraints aren't met.
type Command_SyncApplicationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Command_SyncApplicationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Command_SyncApplicationMultiError) AllErrors() []error { return m }

// Command_SyncApplicationValidationError is the validation error returned by
// Command_SyncApplication.Validate if the designated constraints aren't met.
type Command_SyncApplicationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Command_SyncApplicationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Command_SyncApplicationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Command_SyncApplicationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Command_SyncApplicationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Command_SyncApplicationValidationError) ErrorName() string {
	return "Command_SyncApplicationValidationError"
}

// Error satisfies the builtin error interface
func (e Command_SyncApplicationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommand_SyncApplication.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Command_SyncApplicationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Command_SyncApplicationValidationError{}

// Validate checks the field values on Command_UpdateApplicationConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Command_UpdateApplicationConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Command_UpdateApplicationConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// Command_UpdateApplicationConfigMultiError, or nil if none found.
func (m *Command_UpdateApplicationConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *Command_UpdateApplicationConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetApplicationId()) < 1 {
		err := Command_UpdateApplicationConfigValidationError{
			field:  "ApplicationId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetConfigPath()) < 1 {
		err := Command_UpdateApplicationConfigValidationError{
			field:  "ConfigPath",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetConfig()) < 1 {
		err := Command_UpdateApplicationConfigValidationError{
			field:  "Config",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Command_UpdateApplicationConfigMultiError(errors)
	}

	return nil
}

// Command_UpdateApplicationConfigMultiError is an error wrapping multiple
// validation errors returned by Command_UpdateApplicationConfig.ValidateAll()
// if the designated constraints aren't met.
type Command_UpdateApplicationConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Command_UpdateApplicationConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Command_UpdateApplicationConfigMultiError) AllErrors() []error { return m }

// Command_UpdateApplicationConfigValidationError is the validation error
// returned by Command_UpdateApplicationConfig.Validate if the designated
// constraints aren't met.
type Command_UpdateApplicationConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Command_UpdateApplicationConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Command_UpdateApplicationConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Command_UpdateApplicationConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Command_UpdateApplicationConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Command_UpdateApplicationConfigValidationError) ErrorName() string {
	return "Command_UpdateApplicationConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Command_UpdateApplicationConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommand_UpdateApplicationConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Command_UpdateApplicationConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Command_UpdateApplicationConfigValidationError{}

// Validate checks the field values on Command_CancelDeployment with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Command_CancelDeployment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Command_CancelDeployment with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Command_CancelDeploymentMultiError, or nil if none found.
func (m *Command_CancelDeployment) ValidateAll() error {
	return m.validate(true)
}

func (m *Command_CancelDeployment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDeploymentId()) < 1 {
		err := Command_CancelDeploymentValidationError{
			field:  "DeploymentId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ForceRollback

	// no validation rules for ForceNoRollback

	if len(errors) > 0 {
		return Command_CancelDeploymentMultiError(errors)
	}

	return nil
}

// Command_CancelDeploymentMultiError is an error wrapping multiple validation
// errors returned by Command_CancelDeployment.ValidateAll() if the designated
// constraints aren't met.
type Command_CancelDeploymentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Command_CancelDeploymentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Command_CancelDeploymentMultiError) AllErrors() []error { return m }

// Command_CancelDeploymentValidationError is the validation error returned by
// Command_CancelDeployment.Validate if the designated constraints aren't met.
type Command_CancelDeploymentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Command_CancelDeploymentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Command_CancelDeploymentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Command_CancelDeploymentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Command_CancelDeploymentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Command_CancelDeploymentValidationError) ErrorName() string {
	return "Command_CancelDeploymentValidationError"
}

// Error satisfies the builtin error interface
func (e Command_CancelDeploymentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommand_CancelDeployment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Command_CancelDeploymentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Command_CancelDeploymentValidationError{}

// Validate checks the field values on Command_ApproveStage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Command_ApproveStage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Command_ApproveStage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Command_ApproveStageMultiError, or nil if none found.
func (m *Command_ApproveStage) ValidateAll() error {
	return m.validate(true)
}

func (m *Command_ApproveStage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDeploymentId()) < 1 {
		err := Command_ApproveStageValidationError{
			field:  "DeploymentId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetStageId()) < 1 {
		err := Command_ApproveStageValidationError{
			field:  "StageId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Command_ApproveStageMultiError(errors)
	}

	return nil
}

// Command_ApproveStageMultiError is an error wrapping multiple validation
// errors returned by Command_ApproveStage.ValidateAll() if the designated
// constraints aren't met.
type Command_ApproveStageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Command_ApproveStageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Command_ApproveStageMultiError) AllErrors() []error { return m }

// Command_ApproveStageValidationError is the validation error returned by
// Command_ApproveStage.Validate if the designated constraints aren't met.
type Command_ApproveStageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Command_ApproveStageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Command_ApproveStageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Command_ApproveStageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Command_ApproveStageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Command_ApproveStageValidationError) ErrorName() string {
	return "Command_ApproveStageValidationError"
}

// Error satisfies the builtin error interface
func (e Command_ApproveStageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommand_ApproveStage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Command_ApproveStageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Command_ApproveStageValidationError{}

// Validate checks the field values on Command_BuildPlanPreview with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Command_BuildPlanPreview) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Command_BuildPlanPreview with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Command_BuildPlanPreviewMultiError, or nil if none found.
func (m *Command_BuildPlanPreview) ValidateAll() error {
	return m.validate(true)
}

func (m *Command_BuildPlanPreview) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetRepositoryId()) < 1 {
		err := Command_BuildPlanPreviewValidationError{
			field:  "RepositoryId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetHeadBranch()) < 1 {
		err := Command_BuildPlanPreviewValidationError{
			field:  "HeadBranch",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetHeadCommit()) < 1 {
		err := Command_BuildPlanPreviewValidationError{
			field:  "HeadCommit",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetBaseBranch()) < 1 {
		err := Command_BuildPlanPreviewValidationError{
			field:  "BaseBranch",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTimeout() < 0 {
		err := Command_BuildPlanPreviewValidationError{
			field:  "Timeout",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Command_BuildPlanPreviewMultiError(errors)
	}

	return nil
}

// Command_BuildPlanPreviewMultiError is an error wrapping multiple validation
// errors returned by Command_BuildPlanPreview.ValidateAll() if the designated
// constraints aren't met.
type Command_BuildPlanPreviewMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Command_BuildPlanPreviewMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Command_BuildPlanPreviewMultiError) AllErrors() []error { return m }

// Command_BuildPlanPreviewValidationError is the validation error returned by
// Command_BuildPlanPreview.Validate if the designated constraints aren't met.
type Command_BuildPlanPreviewValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Command_BuildPlanPreviewValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Command_BuildPlanPreviewValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Command_BuildPlanPreviewValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Command_BuildPlanPreviewValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Command_BuildPlanPreviewValidationError) ErrorName() string {
	return "Command_BuildPlanPreviewValidationError"
}

// Error satisfies the builtin error interface
func (e Command_BuildPlanPreviewValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommand_BuildPlanPreview.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Command_BuildPlanPreviewValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Command_BuildPlanPreviewValidationError{}

// Validate checks the field values on Command_ChainSyncApplication with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Command_ChainSyncApplication) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Command_ChainSyncApplication with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Command_ChainSyncApplicationMultiError, or nil if none found.
func (m *Command_ChainSyncApplication) ValidateAll() error {
	return m.validate(true)
}

func (m *Command_ChainSyncApplication) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DeploymentChainId

	// no validation rules for BlockIndex

	if utf8.RuneCountInString(m.GetApplicationId()) < 1 {
		err := Command_ChainSyncApplicationValidationError{
			field:  "ApplicationId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for SyncStrategy

	if len(errors) > 0 {
		return Command_ChainSyncApplicationMultiError(errors)
	}

	return nil
}

// Command_ChainSyncApplicationMultiError is an error wrapping multiple
// validation errors returned by Command_ChainSyncApplication.ValidateAll() if
// the designated constraints aren't met.
type Command_ChainSyncApplicationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Command_ChainSyncApplicationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Command_ChainSyncApplicationMultiError) AllErrors() []error { return m }

// Command_ChainSyncApplicationValidationError is the validation error returned
// by Command_ChainSyncApplication.Validate if the designated constraints
// aren't met.
type Command_ChainSyncApplicationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Command_ChainSyncApplicationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Command_ChainSyncApplicationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Command_ChainSyncApplicationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Command_ChainSyncApplicationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Command_ChainSyncApplicationValidationError) ErrorName() string {
	return "Command_ChainSyncApplicationValidationError"
}

// Error satisfies the builtin error interface
func (e Command_ChainSyncApplicationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommand_ChainSyncApplication.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Command_ChainSyncApplicationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Command_ChainSyncApplicationValidationError{}