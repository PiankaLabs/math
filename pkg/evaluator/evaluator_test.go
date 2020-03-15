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

package evaluator

import (
	// System
	"testing"

	// Third Party
	"github.com/stretchr/testify/assert"
)

func test(t *testing.T, expression string, expected float64, message string) {
	eval := new(Evaluator)
	tree, err := eval.Evaluate(expression)

	if err != nil {
		t.Log(err)
	}

	result := tree.Calculate()

	assert.Equal(t, expected, result, message)
}

func TestNumber(t *testing.T) {
	test(t, "4", 4, "4 should equal 4")
}

func TestUnary(t *testing.T) {
	test(t, "-4", -4, "-4 should equal -4")
}

func TestInfixAdd(t *testing.T) {
	test(t, "4 + 4", 8, "4 + 4 should equal 8")
}

func TestInfixSubtract(t *testing.T) {
	test(t, "4 - 4", 0, "4 - 4 should equal 0")
}

func TestInfixMultiply(t *testing.T) {
	test(t, "4 * 4", 16, "4 * 4 should equal 16")
}

func TestInfixDivide(t *testing.T) {
	test(t, "4 / 4", 1, "4 / 4 should equal 1")
}

func TestParenthetical(t *testing.T) {
	test(t, "4 / (4 * 4)", 0.25, "4 / (4 * 4) should equal 0.25")
}