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

// Represents a mathematical expression tree.
package tree

// An operator node comprising a mathematical operation and the nodes it operates on.
type Operator struct {
	Operator string

	Left  Node
	Right Node
}

// Computes the mathematical operation.
func (operator Operator) Calculate() float64 {
	left := operator.Left.Calculate()
	right := operator.Right.Calculate()

	switch operator.Operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}

	panic("Unknown operator: " + operator.Operator)
}

// Returns the left and right nodes.
func (operator Operator) Children() []Node {
	return []Node{
		operator.Left,
		operator.Right,
	}
}
