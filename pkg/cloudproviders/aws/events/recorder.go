/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package events

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
)

type Event struct {
	InvolvedObject runtime.Object
	Type           string
	Reason         string
	Message        string
}

type Recorder struct {
	rec record.EventRecorder
}

func NewRecorder(r record.EventRecorder) *Recorder {
	return &Recorder{
		rec: r,
	}
}

// Publish creates a Kubernetes event using the passed event struct
func (r *Recorder) Publish(evt Event) {
	r.rec.Event(evt.InvolvedObject, evt.Type, evt.Reason, evt.Message)
}
