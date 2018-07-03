/*
 * Copyright 2018 The ThunderDB Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by mockery v1.0.0. DO NOT EDIT.
package kayak

import (
	"context"

	"github.com/stretchr/testify/mock"
	"gitlab.com/thunderdb/ThunderDB/twopc"
)

// MockWorker is an autogenerated mock type for the Worker type
type MockWorker struct {
	mock.Mock
}

// Commit provides a mock function with given fields: ctx, wb
func (_m *MockWorker) Commit(ctx context.Context, wb twopc.WriteBatch) error {
	ret := _m.Called(ctx, wb)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, twopc.WriteBatch) error); ok {
		r0 = rf(ctx, wb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Prepare provides a mock function with given fields: ctx, wb
func (_m *MockWorker) Prepare(ctx context.Context, wb twopc.WriteBatch) error {
	ret := _m.Called(ctx, wb)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, twopc.WriteBatch) error); ok {
		r0 = rf(ctx, wb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields: ctx, wb
func (_m *MockWorker) Rollback(ctx context.Context, wb twopc.WriteBatch) error {
	ret := _m.Called(ctx, wb)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, twopc.WriteBatch) error); ok {
		r0 = rf(ctx, wb)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
