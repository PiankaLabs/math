/*
 *   Copyright 2020 Pianka Labs LLC (https://www.piankalabs.com)
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

// Package tree represents a mathematical expression tree.
package tree

// A Number node in a mathematical expression tree comprising a single number.
type Number struct {
	Value float64
}

// Calculate returns just the number's value.
func (number Number) Calculate() float64 {
	return number.Value
}

// Children returns nil.
func (number Number) Children() []Node {
	return nil
}
