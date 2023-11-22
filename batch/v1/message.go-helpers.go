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

package batch

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type BatchOperationInfo to the protobuf v3 wire format
func (val *BatchOperationInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BatchOperationInfo from the protobuf v3 wire format
func (val *BatchOperationInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BatchOperationInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BatchOperationInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BatchOperationInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BatchOperationInfo
	switch t := that.(type) {
	case *BatchOperationInfo:
		that1 = t
	case BatchOperationInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BatchOperationTermination to the protobuf v3 wire format
func (val *BatchOperationTermination) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BatchOperationTermination from the protobuf v3 wire format
func (val *BatchOperationTermination) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BatchOperationTermination) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BatchOperationTermination values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BatchOperationTermination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BatchOperationTermination
	switch t := that.(type) {
	case *BatchOperationTermination:
		that1 = t
	case BatchOperationTermination:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BatchOperationSignal to the protobuf v3 wire format
func (val *BatchOperationSignal) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BatchOperationSignal from the protobuf v3 wire format
func (val *BatchOperationSignal) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BatchOperationSignal) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BatchOperationSignal values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BatchOperationSignal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BatchOperationSignal
	switch t := that.(type) {
	case *BatchOperationSignal:
		that1 = t
	case BatchOperationSignal:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BatchOperationCancellation to the protobuf v3 wire format
func (val *BatchOperationCancellation) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BatchOperationCancellation from the protobuf v3 wire format
func (val *BatchOperationCancellation) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BatchOperationCancellation) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BatchOperationCancellation values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BatchOperationCancellation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BatchOperationCancellation
	switch t := that.(type) {
	case *BatchOperationCancellation:
		that1 = t
	case BatchOperationCancellation:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BatchOperationDeletion to the protobuf v3 wire format
func (val *BatchOperationDeletion) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BatchOperationDeletion from the protobuf v3 wire format
func (val *BatchOperationDeletion) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BatchOperationDeletion) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BatchOperationDeletion values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BatchOperationDeletion) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BatchOperationDeletion
	switch t := that.(type) {
	case *BatchOperationDeletion:
		that1 = t
	case BatchOperationDeletion:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BatchOperationReset to the protobuf v3 wire format
func (val *BatchOperationReset) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BatchOperationReset from the protobuf v3 wire format
func (val *BatchOperationReset) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BatchOperationReset) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BatchOperationReset values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BatchOperationReset) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BatchOperationReset
	switch t := that.(type) {
	case *BatchOperationReset:
		that1 = t
	case BatchOperationReset:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}