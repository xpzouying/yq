package yqlib

import (
	"fmt"
	"strings"
)

var myPathTokeniser = newExpressionTokeniser()
var myPathPostfixer = newExpressionPostFixer()

type ExpressionNode struct {
	Operation *Operation
	Lhs       *ExpressionNode
	Rhs       *ExpressionNode
}

type ExpressionParser interface {
	ParseExpression(expression string) (*ExpressionNode, error)
}

type expressionParserImpl struct {
}

func NewExpressionParser() ExpressionParser {
	return &expressionParserImpl{}
}

func (p *expressionParserImpl) ParseExpression(expression string) (*ExpressionNode, error) {
	tokens, err := myPathTokeniser.Tokenise(expression)
	if err != nil {
		return nil, err
	}
	var Operations []*Operation
	Operations, err = myPathPostfixer.ConvertToPostfix(tokens)
	if err != nil {
		return nil, err
	}
	return p.createExpressionTree(Operations)
}

func (p *expressionParserImpl) createExpressionTree(postFixPath []*Operation) (*ExpressionNode, error) {
	var stack = make([]*ExpressionNode, 0)

	if len(postFixPath) == 0 {
		return nil, nil
	}

	for _, Operation := range postFixPath {
		var newNode = ExpressionNode{Operation: Operation}
		log.Debugf("pathTree %v ", Operation.toString())
		if Operation.OperationType.NumArgs > 0 {
			numArgs := Operation.OperationType.NumArgs
			if numArgs == 1 {
				if len(stack) < 1 {
					return nil, fmt.Errorf("'%v' expects 1 arg but received none", strings.TrimSpace(Operation.StringValue))
				}
				remaining, rhs := stack[:len(stack)-1], stack[len(stack)-1]
				newNode.Rhs = rhs
				stack = remaining
			} else if numArgs == 2 {
				if len(stack) < 2 {
					return nil, fmt.Errorf("'%v' expects 2 args but there is %v", strings.TrimSpace(Operation.StringValue), len(stack))
				}
				remaining, lhs, rhs := stack[:len(stack)-2], stack[len(stack)-2], stack[len(stack)-1]
				newNode.Lhs = lhs
				newNode.Rhs = rhs
				stack = remaining
			}
		}
		stack = append(stack, &newNode)
	}
	if len(stack) != 1 {
		return nil, fmt.Errorf("expected end of expression but found '%v', please check expression syntax", strings.TrimSpace(stack[1].Operation.StringValue))
	}
	return stack[0], nil
}
