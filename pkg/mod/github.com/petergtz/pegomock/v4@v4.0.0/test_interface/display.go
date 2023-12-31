// Copyright 2015 Peter Goetz
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

package test_interface

import (
	"io"
	"net/http"
	"time"
)

// Display is some sample interface to be mocked.
type Display interface {
	Flash(_param0 string, _param1 int)
	Show(_param0 string)
	SomeValue() string
	MultipleValues() (string, int, float32)
	MultipleParamsAndReturnValue(s string, i int) string
	ArrayParam(array []string)
	MapParam(m map[string]http.Request)
	FloatParam(float32)
	InterfaceParam(interface{})
	InterfaceReturnValue() interface{}
	ErrorReturnValue() error
	ErrorParam(e error)
	NetHttpRequestParam(r http.Request)
	NetHttpRequestPtrParam(r *http.Request)
	FuncReturnValue() func()
	VariadicParam(v ...string)
	NormalAndVariadicParam(s string, i int, v ...string)
	CamelCaseTypeParam(camelCaseParam io.ReadCloser)
	MapOfStringToInterfaceParam(m map[string]interface{})
	UseTime(t time.Time)
	ChanParams(<-chan string, chan<- error)
	ChanReturnValues() (<-chan string, chan<- error)
	VariadicWithNonPrimitiveType(m ...map[int]int)
	MapWithRedundantImports(m map[http.File]http.File)
	MapOfStringToEmptyUnnamedStruct(m map[string]struct{})
}

type GenericDisplay[N comparable, V Number] interface {
	GenericParams(m map[N]V) V
}

type Number interface {
	int64 | float64
}
