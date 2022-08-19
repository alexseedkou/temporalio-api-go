// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/failed_cause.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Workflow tasks can fail for various reasons. Note that some of these reasons can only originate
// from the server, and some of them can only originate from the SDK/worker.
type WorkflowTaskFailedCause int32

const (
	WORKFLOW_TASK_FAILED_CAUSE_UNSPECIFIED WorkflowTaskFailedCause = 0
	// Between starting and completing the workflow task (with a workflow completion command), some
	// new command (like a signal) was processed into workflow history. The outstanding task will be
	// failed with this reason, and a worker must pick up a new task.
	WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND                                         WorkflowTaskFailedCause = 1
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES                          WorkflowTaskFailedCause = 2
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES                    WorkflowTaskFailedCause = 3
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES                                WorkflowTaskFailedCause = 4
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES                               WorkflowTaskFailedCause = 5
	WORKFLOW_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES                              WorkflowTaskFailedCause = 6
	WORKFLOW_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES                WorkflowTaskFailedCause = 7
	WORKFLOW_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES                    WorkflowTaskFailedCause = 8
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES                  WorkflowTaskFailedCause = 9
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES WorkflowTaskFailedCause = 10
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES                            WorkflowTaskFailedCause = 11
	WORKFLOW_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID                                  WorkflowTaskFailedCause = 12
	// The worker wishes to fail the task and have the next one be generated on a normal, not sticky
	// queue. Generally workers should prefer to use the explicit `ResetStickyTaskQueue` RPC call.
	WORKFLOW_TASK_FAILED_CAUSE_RESET_STICKY_TASK_QUEUE                  WorkflowTaskFailedCause = 13
	WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE        WorkflowTaskFailedCause = 14
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES WorkflowTaskFailedCause = 15
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES     WorkflowTaskFailedCause = 16
	WORKFLOW_TASK_FAILED_CAUSE_FORCE_CLOSE_COMMAND                      WorkflowTaskFailedCause = 17
	WORKFLOW_TASK_FAILED_CAUSE_FAILOVER_CLOSE_COMMAND                   WorkflowTaskFailedCause = 18
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE                    WorkflowTaskFailedCause = 19
	WORKFLOW_TASK_FAILED_CAUSE_RESET_WORKFLOW                           WorkflowTaskFailedCause = 20
	WORKFLOW_TASK_FAILED_CAUSE_BAD_BINARY                               WorkflowTaskFailedCause = 21
	WORKFLOW_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID           WorkflowTaskFailedCause = 22
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES                    WorkflowTaskFailedCause = 23
	// The worker encountered a mismatch while replaying history between what was expected, and
	// what the workflow code actually did.
	WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR WorkflowTaskFailedCause = 24
	WORKFLOW_TASK_FAILED_CAUSE_BAD_MEMO                WorkflowTaskFailedCause = 25
)

var WorkflowTaskFailedCause_name = map[int32]string{
	0:  "Unspecified",
	1:  "UnhandledCommand",
	2:  "BadScheduleActivityAttributes",
	3:  "BadRequestCancelActivityAttributes",
	4:  "BadStartTimerAttributes",
	5:  "BadCancelTimerAttributes",
	6:  "BadRecordMarkerAttributes",
	7:  "BadCompleteWorkflowExecutionAttributes",
	8:  "BadFailWorkflowExecutionAttributes",
	9:  "BadCancelWorkflowExecutionAttributes",
	10: "BadRequestCancelExternalWorkflowExecutionAttributes",
	11: "BadContinueAsNewAttributes",
	12: "StartTimerDuplicateId",
	13: "ResetStickyTaskQueue",
	14: "WorkflowWorkerUnhandledFailure",
	15: "BadSignalWorkflowExecutionAttributes",
	16: "BadStartChildExecutionAttributes",
	17: "ForceCloseCommand",
	18: "FailoverCloseCommand",
	19: "BadSignalInputSize",
	20: "ResetWorkflow",
	21: "BadBinary",
	22: "ScheduleActivityDuplicateId",
	23: "BadSearchAttributes",
	24: "NonDeterministicError",
	25: "BadMemo",
}

