// Copyright 2018 Authors of Cilium
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

package option

var (
	endpointMutableOptionLibrary = OptionLibrary{
		ConntrackAccounting: &specConntrackAccounting,
		ConntrackLocal:      &specConntrackLocal,
		Conntrack:           &specConntrack,
		Debug:               &specDebug,
		DebugLB:             &specDebugLB,
		DropNotify:          &specDropNotify,
		TraceNotify:         &specTraceNotify,
		MonitorAggregation:  &specMonitorAggregation,
		NAT46:               &specNAT46,
	}
)

func GetEndpointMutableOptionLibrary() OptionLibrary {
	opt := OptionLibrary{}
	for k, v := range endpointMutableOptionLibrary {
		opt[k] = v
	}
	return opt
}
