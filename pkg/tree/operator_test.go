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

func operator(left float64, operator string, right float64) Operator {
	return Operator{
		Operator: operator,
		Left:     number(left),
		Right:    number(right),
	}
}

func number(value float64) Number {
	return Number{
		Value: value,
	}
}

func TestOperatorCalculate(t *testing.T) {
	add := operator(4, "+", 4).Calculate()
	assert.Equal(t, 8.0, add, "Operator.Calculate() should correctly add values")

	subtract := operator(4, "-", 4).Calculate()
	assert.Equal(t, 0.0, subtract, "Operator.Calculate() should correctly subtract values")

	multiply := operator(4, "*", 4).Calculate()
	assert.Equal(t, 16.0, multiply, "Operator.Calculate() should correctly multiply values")

	divide := operator(4, "/", 4).Calculate()
	assert.Equal(t, 1.0, divide, "Operator.Calculate() should correctly divide values")
}

func TestOperatorChildren(t *testing.T) {
	result := operator(4, "+", 4).Children()

	assert.Len(t, result, 2, "Number.Children() should always have two children")
}