var WorkflowTaskFailedCause_value = map[string]int32{
	"Unspecified":                                         0,
	"UnhandledCommand":                                    1,
	"BadScheduleActivityAttributes":                       2,
	"BadRequestCancelActivityAttributes":                  3,
	"BadStartTimerAttributes":                             4,
	"BadCancelTimerAttributes":                            5,
	"BadRecordMarkerAttributes":                           6,
	"BadCompleteWorkflowExecutionAttributes":              7,
	"BadFailWorkflowExecutionAttributes":                  8,
	"BadCancelWorkflowExecutionAttributes":                9,
	"BadRequestCancelExternalWorkflowExecutionAttributes": 10,
	"BadContinueAsNewAttributes":                          11,
	"StartTimerDuplicateId":                               12,
	"ResetStickyTaskQueue":                                13,
	"WorkflowWorkerUnhandledFailure":                      14,
	"BadSignalWorkflowExecutionAttributes":                15,
	"BadStartChildExecutionAttributes":                    16,
	"ForceCloseCommand":                                   17,
	"FailoverCloseCommand":                                18,
	"BadSignalInputSize":                                  19,
	"ResetWorkflow":                                       20,
	"BadBinary":                                           21,
	"ScheduleActivityDuplicateId":                         22,
	"BadSearchAttributes":                                 23,
	"NonDeterministicError":                               24,
	"BadMemo":                                             25,
}

func (WorkflowTaskFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{0}
}

type StartChildWorkflowExecutionFailedCause int32

const (
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED             StartChildWorkflowExecutionFailedCause = 0
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS StartChildWorkflowExecutionFailedCause = 1
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND     StartChildWorkflowExecutionFailedCause = 2
)

var StartChildWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "WorkflowAlreadyExists",
	2: "NamespaceNotFound",
}

var StartChildWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":           0,
	"WorkflowAlreadyExists": 1,
	"NamespaceNotFound":     2,
}

func (StartChildWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{1}
}

type CancelExternalWorkflowExecutionFailedCause int32

const (
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           CancelExternalWorkflowExecutionFailedCause = 0
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND CancelExternalWorkflowExecutionFailedCause = 1
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND                   CancelExternalWorkflowExecutionFailedCause = 2
)

var CancelExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
	2: "NamespaceNotFound",
}

var CancelExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
	"NamespaceNotFound":                 2,
}

func (CancelExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{2}
}

type SignalExternalWorkflowExecutionFailedCause int32

const (
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           SignalExternalWorkflowExecutionFailedCause = 0
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND SignalExternalWorkflowExecutionFailedCause = 1
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND                   SignalExternalWorkflowExecutionFailedCause = 2
)

var SignalExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
	2: "NamespaceNotFound",
}

var SignalExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
	"NamespaceNotFound":                 2,
}

func (SignalExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{3}
}

type ResourceExhaustedCause int32

const (
	RESOURCE_EXHAUSTED_CAUSE_UNSPECIFIED ResourceExhaustedCause = 0
	// Caller exceeds request per second limit.
	RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT ResourceExhaustedCause = 1
	// Caller exceeds max concurrent request limit.
	RESOURCE_EXHAUSTED_CAUSE_CONCURRENT_LIMIT ResourceExhaustedCause = 2
	// System overloaded.
	RESOURCE_EXHAUSTED_CAUSE_SYSTEM_OVERLOADED ResourceExhaustedCause = 3
	// Namespace exceeds persistence rate limit.
	RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_LIMIT ResourceExhaustedCause = 4
)

var ResourceExhaustedCause_name = map[int32]string{
	0: "Unspecified",
	1: "RpsLimit",
	2: "ConcurrentLimit",
	3: "SystemOverloaded",
	4: "PersistenceLimit",
}

var ResourceExhaustedCause_value = map[string]int32{
	"Unspecified":      0,
	"RpsLimit":         1,
	"ConcurrentLimit":  2,
	"SystemOverloaded": 3,
	"PersistenceLimit": 4,
}

func (ResourceExhaustedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{4}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowTaskFailedCause", WorkflowTaskFailedCause_name, WorkflowTaskFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.StartChildWorkflowExecutionFailedCause", StartChildWorkflowExecutionFailedCause_name, StartChildWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.CancelExternalWorkflowExecutionFailedCause", CancelExternalWorkflowExecutionFailedCause_name, CancelExternalWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.SignalExternalWorkflowExecutionFailedCause", SignalExternalWorkflowExecutionFailedCause_name, SignalExternalWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.ResourceExhaustedCause", ResourceExhaustedCause_name, ResourceExhaustedCause_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/failed_cause.proto", fileDescriptor_b293cf8d1d965f2d)
}

