// Copyright 2014 Google Inc. All Rights Reserved.
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/gwos/bokzer/utils/fs (interfaces: FileSystem)

package mockfs

import (
	gomock "code.google.com/p/gomock/gomock"
	fs "github.com/gwos/bokzer/utils/fs"
)

// Mock of FileSystem interface
type MockFileSystem struct {
	ctrl     *gomock.Controller
	recorder *_MockFileSystemRecorder
}

// Recorder for MockFileSystem (not exported)
type _MockFileSystemRecorder struct {
	mock *MockFileSystem
}

func NewMockFileSystem(ctrl *gomock.Controller) *MockFileSystem {
	mock := &MockFileSystem{ctrl: ctrl}
	mock.recorder = &_MockFileSystemRecorder{mock}
	return mock
}

func (_m *MockFileSystem) EXPECT() *_MockFileSystemRecorder {
	return _m.recorder
}

func (_m *MockFileSystem) Open(_param0 string) (fs.File, error) {
	ret := _m.ctrl.Call(_m, "Open", _param0)
	ret0, _ := ret[0].(fs.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemRecorder) Open(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open", arg0)
}
