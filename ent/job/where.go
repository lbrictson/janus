// Code generated by ent, DO NOT EDIT.

package job

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lbrictson/janus/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldDescription, v))
}

// CronSchedule applies equality check predicate on the "cron_schedule" field. It's identical to CronScheduleEQ.
func CronSchedule(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCronSchedule, v))
}

// ScheduleEnabled applies equality check predicate on the "schedule_enabled" field. It's identical to ScheduleEnabledEQ.
func ScheduleEnabled(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldScheduleEnabled, v))
}

// AllowConcurrentRuns applies equality check predicate on the "allow_concurrent_runs" field. It's identical to AllowConcurrentRunsEQ.
func AllowConcurrentRuns(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldAllowConcurrentRuns, v))
}

// RequiresFileUpload applies equality check predicate on the "requires_file_upload" field. It's identical to RequiresFileUploadEQ.
func RequiresFileUpload(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldRequiresFileUpload, v))
}

// AverageDurationMs applies equality check predicate on the "average_duration_ms" field. It's identical to AverageDurationMsEQ.
func AverageDurationMs(v int64) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldAverageDurationMs, v))
}

// TimeoutSeconds applies equality check predicate on the "timeout_seconds" field. It's identical to TimeoutSecondsEQ.
func TimeoutSeconds(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldTimeoutSeconds, v))
}

// LastEditTime applies equality check predicate on the "last_edit_time" field. It's identical to LastEditTimeEQ.
func LastEditTime(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldLastEditTime, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCreatedAt, v))
}

// LastRunTime applies equality check predicate on the "last_run_time" field. It's identical to LastRunTimeEQ.
func LastRunTime(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldLastRunTime, v))
}

// NextCronRunTime applies equality check predicate on the "next_cron_run_time" field. It's identical to NextCronRunTimeEQ.
func NextCronRunTime(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldNextCronRunTime, v))
}

// Script applies equality check predicate on the "script" field. It's identical to ScriptEQ.
func Script(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldScript, v))
}

// LastRunSuccess applies equality check predicate on the "last_run_success" field. It's identical to LastRunSuccessEQ.
func LastRunSuccess(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldLastRunSuccess, v))
}

// CreatedByAPI applies equality check predicate on the "created_by_api" field. It's identical to CreatedByAPIEQ.
func CreatedByAPI(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCreatedByAPI, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Job {
	return predicate.Job(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Job {
	return predicate.Job(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldDescription, v))
}

// CronScheduleEQ applies the EQ predicate on the "cron_schedule" field.
func CronScheduleEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCronSchedule, v))
}

// CronScheduleNEQ applies the NEQ predicate on the "cron_schedule" field.
func CronScheduleNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldCronSchedule, v))
}

// CronScheduleIn applies the In predicate on the "cron_schedule" field.
func CronScheduleIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldCronSchedule, vs...))
}

// CronScheduleNotIn applies the NotIn predicate on the "cron_schedule" field.
func CronScheduleNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldCronSchedule, vs...))
}

// CronScheduleGT applies the GT predicate on the "cron_schedule" field.
func CronScheduleGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldCronSchedule, v))
}

// CronScheduleGTE applies the GTE predicate on the "cron_schedule" field.
func CronScheduleGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldCronSchedule, v))
}

// CronScheduleLT applies the LT predicate on the "cron_schedule" field.
func CronScheduleLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldCronSchedule, v))
}

// CronScheduleLTE applies the LTE predicate on the "cron_schedule" field.
func CronScheduleLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldCronSchedule, v))
}

// CronScheduleContains applies the Contains predicate on the "cron_schedule" field.
func CronScheduleContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldCronSchedule, v))
}

// CronScheduleHasPrefix applies the HasPrefix predicate on the "cron_schedule" field.
func CronScheduleHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldCronSchedule, v))
}

// CronScheduleHasSuffix applies the HasSuffix predicate on the "cron_schedule" field.
func CronScheduleHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldCronSchedule, v))
}

// CronScheduleIsNil applies the IsNil predicate on the "cron_schedule" field.
func CronScheduleIsNil() predicate.Job {
	return predicate.Job(sql.FieldIsNull(FieldCronSchedule))
}

// CronScheduleNotNil applies the NotNil predicate on the "cron_schedule" field.
func CronScheduleNotNil() predicate.Job {
	return predicate.Job(sql.FieldNotNull(FieldCronSchedule))
}

// CronScheduleEqualFold applies the EqualFold predicate on the "cron_schedule" field.
func CronScheduleEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldCronSchedule, v))
}

// CronScheduleContainsFold applies the ContainsFold predicate on the "cron_schedule" field.
func CronScheduleContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldCronSchedule, v))
}

