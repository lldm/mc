/*
 * Minio Client (C) 2014, 2015 Minio, Inc.
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

package main

import (
	"encoding/json"
	"os"

	"github.com/minio/mc/internal/github.com/minio/minio/pkg/probe"
	"github.com/minio/mc/pkg/console"
)

// fatalIf wrapper function which takes error and selectively prints stack frames if available on debug
func fatalIf(err *probe.Error) {
	if err == nil {
		return
	}
	if globalJSONFlag {
		errorMessage := struct {
			Cause     error              `json:"cause"`
			Type      string             `json:"type"`
			CallTrace []probe.TracePoint `json:"trace,omitempty"`
			SysInfo   map[string]string  `json:"sysinfo"`
		}{
			Type:    "Fatal",
			Cause:   err.ToGoError(),
			SysInfo: err.SysInfo,
		}
		if globalDebugFlag {
			errorMessage.CallTrace = err.CallTrace
		}
		json, err := json.Marshal(errorMessage)
		if err != nil {
			panic(err)
		}
		console.Println(string(json))
		os.Exit(1)
	}
	if !globalDebugFlag {
		console.Fatalln(err.ToGoError())
	}
	console.Fatalln(err)
}

// errorIf synonymous with fatalIf but doesn't exit on error != nil
func errorIf(err *probe.Error) {
	if err == nil {
		return
	}
	if globalJSONFlag {
		errorMessage := struct {
			Cause     error              `json:"cause"`
			Type      string             `json:"type"`
			CallTrace []probe.TracePoint `json:"trace,omitempty"`
			SysInfo   map[string]string  `json:"sysinfo"`
		}{
			Type:    "Error",
			Cause:   err.ToGoError(),
			SysInfo: err.SysInfo,
		}
		if globalDebugFlag {
			errorMessage.CallTrace = err.CallTrace
		}
		json, err := json.Marshal(errorMessage)
		if err != nil {
			panic(err)
		}
		console.Println(string(json))
		return
	}
	if !globalDebugFlag {
		console.Errorln(err.ToGoError())
		return
	}
	console.Errorln(err)
}