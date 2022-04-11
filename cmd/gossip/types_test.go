/*
 * MIT License
 *
 * Copyright (c) 2022 Ludovic Cleroux
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package gossip

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name   string
		from   ClusterMetadata
		to     ClusterMetadata
		expect ClusterMetadata
	}{
		{
			name: "different",
			to: ClusterMetadata{
				"b": NodeState{
					"b": &VersionedStr{Version: 0, Value: "b"},
				},
			},
			from: ClusterMetadata{
				"a": NodeState{
					"a": &VersionedStr{Version: 0, Value: "a"},
				},
			},
			expect: ClusterMetadata{
				"a": NodeState{
					"a": &VersionedStr{Version: 0, Value: "a"},
				},
				"b": NodeState{
					"b": &VersionedStr{Version: 0, Value: "b"},
				},
			},
		}, {
			name: "different version",
			to: ClusterMetadata{
				"a": NodeState{
					"a": &VersionedStr{Version: 0, Value: "a"},
					"b": &VersionedStr{Version: 0, Value: "b"},
				},
			},
			from: ClusterMetadata{
				"a": NodeState{
					"a": &VersionedStr{Version: 1, Value: "a"},
				},
			},
			expect: ClusterMetadata{
				"a": NodeState{
					"a": &VersionedStr{Version: 1, Value: "a"},
					"b": &VersionedStr{Version: 0, Value: "b"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			jsonBytes, _ := json.Marshal(tt.to)
			t.Log(string(jsonBytes))

			merge(tt.to, tt.from)

			jsonBytes, _ = json.Marshal(tt.to)
			t.Log(string(jsonBytes))

			assert.Equal(t, tt.expect, tt.to)
		})
	}
}
