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

// based on https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/

// Turns a mathematical expression into a tree.
package evaluator

import (
	"log"
	"strconv"

	"github.com/golang-collections/collections/stack"

	"piankalabs.com/math/pkg/parser"
	"piankalabs.com/math/pkg/tree"
)

// For handling tokens when walking a parse tree.
type Listener struct {
	*parser.BaseMathListener

	Tree tree.Node

	stack *stack.Stack
}

// Initializes the listener.
func (listener *Listener) EnterExpression(c *parser.ExpressionContext) {
	listener.stack = new(stack.Stack)
}

// Expression parse tree has been fully walked.
func (listener *Listener) ExitExpression(c *parser.ExpressionContext) {
	listener.Tree = listener.stack.Pop().(tree.Node)
}

// A unary statement has been walked.
func (listener *Listener) ExitUnaryStatement(c *parser.UnaryStatementContext) {
	operator := c.Unary().GetOperator().GetText()
	if operator == "-" {
		var negativeOne tree.Node = tree.Number{
			Value: -1,
		}

		left := negativeOne
		right := listener.stack.Pop()

		var node tree.Node = tree.Operator{
			Operator: "*",
			Left:     left,
			Right:    right.(tree.Node),
		}

		listener.stack.Push(node)
	}
}

// A number statement has been walked.
func (listener *Listener) ExitNumberStatement(c *parser.NumberStatementContext) {
	value, err := strconv.ParseFloat(c.Number().GetText(), 64)
	if err != nil {
		log.Fatal(err)
	}

	var node tree.Node = tree.Number{
		Value: value,
	}

	listener.stack.Push(node)
}

// An infix statement has been walked.
func (listener *Listener) ExitInfixStatement(c *parser.InfixStatementContext) {
	operator := c.GetOperator().GetText()
	right := listener.stack.Pop()
	left := listener.stack.Pop()

	var node tree.Node = tree.Operator{
		Operator: operator,
		Left:     left.(tree.Node),
		Right:    right.(tree.Node),
	}

	listener.stack.Push(node)
}
