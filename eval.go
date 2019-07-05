package main

import (
	"fmt"
	"os"

	"github.com/DanBrezeanu/eval/evaluators"
	"github.com/DanBrezeanu/eval/gui"
)

func main() {
	var x *evaluators.GccCompiler = evaluators.NewGccCompiler()

	gui.MainGui()

	x.AddSources("test.c", "test2.c")
	x.AddFlags("-Wall")
	x.AddLinks("-lncurses")

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