// ScheduleEnabledEQ applies the EQ predicate on the "schedule_enabled" field.
func ScheduleEnabledEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldScheduleEnabled, v))
}

// ScheduleEnabledNEQ applies the NEQ predicate on the "schedule_enabled" field.
func ScheduleEnabledNEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldScheduleEnabled, v))
}

// AllowConcurrentRunsEQ applies the EQ predicate on the "allow_concurrent_runs" field.
func AllowConcurrentRunsEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldAllowConcurrentRuns, v))
}

// AllowConcurrentRunsNEQ applies the NEQ predicate on the "allow_concurrent_runs" field.
func AllowConcurrentRunsNEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldAllowConcurrentRuns, v))
}

// ArgumentsIsNil applies the IsNil predicate on the "arguments" field.
func ArgumentsIsNil() predicate.Job {
	return predicate.Job(sql.FieldIsNull(FieldArguments))
}

// ArgumentsNotNil applies the NotNil predicate on the "arguments" field.
func ArgumentsNotNil() predicate.Job {
	return predicate.Job(sql.FieldNotNull(FieldArguments))
}

// RequiresFileUploadEQ applies the EQ predicate on the "requires_file_upload" field.
func RequiresFileUploadEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldRequiresFileUpload, v))
}

// RequiresFileUploadNEQ applies the NEQ predicate on the "requires_file_upload" field.
func RequiresFileUploadNEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldRequiresFileUpload, v))
}

// AverageDurationMsEQ applies the EQ predicate on the "average_duration_ms" field.
func AverageDurationMsEQ(v int64) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldAverageDurationMs, v))
}

// AverageDurationMsNEQ applies the NEQ predicate on the "average_duration_ms" field.
func AverageDurationMsNEQ(v int64) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldAverageDurationMs, v))
}

// AverageDurationMsIn applies the In predicate on the "average_duration_ms" field.
func AverageDurationMsIn(vs ...int64) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldAverageDurationMs, vs...))
}

// AverageDurationMsNotIn applies the NotIn predicate on the "average_duration_ms" field.
func AverageDurationMsNotIn(vs ...int64) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldAverageDurationMs, vs...))
}

// AverageDurationMsGT applies the GT predicate on the "average_duration_ms" field.
func AverageDurationMsGT(v int64) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldAverageDurationMs, v))
}

// AverageDurationMsGTE applies the GTE predicate on the "average_duration_ms" field.
func AverageDurationMsGTE(v int64) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldAverageDurationMs, v))
}

// AverageDurationMsLT applies the LT predicate on the "average_duration_ms" field.
func AverageDurationMsLT(v int64) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldAverageDurationMs, v))
}

// AverageDurationMsLTE applies the LTE predicate on the "average_duration_ms" field.
func AverageDurationMsLTE(v int64) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldAverageDurationMs, v))
}

// AverageDurationMsIsNil applies the IsNil predicate on the "average_duration_ms" field.
func AverageDurationMsIsNil() predicate.Job {
	return predicate.Job(sql.FieldIsNull(FieldAverageDurationMs))
}

// AverageDurationMsNotNil applies the NotNil predicate on the "average_duration_ms" field.
func AverageDurationMsNotNil() predicate.Job {
	return predicate.Job(sql.FieldNotNull(FieldAverageDurationMs))
}

// TimeoutSecondsEQ applies the EQ predicate on the "timeout_seconds" field.
func TimeoutSecondsEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldTimeoutSeconds, v))
}

// TimeoutSecondsNEQ applies the NEQ predicate on the "timeout_seconds" field.
func TimeoutSecondsNEQ(v int) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldTimeoutSeconds, v))
}

// TimeoutSecondsIn applies the In predicate on the "timeout_seconds" field.
func TimeoutSecondsIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldTimeoutSeconds, vs...))
}

// TimeoutSecondsNotIn applies the NotIn predicate on the "timeout_seconds" field.
func TimeoutSecondsNotIn(vs ...int) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldTimeoutSeconds, vs...))
}

// TimeoutSecondsGT applies the GT predicate on the "timeout_seconds" field.
func TimeoutSecondsGT(v int) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldTimeoutSeconds, v))
}

// TimeoutSecondsGTE applies the GTE predicate on the "timeout_seconds" field.
func TimeoutSecondsGTE(v int) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldTimeoutSeconds, v))
}

// TimeoutSecondsLT applies the LT predicate on the "timeout_seconds" field.
func TimeoutSecondsLT(v int) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldTimeoutSeconds, v))
}

// TimeoutSecondsLTE applies the LTE predicate on the "timeout_seconds" field.
func TimeoutSecondsLTE(v int) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldTimeoutSeconds, v))
}

// TimeoutSecondsIsNil applies the IsNil predicate on the "timeout_seconds" field.
func TimeoutSecondsIsNil() predicate.Job {
	return predicate.Job(sql.FieldIsNull(FieldTimeoutSeconds))
}

