package main

import (
	"fmt"
	"testing"

	"github.com/Knetic/govaluate"
	"github.com/expr-lang/expr"
)

func BenchmarkGovaluate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		expression, _ := govaluate.NewEvaluableExpression("10 > 0")
		expression.Evaluate(nil)
	}
}

func BenchmarkExpr(b *testing.B) {
}

func TestName(t *testing.T) {
	env := map[string]interface{}{
		"v1": 5,
		"v2": 2,
		"v3": 2,
		"v4": 2,
		"v4": 2,
	}
	program, err := expr.Compile(`(v1>v2 || v2 < v3) && v3==v4`, expr.Env(env))
	if err != nil {
		fmt.Println(err)
		return
	}

	v, err := expr.Run(program, env)
	fmt.Println(v, err)
}
