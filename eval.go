package main

import (
	"fmt"

	"github.com/DanBrezeanu/eval/evaluators"
)

func main() {
	var x *evaluators.GccCompiler = evaluators.NewGccCompiler()
	z := new(error)
	var y *evaluators.ErrorHandler = evaluators.NewErrorHandler(evaluators.NoCompilerFound, *z)

	x.AddFlags("-Wall")
	x.AddSources("test.c")
	x.AddLinks("-lm")

	x.CompileSources()
	x.RunExec()
	fmt.Println(x, y)
}