// TimeoutSecondsNotNil applies the NotNil predicate on the "timeout_seconds" field.
func TimeoutSecondsNotNil() predicate.Job {
	return predicate.Job(sql.FieldNotNull(FieldTimeoutSeconds))
}

// LastEditTimeEQ applies the EQ predicate on the "last_edit_time" field.
func LastEditTimeEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldLastEditTime, v))
}

// LastEditTimeNEQ applies the NEQ predicate on the "last_edit_time" field.
func LastEditTimeNEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldLastEditTime, v))
}

// LastEditTimeIn applies the In predicate on the "last_edit_time" field.
func LastEditTimeIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldLastEditTime, vs...))
}

// LastEditTimeNotIn applies the NotIn predicate on the "last_edit_time" field.
func LastEditTimeNotIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldLastEditTime, vs...))
}

// LastEditTimeGT applies the GT predicate on the "last_edit_time" field.
func LastEditTimeGT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldLastEditTime, v))
}

// LastEditTimeGTE applies the GTE predicate on the "last_edit_time" field.
func LastEditTimeGTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldLastEditTime, v))
}

// LastEditTimeLT applies the LT predicate on the "last_edit_time" field.
func LastEditTimeLT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldLastEditTime, v))
}

// LastEditTimeLTE applies the LTE predicate on the "last_edit_time" field.
func LastEditTimeLTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldLastEditTime, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldCreatedAt, v))
}

// NotifyOnStartChannelIdsIsNil applies the IsNil predicate on the "notify_on_start_channel_ids" field.
func NotifyOnStartChannelIdsIsNil() predicate.Job {
	return predicate.Job(sql.FieldIsNull(FieldNotifyOnStartChannelIds))
}

// NotifyOnStartChannelIdsNotNil applies the NotNil predicate on the "notify_on_start_channel_ids" field.
func NotifyOnStartChannelIdsNotNil() predicate.Job {
	return predicate.Job(sql.FieldNotNull(FieldNotifyOnStartChannelIds))
}

// NotifyOnSuccessChannelIdsIsNil applies the IsNil predicate on the "notify_on_success_channel_ids" field.
func NotifyOnSuccessChannelIdsIsNil() predicate.Job {
	return predicate.Job(sql.FieldIsNull(FieldNotifyOnSuccessChannelIds))
}

// NotifyOnSuccessChannelIdsNotNil applies the NotNil predicate on the "notify_on_success_channel_ids" field.
func NotifyOnSuccessChannelIdsNotNil() predicate.Job {
	return predicate.Job(sql.FieldNotNull(FieldNotifyOnSuccessChannelIds))
}

// NotifyOnFailureChannelIdsIsNil applies the IsNil predicate on the "notify_on_failure_channel_ids" field.
func NotifyOnFailureChannelIdsIsNil() predicate.Job {
	return predicate.Job(sql.FieldIsNull(FieldNotifyOnFailureChannelIds))
}

// NotifyOnFailureChannelIdsNotNil applies the NotNil predicate on the "notify_on_failure_channel_ids" field.
func NotifyOnFailureChannelIdsNotNil() predicate.Job {
	return predicate.Job(sql.FieldNotNull(FieldNotifyOnFailureChannelIds))
}

// LastRunTimeEQ applies the EQ predicate on the "last_run_time" field.
func LastRunTimeEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldLastRunTime, v))
}

// LastRunTimeNEQ applies the NEQ predicate on the "last_run_time" field.
func LastRunTimeNEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldLastRunTime, v))
}

// LastRunTimeIn applies the In predicate on the "last_run_time" field.
func LastRunTimeIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldLastRunTime, vs...))
}

// LastRunTimeNotIn applies the NotIn predicate on the "last_run_time" field.
func LastRunTimeNotIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldLastRunTime, vs...))
}

// LastRunTimeGT applies the GT predicate on the "last_run_time" field.
func LastRunTimeGT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldLastRunTime, v))
}

// LastRunTimeGTE applies the GTE predicate on the "last_run_time" field.
func LastRunTimeGTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldLastRunTime, v))
}

// LastRunTimeLT applies the LT predicate on the "last_run_time" field.
func LastRunTimeLT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldLastRunTime, v))
}

// LastRunTimeLTE applies the LTE predicate on the "last_run_time" field.
func LastRunTimeLTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldLastRunTime, v))
}

// NextCronRunTimeEQ applies the EQ predicate on the "next_cron_run_time" field.
func NextCronRunTimeEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldNextCronRunTime, v))
}

// NextCronRunTimeNEQ applies the NEQ predicate on the "next_cron_run_time" field.
func NextCronRunTimeNEQ(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldNextCronRunTime, v))
}

