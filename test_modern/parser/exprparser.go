package parser

import (
	"github.com/modern-go/parse"
	"github.com/modern-go/parse/discard"
	"github.com/modern-go/parse/read"
)

const precedenceAssignment = 1
const precedenceConditional = 2
const precedenceSum = 3
const precedenceProduct = 4
const precedenceExponent = 5
const precedencePrefix = 6
const precedencePostfix = 7
const precedenceCall = 8

type exprLexer struct {
	value    *valueToken
	plus     *plusToken
	minus    *minusToken
	multiply *multiplyToken
	divide   *divideToken
	group    *groupToken
}

var expr = NewExprLexer()

func NewExprLexer() *exprLexer {
	return &exprLexer{}
}

func (lexer *exprLexer) Parse(src *parse.Source, precedence int) interface{} {
	return parse.Parse(src, lexer, precedence)
}

func (lexer *exprLexer) InfixToken(src *parse.Source) (parse.InfixToken, int) {
	discard.UnicodeSpace(src)
	switch src.Peek1() {
	case '+':
		return lexer.plus, precedenceSum
	case '-':
		return lexer.minus, precedenceSum
	case '*':
		return lexer.multiply, precedenceProduct
	case '/':
		return lexer.divide, precedenceProduct
	default:
		return nil, 0
	}
}

func (lexer *exprLexer) PrefixToken(src *parse.Source) parse.PrefixToken {
	discard.UnicodeSpace(src)
	switch src.Peek1() {
	case '(':
		return lexer.group
	case '-':
		return lexer.minus
	default:
		return lexer.value
	}
}

type valueToken struct {
}

func (token *valueToken) PrefixParse(src *parse.Source) interface{} {
	return read.Int(src)
}

type plusToken struct {
}

func (token *plusToken) InfixParse(src *parse.Source, left interface{}) interface{} {
	leftValue := left.(int)
	discard.UnicodeSpace(src)
	src.Consume1('+')
	discard.UnicodeSpace(src)
	rightValue := expr.Parse(src, precedenceSum).(int)
	return leftValue + rightValue
}

type minusToken struct {
}

func (token *minusToken) PrefixParse(src *parse.Source) interface{} {
	discard.UnicodeSpace(src)
	src.Consume1('-')
	expr := expr.Parse(src, precedencePrefix).(int)
	return -expr
}

func (token *minusToken) InfixParse(src *parse.Source, left interface{}) interface{} {
	leftValue := left.(int)
	discard.UnicodeSpace(src)
	src.Consume1('-')
	discard.UnicodeSpace(src)
	rightValue := expr.Parse(src, precedenceSum).(int)
	return leftValue - rightValue
}

type multiplyToken struct {
}

func (token *multiplyToken) InfixParse(src *parse.Source, left interface{}) interface{} {
	leftValue := left.(int)
	discard.UnicodeSpace(src)
	src.Consume1('*')
	discard.UnicodeSpace(src)
	rightValue := expr.Parse(src, precedenceProduct).(int)
	return leftValue * rightValue
}

type divideToken struct {
}

func (token *divideToken) InfixParse(src *parse.Source, left interface{}) interface{} {
	leftValue := left.(int)
	discard.UnicodeSpace(src)
	src.Consume1('/')
	discard.UnicodeSpace(src)
	rightValue := expr.Parse(src, precedenceProduct).(int)
	return leftValue / rightValue
}

type groupToken struct {
}

func (token *groupToken) PrefixParse(src *parse.Source) interface{} {
	discard.UnicodeSpace(src)
	src.Consume1('(')
	discard.UnicodeSpace(src)
	expr := expr.Parse(src, 0)
	discard.UnicodeSpace(src)
	src.Consume1(')')
	discard.UnicodeSpace(src)
	return expr
}
