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

import (
	// System
	"testing"

	// Third Party
	"github.com/stretchr/testify/assert"
)

func node() Number {
	return Number{
		Value: 4,
	}
}

func TestNumberCalculate(t *testing.T) {
	result := node().Calculate()

	assert.Equal(t, 4.0, result, "Number.Calculate() should return its numerical value")
}

func TestNumberChildren(t *testing.T) {
	result := node().Children()

	assert.Len(t, result, 0, "Number.Children() should not have any children")
}
