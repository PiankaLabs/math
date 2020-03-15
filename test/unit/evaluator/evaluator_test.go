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
	asserter "github.com/stretchr/testify/assert"

	// Internal
	"piankalabs.com/math/pkg/evaluator"
)

func TestParser(t *testing.T) {
	eval := new(evaluator.Evaluator)
	tree, err := eval.Evaluate("2 + 2")

	if err != nil {
		t.Log(err)
	}

	result := tree.Calculate()

	assert := asserter.New(t)
	assert.Equal(result, 4.0, "2 + 2 should equal 4")
}