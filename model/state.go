// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package model

// DeviceStates maps integration IDs to device states
type DeviceStates map[string]DeviceState

// DeviceState is the device state for/from a specific integration
type DeviceState struct {
	// Desired state is only mutable for the cloud.
	Desired map[string]interface{} `json:"desired"`
	// Reported state is only mutable for the device.
	Reported map[string]interface{} `json:"reported"`
}
