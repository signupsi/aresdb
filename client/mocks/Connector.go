//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v1.0.0
package mocks

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/uber/aresdb/client"
	"github.com/uber/aresdb/memstore/common"
)

// Connector is an autogenerated mock type for the Connector type
type Connector struct {
	mock.Mock
}

// Insert provides a mock function with given fields: tableName, columnNames, rows
func (_m *Connector) Insert(tableName string, columnNames []string, rows []client.Row, updateModes ...common.ColumnUpdateMode) (int, error) {
	ret := _m.Called(tableName, columnNames, rows)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, []string, []client.Row, ...common.ColumnUpdateMode) int); ok {
		r0 = rf(tableName, columnNames, rows, updateModes...)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, []client.Row, ...common.ColumnUpdateMode) error); ok {
		r1 = rf(tableName, columnNames, rows, updateModes...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close the connection
func (_m *Connector) Close() {
	_m.Close()
}
