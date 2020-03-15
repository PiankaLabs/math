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
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 38, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 18, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7,
	3, 26, 10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 2, 3, 4, 6, 2, 4, 6, 8, 2, 4, 3, 2, 7, 8, 3, 2, 5, 6, 2,
	37, 2, 10, 3, 2, 2, 2, 4, 17, 3, 2, 2, 2, 6, 30, 3, 2, 2, 2, 8, 34, 3,
	2, 2, 2, 10, 11, 5, 4, 3, 2, 11, 12, 7, 2, 2, 3, 12, 3, 3, 2, 2, 2, 13,
	14, 8, 3, 1, 2, 14, 18, 5, 6, 4, 2, 15, 18, 5, 8, 5, 2, 16, 18, 7, 9, 2,
	2, 17, 13, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 16, 3, 2, 2, 2, 18, 27,
	3, 2, 2, 2, 19, 20, 12, 5, 2, 2, 20, 21, 9, 2, 2, 2, 21, 26, 5, 4, 3, 6,
	22, 23, 12, 4, 2, 2, 23, 24, 9, 3, 2, 2, 24, 26, 5, 4, 3, 5, 25, 19, 3,
	2, 2, 2, 25, 22, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27,
	28, 3, 2, 2, 2, 28, 5, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 31, 7, 3, 2,
	2, 31, 32, 5, 4, 3, 2, 32, 33, 7, 4, 2, 2, 33, 7, 3, 2, 2, 2, 34, 35, 9,
	3, 2, 2, 35, 36, 5, 4, 3, 2, 36, 9, 3, 2, 2, 2, 5, 17, 25, 27,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'+'", "'-'", "'*'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "Operator_Add", "Operator_Subtract", "Operator_Multiply", "Operator_Divide",
	"Number", "Whitespace",
}

var ruleNames = []string{
	"expression", "statement", "parenthetical", "unary",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MathParser struct {
	*antlr.BaseParser
}

func NewMathParser(input antlr.TokenStream) *MathParser {
	this := new(MathParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Math.g4"

	return this
}

// MathParser tokens.
const (
	MathParserEOF               = antlr.TokenEOF
	MathParserT__0              = 1
	MathParserT__1              = 2
	MathParserOperator_Add      = 3
	MathParserOperator_Subtract = 4
	MathParserOperator_Multiply = 5
	MathParserOperator_Divide   = 6
	MathParserNumber            = 7
	MathParserWhitespace        = 8
)

// MathParser rules.
const (
	MathParserRULE_expression    = 0
	MathParserRULE_statement     = 1
	MathParserRULE_parenthetical = 2
	MathParserRULE_unary         = 3
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MathParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MathParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MathParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MathParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.statement(0)
	}
	{
		p.SetState(9)
		p.Match(MathParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MathParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaryStatementContext struct {
	*StatementContext
}

func NewUnaryStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryStatementContext {
	var p = new(UnaryStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *UnaryStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryStatementContext) Unary() IUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *UnaryStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterUnaryStatement(s)
	}
}

func (s *UnaryStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitUnaryStatement(s)
	}
}

type ParentheticalStatementContext struct {
	*StatementContext
}

func NewParentheticalStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalStatementContext {
	var p = new(ParentheticalStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ParentheticalStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalStatementContext) Parenthetical() IParentheticalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParentheticalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParentheticalContext)
}

func (s *ParentheticalStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterParentheticalStatement(s)
	}
}

func (s *ParentheticalStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitParentheticalStatement(s)
	}
}

type NumberStatementContext struct {
	*StatementContext
}

func NewNumberStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberStatementContext {
	var p = new(NumberStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *NumberStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberStatementContext) Number() antlr.TerminalNode {
	return s.GetToken(MathParserNumber, 0)
}

func (s *NumberStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterNumberStatement(s)
	}
}

func (s *NumberStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitNumberStatement(s)
	}
}

type InfixStatementContext struct {
	*StatementContext
	left     IStatementContext
	operator antlr.Token
	right    IStatementContext
}

func NewInfixStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InfixStatementContext {
	var p = new(InfixStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *InfixStatementContext) GetOperator() antlr.Token { return s.operator }

func (s *InfixStatementContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *InfixStatementContext) GetLeft() IStatementContext { return s.left }

func (s *InfixStatementContext) GetRight() IStatementContext { return s.right }

func (s *InfixStatementContext) SetLeft(v IStatementContext) { s.left = v }

func (s *InfixStatementContext) SetRight(v IStatementContext) { s.right = v }

func (s *InfixStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfixStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *InfixStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *InfixStatementContext) Operator_Multiply() antlr.TerminalNode {
	return s.GetToken(MathParserOperator_Multiply, 0)
}

func (s *InfixStatementContext) Operator_Divide() antlr.TerminalNode {
	return s.GetToken(MathParserOperator_Divide, 0)
}

func (s *InfixStatementContext) Operator_Add() antlr.TerminalNode {
	return s.GetToken(MathParserOperator_Add, 0)
}

func (s *InfixStatementContext) Operator_Subtract() antlr.TerminalNode {
	return s.GetToken(MathParserOperator_Subtract, 0)
}

func (s *InfixStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterInfixStatement(s)
	}
}

func (s *InfixStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitInfixStatement(s)
	}
}

func (p *MathParser) Statement() (localctx IStatementContext) {
	return p.statement(0)
}

func (p *MathParser) statement(_p int) (localctx IStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, MathParserRULE_statement, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MathParserT__0:
		localctx = NewParentheticalStatementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(12)
			p.Parenthetical()
		}

	case MathParserOperator_Add, MathParserOperator_Subtract:
		localctx = NewUnaryStatementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)
			p.Unary()
		}

	case MathParserNumber:
		localctx = NewNumberStatementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(14)
			p.Match(MathParserNumber)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(23)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewInfixStatementContext(p, NewStatementContext(p, _parentctx, _parentState))
				localctx.(*InfixStatementContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MathParserRULE_statement)
				p.SetState(17)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(18)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*InfixStatementContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MathParserOperator_Multiply || _la == MathParserOperator_Divide) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*InfixStatementContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(19)

					var _x = p.statement(4)

					localctx.(*InfixStatementContext).right = _x
				}

			case 2:
				localctx = NewInfixStatementContext(p, NewStatementContext(p, _parentctx, _parentState))
				localctx.(*InfixStatementContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, MathParserRULE_statement)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(21)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*InfixStatementContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MathParserOperator_Add || _la == MathParserOperator_Subtract) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*InfixStatementContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(22)

					var _x = p.statement(3)

					localctx.(*InfixStatementContext).right = _x
				}

			}

		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IParentheticalContext is an interface to support dynamic dispatch.
type IParentheticalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParentheticalContext differentiates from other interfaces.
	IsParentheticalContext()
}

type ParentheticalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParentheticalContext() *ParentheticalContext {
	var p = new(ParentheticalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MathParserRULE_parenthetical
	return p
}

func (*ParentheticalContext) IsParentheticalContext() {}

func NewParentheticalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParentheticalContext {
	var p = new(ParentheticalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathParserRULE_parenthetical

	return p
}

func (s *ParentheticalContext) GetParser() antlr.Parser { return s.parser }

func (s *ParentheticalContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ParentheticalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParentheticalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterParenthetical(s)
	}
}

func (s *ParentheticalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitParenthetical(s)
	}
}

func (p *MathParser) Parenthetical() (localctx IParentheticalContext) {
	localctx = NewParentheticalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MathParserRULE_parenthetical)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(MathParserT__0)
	}
	{
		p.SetState(29)
		p.statement(0)
	}
	{
		p.SetState(30)
		p.Match(MathParserT__1)
	}

	return localctx
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	operator antlr.Token
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MathParserRULE_unary
	return p
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MathParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) GetOperator() antlr.Token { return s.operator }

func (s *UnaryContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *UnaryContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *UnaryContext) Operator_Add() antlr.TerminalNode {
	return s.GetToken(MathParserOperator_Add, 0)
}

func (s *UnaryContext) Operator_Subtract() antlr.TerminalNode {
	return s.GetToken(MathParserOperator_Subtract, 0)
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MathListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (p *MathParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MathParserRULE_unary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*UnaryContext).operator = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == MathParserOperator_Add || _la == MathParserOperator_Subtract) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*UnaryContext).operator = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(33)
		p.statement(0)
	}

	return localctx
}

func (p *MathParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *StatementContext = nil
		if localctx != nil {
			t = localctx.(*StatementContext)
		}
		return p.Statement_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MathParser) Statement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