var fileDescriptor_b293cf8d1d965f2d = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x73, 0xdb, 0x44,
	0x1c, 0xb5, 0xdc, 0x52, 0x60, 0x29, 0x20, 0x16, 0xd2, 0x86, 0x8b, 0x66, 0xf8, 0x17, 0x5a, 0x0f,
	0x38, 0x75, 0x4b, 0x9b, 0xa9, 0x0d, 0x13, 0xd6, 0xab, 0x9f, 0xe3, 0x9d, 0x48, 0x2b, 0x77, 0x77,
	0x95, 0xd8, 0xbd, 0xec, 0x88, 0x54, 0x6d, 0x35, 0x75, 0x2d, 0x8f, 0xff, 0x94, 0x1c, 0xf9, 0x08,
	0xf0, 0x2d, 0x18, 0x3e, 0x01, 0x7c, 0x03, 0x8e, 0x39, 0xf6, 0x48, 0x9c, 0x0b, 0xc3, 0xa9, 0x33,
	0x9c, 0xb8, 0x31, 0x32, 0x71, 0x50, 0xa8, 0x2d, 0xd9, 0xdc, 0x24, 0xed, 0x7b, 0x6f, 0xf5, 0x7b,
	0xfb, 0xf6, 0xb7, 0x8b, 0xae, 0x8d, 0xc2, 0xa7, 0xfd, 0x78, 0x10, 0x74, 0x37, 0x83, 0x7e, 0xb4,
	0x19, 0xf6, 0xc6, 0x4f, 0x87, 0x9b, 0xcf, 0x2a, 0x9b, 0x0f, 0x83, 0xa8, 0x1b, 0x3e, 0xd0, 0x07,
	0xc1, 0x78, 0x18, 0x96, 0xfb, 0x83, 0x78, 0x14, 0xe3, 0xb5, 0x19, 0xb2, 0x1c, 0xf4, 0xa3, 0xf2,
	0x14, 0x59, 0x7e, 0x56, 0x29, 0xfd, 0x72, 0x19, 0x5d, 0xdd, 0x8f, 0x07, 0x4f, 0x1e, 0x76, 0xe3,
	0x6f, 0x55, 0x30, 0x7c, 0xd2, 0x98, 0x32, 0x69, 0x42, 0xc4, 0x25, 0xb4, 0xb1, 0xef, 0x89, 0xdd,
	0x86, 0xe3, 0xed, 0x6b, 0x45, 0xe4, 0xae, 0x6e, 0x10, 0xe6, 0x80, 0xad, 0x29, 0xf1, 0x25, 0x68,
	0x9f, 0xcb, 0x16, 0x50, 0xd6, 0x60, 0x60, 0x9b, 0x05, 0x7c, 0x03, 0x7d, 0x96, 0x89, 0x6d, 0x12,
	0x6e, 0x4f, 0xdf, 0x3d, 0xd7, 0x25, 0xdc, 0x36, 0x0d, 0xbc, 0x8d, 0x6a, 0x19, 0x8c, 0x3a, 0xb1,
	0xb5, 0xa4, 0x4d, 0xb0, 0x7d, 0x07, 0x34, 0xa1, 0x8a, 0xed, 0x31, 0xd5, 0xd1, 0x44, 0x29, 0xc1,
	0xea, 0xbe, 0x02, 0x69, 0x16, 0x31, 0x20, 0x92, 0x23, 0x20, 0xe0, 0x9e, 0x0f, 0x52, 0x69, 0x4a,
	0x38, 0x05, 0x67, 0xae, 0xcc, 0x05, 0x7c, 0x17, 0xdd, 0xce, 0xfb, 0x0f, 0x45, 0x84, 0xd2, 0x8a,
	0xb9, 0x20, 0xd2, 0xd4, 0x8b, 0xb8, 0x8a, 0xee, 0xe4, 0x50, 0x4f, 0x67, 0x7e, 0x89, 0xfb, 0x0a,
	0xae, 0xa1, 0xad, 0xdc, 0xbf, 0xa7, 0x9e, 0xb0, 0xb5, 0x4b, 0xc4, 0xee, 0x79, 0xf2, 0x25, 0xcc,
	0x10, 0xe4, 0x4d, 0xec, 0xb9, 0x2d, 0x07, 0x14, 0xe8, 0x33, 0x1c, 0xb4, 0x81, 0xfa, 0x8a, 0x79,
	0x3c, 0x2d, 0xf5, 0xea, 0x12, 0x2e, 0x26, 0x1f, 0x72, 0x64, 0x5e, 0xc3, 0x3b, 0x88, 0x2e, 0x67,
	0x45, 0xb6, 0xd0, 0xeb, 0xb8, 0x8d, 0xd4, 0x6a, 0xab, 0x0a, 0x6d, 0x05, 0x82, 0x93, 0x3c, 0x65,
	0x84, 0xbf, 0x42, 0x77, 0x73, 0x4d, 0xe3, 0x8a, 0x71, 0x1f, 0x34, 0x91, 0x9a, 0xc3, 0x7e, 0x9a,
	0xfe, 0x06, 0xde, 0x42, 0xb7, 0x32, 0xe8, 0xe9, 0x8c, 0xd8, 0x7e, 0xcb, 0x61, 0x94, 0x28, 0xd0,
	0xcc, 0x36, 0x2f, 0xe3, 0x3b, 0xe8, 0x66, 0x06, 0x51, 0x80, 0x04, 0xa5, 0xa5, 0x62, 0x74, 0xb7,
	0xf3, 0xcf, 0xf0, 0x3d, 0x1f, 0x7c, 0x30, 0xdf, 0xc4, 0x5f, 0xa3, 0x2f, 0x33, 0x78, 0x67, 0x43,
	0xc9, 0x03, 0x88, 0xd4, 0x16, 0x4b, 0x60, 0xbe, 0x00, 0xf3, 0xad, 0x25, 0x16, 0x45, 0xb2, 0x9d,
	0x7c, 0xeb, 0xde, 0xc6, 0x14, 0x6d, 0x2f, 0xb5, 0x47, 0x68, 0x93, 0x39, 0xf6, 0x7c, 0x11, 0x13,
	0xdf, 0x44, 0xe5, 0x0c, 0x91, 0x86, 0x27, 0x28, 0x68, 0xea, 0x78, 0x12, 0xce, 0x9a, 0xc4, 0x3b,
	0xf8, 0x36, 0xaa, 0x64, 0x71, 0x08, 0x73, 0xbc, 0x3d, 0x10, 0xff, 0xa1, 0x61, 0xfc, 0x05, 0xba,
	0xb1, 0x5c, 0xe1, 0x8c, 0xb7, 0x7c, 0xa5, 0x25, 0xbb, 0x0f, 0xe6, 0xbb, 0xf8, 0x73, 0x74, 0x3d,
	0x77, 0xa1, 0x66, 0x00, 0xf3, 0x3d, 0x7c, 0x1d, 0x7d, 0x92, 0x33, 0x49, 0x9d, 0x71, 0x22, 0x3a,
	0xe6, 0x5a, 0x4e, 0xf4, 0x5e, 0xee, 0x73, 0xe7, 0x12, 0x74, 0x65, 0x99, 0x72, 0x80, 0x08, 0xda,
	0x4c, 0xfb, 0x7d, 0x35, 0x27, 0x77, 0xdc, 0xe3, 0xda, 0x06, 0x05, 0xc2, 0x65, 0x9c, 0x25, 0xf1,
	0xd3, 0x20, 0x84, 0x27, 0xcc, 0x75, 0xfc, 0x29, 0xfa, 0x28, 0x67, 0x36, 0x17, 0x5c, 0xcf, 0x7c,
	0xbf, 0xf4, 0xa7, 0x81, 0x36, 0xe4, 0x28, 0x18, 0x8c, 0xe8, 0xe3, 0xa8, 0xfb, 0x60, 0x76, 0x8a,
	0xc0, 0x61, 0x78, 0x30, 0x1e, 0x45, 0x71, 0x2f, 0x7d, 0x94, 0xd4, 0xd0, 0x56, 0x3a, 0x21, 0x73,
	0xf2, 0x96, 0x71, 0xb6, 0xec, 0x20, 0xba, 0x0a, 0xf9, 0x6c, 0x9c, 0x38, 0x02, 0x88, 0xdd, 0xd1,
	0xd0, 0x66, 0x52, 0x49, 0xd3, 0x48, 0x62, 0xbc, 0x8a, 0x10, 0x27, 0x2e, 0xc8, 0x16, 0xa1, 0x89,
	0x59, 0x4a, 0x37, 0x3c, 0x9f, 0xdb, 0x66, 0xb1, 0xf4, 0x43, 0x11, 0x95, 0x68, 0xd0, 0x3b, 0x08,
	0xbb, 0x70, 0x38, 0x0a, 0x07, 0xbd, 0xa0, 0x9b, 0x59, 0xf9, 0x36, 0xaa, 0x2d, 0xd1, 0xa8, 0x32,
	0xaa, 0xef, 0x20, 0x7f, 0x55, 0x81, 0x2c, 0xe0, 0xbf, 0xa5, 0x18, 0x89, 0xb1, 0xab, 0x4a, 0x2f,
	0xf6, 0x44, 0x46, 0x8f, 0x7a, 0xc1, 0xd2, 0x9e, 0x9c, 0xee, 0xbf, 0xff, 0xef, 0xc9, 0xaa, 0x02,
	0x2b, 0x78, 0xb2, 0xaa, 0xf4, 0x7c, 0x4f, 0xfe, 0x32, 0xd0, 0x15, 0x11, 0x0e, 0xe3, 0xf1, 0xe0,
	0x20, 0x84, 0xc3, 0xc7, 0xc1, 0x78, 0x38, 0x9a, 0xd5, 0x7f, 0x0d, 0x7d, 0x2c, 0x40, 0x7a, 0x7e,
	0xd2, 0xf1, 0xa0, 0xdd, 0x24, 0xbe, 0x54, 0x0b, 0x0a, 0xdd, 0x40, 0x1f, 0x2e, 0x44, 0x8a, 0x96,
	0xd4, 0x0e, 0x73, 0x99, 0x32, 0x8d, 0xa4, 0x75, 0x2d, 0xc4, 0x51, 0x8f, 0x53, 0x5f, 0x08, 0xe0,
	0xea, 0x14, 0x5e, 0xc4, 0x65, 0x54, 0x5a, 0x08, 0x97, 0x1d, 0xa9, 0xc0, 0xd5, 0x49, 0x5f, 0x75,
	0x3c, 0x62, 0x83, 0x6d, 0x5e, 0xc8, 0xc4, 0xb7, 0x40, 0x48, 0x26, 0x15, 0x70, 0x0a, 0xa7, 0xfa,
	0x17, 0xeb, 0x3f, 0x1b, 0x47, 0xc7, 0x56, 0xe1, 0xf9, 0xb1, 0x55, 0x78, 0x71, 0x6c, 0x19, 0xdf,
	0x4d, 0x2c, 0xe3, 0xc7, 0x89, 0x65, 0xfc, 0x3a, 0xb1, 0x8c, 0xa3, 0x89, 0x65, 0xfc, 0x36, 0xb1,
	0x8c, 0xdf, 0x27, 0x56, 0xe1, 0xc5, 0xc4, 0x32, 0xbe, 0x3f, 0xb1, 0x0a, 0x47, 0x27, 0x56, 0xe1,
	0xf9, 0x89, 0x55, 0x40, 0xeb, 0x51, 0x5c, 0x9e, 0x7b, 0x4d, 0xad, 0x9b, 0xa9, 0xf8, 0xb4, 0x92,
	0xfb, 0x6c, 0xcb, 0xb8, 0xff, 0xc1, 0xa3, 0x14, 0x3a, 0x8a, 0xcf, 0xdd, 0x80, 0x6b, 0xd3, 0x87,
	0x9f, 0x8a, 0x6b, 0x6a, 0x06, 0x20, 0xfd, 0xa8, 0x0c, 0x53, 0xb9, 0xbd, 0xca, 0x1f, 0xc5, 0xf5,
	0xd9, 0xf7, 0x6a, 0x95, 0xf4, 0xa3, 0x6a, 0x75, 0x3a, 0x52, 0xad, 0xee, 0x55, 0xbe, 0xb9, 0x34,
	0xbd, 0x2e, 0xdf, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x55, 0x28, 0xbc, 0x9d, 0x5a, 0x0b, 0x00,
	0x00,
}

func (x WorkflowTaskFailedCause) String() string {
	s, ok := WorkflowTaskFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x StartChildWorkflowExecutionFailedCause) String() string {
	s, ok := StartChildWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x CancelExternalWorkflowExecutionFailedCause) String() string {
	s, ok := CancelExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x SignalExternalWorkflowExecutionFailedCause) String() string {
	s, ok := SignalExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ResourceExhaustedCause) String() string {
	s, ok := ResourceExhaustedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
