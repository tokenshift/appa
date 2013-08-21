package appa

import "strconv"
import "strings"
import "testing"

func Test_SimpleCalculator(t *testing.T) {
	input := CreateStringBuffer(strings.NewReader("1+2+3+4+5"))

	g := NewGrammar()

	oper := Lit("+")
	num := Regex("\\d+")

	expr := g.NonTerminal("EXPR")
	expr.AddReduction(Seq(num, oper, expr), func(matched []Node) Node {
		a, _ := strconv.ParseInt(matched[0].String(), 10, 0)
		b, _ := strconv.ParseInt(matched[2].String(), 10, 0)
		return Int(a + b)
	})

	expr.AddReduction(num, func(matched []Node) Node {
		val, _ := strconv.ParseInt(matched[0].String(), 10, 0)
		return Int(val)
	})

	result, err := expr.Parse(input)

	if err != nil {
		t.Error(err)
		return
	}

	val, ok := result[0].(Int)
	if !ok {
		t.Errorf("Expected an integer result.")
		t.Errorf("Got: %v", result)
		return
	}

	assertIntEquals(t, 15, int(val))
}

func Test_CalcWithAssociativity(t *testing.T) {
	input := CreateStringBuffer(strings.NewReader("1+2*3-4/5"))

	// (1 + ((2 * 3) - (4 / 5))) == 6.2

	g := NewGrammar()

	num := Regex("\\d+")

	op1 := g.NonTerminal("OP1")
	op1.AddReduction(Lit("+"), func(op []Node) Node { return op[0] })
	op1.AddReduction(Lit("-"), func(op []Node) Node { return op[0] })

	op2 := g.NonTerminal("OP2")
	op2.AddReduction(Lit("*"), func(op []Node) Node { return op[0] })
	op2.AddReduction(Lit("/"), func(op []Node) Node { return op[0] })

	val := g.NonTerminal("VAL")
	val.AddReduction(Seq(num, op2, val), func(matched []Node) Node {
		a, _ := strconv.ParseFloat(matched[0].String(), 64)
		b, _ := strconv.ParseFloat(matched[2].String(), 64)

		var val Float

		if matched[1].String() == "*" {
			val = Float(a * b)
		} else {
			val = Float(a / b)
		}

		return val
	})
	val.AddReduction(num, func(matched []Node) Node {
		val, _ := strconv.ParseFloat(matched[0].String(), 64)
		return Float(val)
	})

	expr := g.NonTerminal("EXPR")
	expr.AddReduction(Seq(val, op1, expr), func(matched []Node) Node {
		a := matched[0].(Float)
		b := matched[2].(Float)

		if matched[1].String() == "+" {
			return a + b
		} else {
			return a - b
		}
	})
	expr.AddReduction(val, func(matched []Node) Node {
		return matched[0].(Float)
	})

	math := g.NonTerminal("MATH")
	math.AddRule(expr)

	result, _ := expr.Parse(input)
	assertFloatEquals(t, 6.2, float64(result[0].(Float)))
}
