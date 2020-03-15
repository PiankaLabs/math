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

// BaseMathListener is a complete listener for a parse tree produced by MathParser.
type BaseMathListener struct{}

var _ MathListener = &BaseMathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMathListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMathListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUnaryStatement is called when production unaryStatement is entered.
func (s *BaseMathListener) EnterUnaryStatement(ctx *UnaryStatementContext) {}

// ExitUnaryStatement is called when production unaryStatement is exited.
func (s *BaseMathListener) ExitUnaryStatement(ctx *UnaryStatementContext) {}

// EnterParentheticalStatement is called when production parentheticalStatement is entered.
func (s *BaseMathListener) EnterParentheticalStatement(ctx *ParentheticalStatementContext) {}

// ExitParentheticalStatement is called when production parentheticalStatement is exited.
func (s *BaseMathListener) ExitParentheticalStatement(ctx *ParentheticalStatementContext) {}

// EnterNumberStatement is called when production numberStatement is entered.
func (s *BaseMathListener) EnterNumberStatement(ctx *NumberStatementContext) {}

// ExitNumberStatement is called when production numberStatement is exited.
func (s *BaseMathListener) ExitNumberStatement(ctx *NumberStatementContext) {}

// EnterInfixStatement is called when production infixStatement is entered.
func (s *BaseMathListener) EnterInfixStatement(ctx *InfixStatementContext) {}

// ExitInfixStatement is called when production infixStatement is exited.
func (s *BaseMathListener) ExitInfixStatement(ctx *InfixStatementContext) {}

// EnterParenthetical is called when production parenthetical is entered.
func (s *BaseMathListener) EnterParenthetical(ctx *ParentheticalContext) {}

// ExitParenthetical is called when production parenthetical is exited.
func (s *BaseMathListener) ExitParenthetical(ctx *ParentheticalContext) {}

// EnterUnary is called when production unary is entered.
func (s *BaseMathListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production unary is exited.
func (s *BaseMathListener) ExitUnary(ctx *UnaryContext) {}
