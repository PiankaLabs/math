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

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 48, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 33, 10, 8, 13, 8, 14, 8, 34, 3,
	8, 3, 8, 6, 8, 39, 10, 8, 13, 8, 14, 8, 40, 5, 8, 43, 10, 8, 3, 9, 3, 9,
	3, 9, 3, 9, 2, 2, 10, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 3, 2, 4, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 50, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19,
	3, 2, 2, 2, 5, 21, 3, 2, 2, 2, 7, 23, 3, 2, 2, 2, 9, 25, 3, 2, 2, 2, 11,
	27, 3, 2, 2, 2, 13, 29, 3, 2, 2, 2, 15, 32, 3, 2, 2, 2, 17, 44, 3, 2, 2,
	2, 19, 20, 7, 42, 2, 2, 20, 4, 3, 2, 2, 2, 21, 22, 7, 43, 2, 2, 22, 6,
	3, 2, 2, 2, 23, 24, 7, 45, 2, 2, 24, 8, 3, 2, 2, 2, 25, 26, 7, 47, 2, 2,
	26, 10, 3, 2, 2, 2, 27, 28, 7, 44, 2, 2, 28, 12, 3, 2, 2, 2, 29, 30, 7,
	49, 2, 2, 30, 14, 3, 2, 2, 2, 31, 33, 9, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33,
	34, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 42, 3, 2, 2,
	2, 36, 38, 7, 48, 2, 2, 37, 39, 9, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39, 40,
	3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 43, 3, 2, 2, 2,
	42, 36, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 16, 3, 2, 2, 2, 44, 45, 9,
	3, 2, 2, 45, 46, 3, 2, 2, 2, 46, 47, 8, 9, 2, 2, 47, 18, 3, 2, 2, 2, 6,
	2, 34, 40, 42, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'+'", "'-'", "'*'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "Operator_Add", "Operator_Subtract", "Operator_Multiply", "Operator_Divide",
	"Number", "Whiteosace",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "Operator_Add", "Operator_Subtract", "Operator_Multiply",
	"Operator_Divide", "Number", "Whiteosace",
}

type MathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMathLexer(input antlr.CharStream) *MathLexer {

	l := new(MathLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Math.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MathLexer tokens.
const (
	MathLexerT__0              = 1
	MathLexerT__1              = 2
	MathLexerOperator_Add      = 3
	MathLexerOperator_Subtract = 4
	MathLexerOperator_Multiply = 5
	MathLexerOperator_Divide   = 6
	MathLexerNumber            = 7
	MathLexerWhiteosace        = 8
)
