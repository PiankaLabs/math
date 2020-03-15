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

// Code generated from grammar/Math.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Math
import "github.com/antlr/antlr4/runtime/Go/antlr"

// MathListener is a complete listener for a parse tree produced by MathParser.
type MathListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterUnaryStatement is called when entering the unaryStatement production.
	EnterUnaryStatement(c *UnaryStatementContext)

	// EnterParentheticalStatement is called when entering the parentheticalStatement production.
	EnterParentheticalStatement(c *ParentheticalStatementContext)

	// EnterNumberStatement is called when entering the numberStatement production.
	EnterNumberStatement(c *NumberStatementContext)

	// EnterInfixStatement is called when entering the infixStatement production.
	EnterInfixStatement(c *InfixStatementContext)

	// EnterParenthetical is called when entering the parenthetical production.
	EnterParenthetical(c *ParentheticalContext)

	// EnterUnary is called when entering the unary production.
	EnterUnary(c *UnaryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitUnaryStatement is called when exiting the unaryStatement production.
	ExitUnaryStatement(c *UnaryStatementContext)

	// ExitParentheticalStatement is called when exiting the parentheticalStatement production.
	ExitParentheticalStatement(c *ParentheticalStatementContext)

	// ExitNumberStatement is called when exiting the numberStatement production.
	ExitNumberStatement(c *NumberStatementContext)

	// ExitInfixStatement is called when exiting the infixStatement production.
	ExitInfixStatement(c *InfixStatementContext)

	// ExitParenthetical is called when exiting the parenthetical production.
	ExitParenthetical(c *ParentheticalContext)

	// ExitUnary is called when exiting the unary production.
	ExitUnary(c *UnaryContext)
}
