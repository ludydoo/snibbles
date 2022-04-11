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

type NodeState map[string]*VersionedStr
type ClusterMetadata map[string]NodeState

// delta returns the delta between two ClusterMetadata
func delta(from ClusterMetadata, to ClusterMetadata) ClusterMetadata {
	diff := make(ClusterMetadata)

	for fNodeId, fState := range from {
		if _, ok := to[fNodeId]; !ok {
			diff[fNodeId] = fState
			continue
		}
		fromState := from[fNodeId]
		toState := to[fNodeId]
		diffS := diffState(fromState, toState)
		if len(diffS) > 0 {
			diff[fNodeId] = diffS
		}
	}

	return diff
}

// diffState returns the delta between two NodeState
func diffState(from, to NodeState) NodeState {
	diff := make(NodeState)
	for fKey, fValue := range from {
		if _, ok := to[fKey]; !ok {
			diff[fKey] = fValue
			continue
		}
		toValue := to[fKey]
		if fValue.Version != toValue.Version || fValue.Value != toValue.Value {
			if fValue.Version > toValue.Version {
				diff[fKey] = fValue
			} else {
				diff[fKey] = toValue
			}
		}
	}
	return diff
}

// merge merges two ClusterMetadata
func merge(to ClusterMetadata, from ClusterMetadata) {
	diff := delta(from, to)
	for nodeId := range diff {
		if _, ok := to[nodeId]; !ok {
			to[nodeId] = diff[nodeId]
		} else {
			if _, ok := to[nodeId]; !ok {
				to[nodeId] = make(NodeState)
			}
			for k, v := range diff[nodeId] {
				to[nodeId][k] = v
			}
		}
	}
}
