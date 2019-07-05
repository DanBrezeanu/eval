package main

import (
	"fmt"
	"os"

	"github.com/DanBrezeanu/eval/evaluators"
)

func main() {
	var x *evaluators.GccCompiler = evaluators.NewGccCompiler()

	x.AddSources("test.c")
	x.AddFlags("-Wall")
	x.AddLinks("-lm")

	x.CompileSources()
	if hasRaised := x.RaisedError(); hasRaised {
		x.GetErrorHandler().WriteToStderr()
		os.Exit(1)
	}

	output := x.RunExec()
	fmt.Println(output)
	if hasRaised := x.RaisedError(); hasRaised {
		x.GetErrorHandler().WriteToStderr()
		os.Exit(1)
	}

	fmt.Println(x)
	os.Exit(0)
}
