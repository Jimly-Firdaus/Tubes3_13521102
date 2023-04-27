package FeatureCalculator

import (
	"TUBES3_13521102/netlify/functions/main/Class"
	"fmt"
	"strconv"
	"strings"
)


func Precedence(op string) int {
  if (op == "+" || op == "-") {
    return 1
  } else if (op == "*" || op == "/") {
    return 2
  }
  return 0
}

// Get all the number in the string
func GetNumber(expression string) []string {
  number := strings.FieldsFunc(expression, func(c rune) bool {
    return c == '-' || c == '+' || c == ' ' || c == '/' || c == '*' || c == '(' || c == ')'
  })
  return number
}

func StringtoFloat(number string) float64 {
  f, _ := strconv.ParseFloat(number, 64)
  return f
}

func applyOp(a float64, b float64, op string) (string, error) {
  switch op {
  case "+":
    return strconv.FormatFloat(a + b, 'f', 2, 64), nil
  case "-":
    return strconv.FormatFloat(a - b, 'f', 2, 64), nil
  case "*":
    return strconv.FormatFloat(a * b, 'f', 2, 64), nil
  case "/":
    if (b == 0) {
      return "", fmt.Errorf("division by zero")
    }
    return strconv.FormatFloat(a / b, 'f', 2, 64), nil
  default:
    return "", fmt.Errorf("invalid operator")
  }
}

func CalculateExpression(expression string) (string, error) {
  // Stack to store all operator
  operatorStack := Class.Stack{}

  // Stack to store all values
  valueStack := Class.Stack{}

  // Get all numbers in the expression
  numbers := GetNumber(expression)

  // Boolean if a number is found
  numberFound := false

  for i := 0; i < len(expression); i++ {
    // If a number is Found iterate until all the number sequence in is skipped
    if (numberFound) {
      if ((expression[i] < '0' || expression[i] > '9') && expression[i] != '.') {
        numberFound = false
      } else {
        continue
      }
    }

    // If a number is found push the number to stack
    if (expression[i] >= '0' && expression[i] <= '9') {
      valueStack.Push(numbers[0])
      numbers = numbers[1:]
      numberFound = true
      continue
    }

    // If current char is whitespace then skip
    if (expression[i] == ' ') {
      continue
    } else if (expression[i] == '(') { // If current token is an opening parenthesis push the token
      operatorStack.Push(string(expression[i]))
      if (i != len(expression)-1){ // If '(' is not the last character in expression and a minus exist after this expression
        if (expression[i+1] == '-'){ // if a minus appears after opening paranthesis, push a 0 number value.
          valueStack.Push("0")
        }
      }

    } else if (expression[i] == ')') { // If current token is an closing parenthesis then solve expression until opening parenthesis is found
      for {

        // If the expression contains parenthesis that contains nothing then return error
        if (expression[i-1] == '(') {
          return "", fmt.Errorf("Invalid Syntax")
        }

        if (operatorStack.IsEmpty()) {
          break
        }

        // All err is for handling Syntax error

        op, _ := operatorStack.Peek()

        if (op == "(") {
          break
        }
        firstNumber, err := valueStack.Pop()

        if (err != nil) {
          return "", err
        }

        secondNumber, err := valueStack.Pop()

        if (err != nil) {
          return "", err
        }

        operator, err := operatorStack.Pop()

        if (err != nil) {
          return "", err
        }

        hasil, err := applyOp(StringtoFloat(secondNumber), StringtoFloat(firstNumber), operator)

        if (err != nil) {
          return "", err
        }

        valueStack.Push(hasil)
      }

      // If opening parenthesis is found then pop the op
      _, err := operatorStack.Pop()

      if (err != nil) {
        return "", err
      }

    } else { // If tokens for mathematical operations is found then
      for {

        // If operatorStack is empty there is no operator to compare so break
        if (operatorStack.IsEmpty()) {
          break
        }

        opTop, _ := operatorStack.Peek()

        // If the operator on top of operatorStack authority is below current token then just push current token to stack
        if (Precedence(opTop) < Precedence(string(expression[i]))) {
          break
        }

        // If the operator on top of operatorStack is more important than the current token then process the expression first
        firstNumber, err := valueStack.Pop()

        if (err != nil) {
          return "", err
        }

        secondNumber, err := valueStack.Pop()

        if (err != nil) {
          return "", err
        }

        operator, err := operatorStack.Pop()

        if (err != nil) {
          return "", err
        }

        hasil, err := applyOp(StringtoFloat(secondNumber), StringtoFloat(firstNumber), operator)

        if (err != nil) {
          return "", err
        }

        valueStack.Push(hasil)
      }
      // Pushing current token operator to stack
      operatorStack.Push(string(expression[i]))
    }
  }

  for {
    // Finishing the rest of the mathematical expression
    if (operatorStack.IsEmpty()) {
      break
    }

    firstNumber, err := valueStack.Pop()

    if (err != nil) {
      return "", err
    }

    secondNumber, err := valueStack.Pop()

    if (err != nil) {
      return "", err
    }

    operator, err := operatorStack.Pop()

    if (err != nil) {
      return "", err
    }

    hasil, err := applyOp(StringtoFloat(secondNumber), StringtoFloat(firstNumber), operator)

    if (err != nil) {
      return "", err
    }
    valueStack.Push(hasil)
  }
  return valueStack.Peek()
}
