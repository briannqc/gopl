package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	// Eval returns the value if this Expr in the environment Env.
	Eval(env Env) float64

	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error

	// String prints the Expr pretty.
	String() string
}

// A Var identifies a variable, e.g. x.
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (v Var) String() string {
	return strings.TrimSpace(string(v))
}

// A literal is a numeric constant, e.g. 3.141.
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (l literal) Check(_ map[Var]bool) error {
	return nil
}

func (l literal) String() string {
	return fmt.Sprintf("%v", float64(l))
}

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+' or '-'
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	default:
		panic(fmt.Errorf("unsupported unary op: %v", u.op))
	}
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unsupported unary op: %v", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) String() string {
	return fmt.Sprintf("%s%f", string(u.op), u.x)
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	x, y := b.x.Eval(env), b.y.Eval(env)
	switch b.op {
	case '+':
		return x + y
	case '-':
		return x - y
	case '*':
		return x * y
	case '/':
		return x / y
	default:
		panic(fmt.Errorf("unsupported binary op: %v", b.op))
	}
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unsupported binary op: %v", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	if err := b.y.Check(vars); err != nil {
		return err
	}
	return nil
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.x, string(b.op), b.y)
}

// A call represents a function call expression, e.g. sin(x) or pow(x, 2).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	default:
		panic(fmt.Errorf("unsupported function: %v", c.fn))
	}
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unsupported function: %v", c.fn)
	}

	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func (c call) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString(c.fn)
	buf.WriteString("(")
	buf.WriteString(c.args[0].String())
	for i := 1; i < len(c.args); i++ {
		buf.WriteString(", ")
		buf.WriteString(c.args[i].String())
	}
	buf.WriteString(")")
	return buf.String()
}

type postUnary struct {
	op rune // one of '!'
	x  Expr
}

func (p postUnary) Eval(env Env) float64 {
	x := p.x.Eval(env)
	val := 1.0
	for i := float64(2); i <= x; i++ {
		val *= i
	}
	return val
}

func (p postUnary) Check(vars map[Var]bool) error {
	if p.op != '!' {
		return fmt.Errorf("unsupported op: %v", string(p.op))
	}
	return p.x.Check(vars)
}

func (p postUnary) String() string {
	return fmt.Sprintf("%s!", p.x)
}

// An Env maps variables to their values.
type Env map[Var]float64

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %v", v)
		}
	}
	return expr, nil
}