// NextCronRunTimeIn applies the In predicate on the "next_cron_run_time" field.
func NextCronRunTimeIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldNextCronRunTime, vs...))
}

// NextCronRunTimeNotIn applies the NotIn predicate on the "next_cron_run_time" field.
func NextCronRunTimeNotIn(vs ...time.Time) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldNextCronRunTime, vs...))
}

// NextCronRunTimeGT applies the GT predicate on the "next_cron_run_time" field.
func NextCronRunTimeGT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldNextCronRunTime, v))
}

// NextCronRunTimeGTE applies the GTE predicate on the "next_cron_run_time" field.
func NextCronRunTimeGTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldNextCronRunTime, v))
}

// NextCronRunTimeLT applies the LT predicate on the "next_cron_run_time" field.
func NextCronRunTimeLT(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldNextCronRunTime, v))
}

// NextCronRunTimeLTE applies the LTE predicate on the "next_cron_run_time" field.
func NextCronRunTimeLTE(v time.Time) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldNextCronRunTime, v))
}

// ScriptEQ applies the EQ predicate on the "script" field.
func ScriptEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldScript, v))
}

// ScriptNEQ applies the NEQ predicate on the "script" field.
func ScriptNEQ(v string) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldScript, v))
}

// ScriptIn applies the In predicate on the "script" field.
func ScriptIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldIn(FieldScript, vs...))
}

// ScriptNotIn applies the NotIn predicate on the "script" field.
func ScriptNotIn(vs ...string) predicate.Job {
	return predicate.Job(sql.FieldNotIn(FieldScript, vs...))
}

// ScriptGT applies the GT predicate on the "script" field.
func ScriptGT(v string) predicate.Job {
	return predicate.Job(sql.FieldGT(FieldScript, v))
}

// ScriptGTE applies the GTE predicate on the "script" field.
func ScriptGTE(v string) predicate.Job {
	return predicate.Job(sql.FieldGTE(FieldScript, v))
}

// ScriptLT applies the LT predicate on the "script" field.
func ScriptLT(v string) predicate.Job {
	return predicate.Job(sql.FieldLT(FieldScript, v))
}

// ScriptLTE applies the LTE predicate on the "script" field.
func ScriptLTE(v string) predicate.Job {
	return predicate.Job(sql.FieldLTE(FieldScript, v))
}

// ScriptContains applies the Contains predicate on the "script" field.
func ScriptContains(v string) predicate.Job {
	return predicate.Job(sql.FieldContains(FieldScript, v))
}

// ScriptHasPrefix applies the HasPrefix predicate on the "script" field.
func ScriptHasPrefix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasPrefix(FieldScript, v))
}

// ScriptHasSuffix applies the HasSuffix predicate on the "script" field.
func ScriptHasSuffix(v string) predicate.Job {
	return predicate.Job(sql.FieldHasSuffix(FieldScript, v))
}

// ScriptEqualFold applies the EqualFold predicate on the "script" field.
func ScriptEqualFold(v string) predicate.Job {
	return predicate.Job(sql.FieldEqualFold(FieldScript, v))
}

// ScriptContainsFold applies the ContainsFold predicate on the "script" field.
func ScriptContainsFold(v string) predicate.Job {
	return predicate.Job(sql.FieldContainsFold(FieldScript, v))
}

// LastRunSuccessEQ applies the EQ predicate on the "last_run_success" field.
func LastRunSuccessEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldLastRunSuccess, v))
}

// LastRunSuccessNEQ applies the NEQ predicate on the "last_run_success" field.
func LastRunSuccessNEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldLastRunSuccess, v))
}

// CreatedByAPIEQ applies the EQ predicate on the "created_by_api" field.
func CreatedByAPIEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldEQ(FieldCreatedByAPI, v))
}

// CreatedByAPINEQ applies the NEQ predicate on the "created_by_api" field.
func CreatedByAPINEQ(v bool) predicate.Job {
	return predicate.Job(sql.FieldNEQ(FieldCreatedByAPI, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHistory applies the HasEdge predicate on the "history" edge.
func HasHistory() predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HistoryTable, HistoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHistoryWith applies the HasEdge predicate on the "history" edge with a given conditions (other predicates).
func HasHistoryWith(preds ...predicate.JobHistory) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := newHistoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVersions applies the HasEdge predicate on the "versions" edge.
func HasVersions() predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VersionsTable, VersionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVersionsWith applies the HasEdge predicate on the "versions" edge with a given conditions (other predicates).
func HasVersionsWith(preds ...predicate.JobVersion) predicate.Job {
	return predicate.Job(func(s *sql.Selector) {
		step := newVersionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Job) predicate.Job {
	return predicate.Job(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Job) predicate.Job {
	return predicate.Job(sql.NotPredicates(p))
}
