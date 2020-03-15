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

// Turns a mathematical expression into a tree.
package evaluator

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"piankalabs.com/math/pkg/parser"
	"piankalabs.com/math/pkg/tree"
)

// For evaluating expressions.
type Evaluator struct {

}

// Parses an expression and turns it into a tree.Node.
func (evaluator Evaluator) Evaluate(expression string) (tree.Node, error) {
	is := antlr.NewInputStream(expression)

	lexer := parser.NewMathLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	antlrParser := parser.NewMathParser(tokens)

	antlrParser.BuildParseTrees = true
	antlrTree := antlrParser.Expression()

	var listener Listener
	antlr.ParseTreeWalkerDefault.Walk(&listener, antlrTree)

	return listener.Tree, nil
}