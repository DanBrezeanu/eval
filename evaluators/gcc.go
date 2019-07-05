package evaluators

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

type GccCompiler struct {
	Compiler
	exists  bool
	version string
	path    string

	flags   []string
	sources []string
	objects []string
	links   []string
	args    []string

	output string

	errorHandler *ErrorHandler
	waitGroup    *sync.WaitGroup
}

func NewGccCompiler() *GccCompiler {
	g := new(GccCompiler)
	g.waitGroup = &sync.WaitGroup{}

	g.checkPath()

	if g.errorHandler != nil {
		g.errorHandler.WriteToStderr()
		return nil
	}

	g.checkVersion()
	return g
}

func (g *GccCompiler) checkPath() string {
	cmd := "which gcc"
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		g.errorHandler = NewErrorHandler(NoCompilerFound, &err, string(out))
		return ""
	}
	g.path = string(out)
	return string(out)
}

func (g *GccCompiler) checkVersion() string {
	cmd := "gcc --version | head -n1"
	out, _ := exec.Command("bash", "-c", cmd).CombinedOutput()
	g.version = string(out)
	return string(out)
}

func (g *GccCompiler) AddFlags(flags ...string) {
	g.flags = append(g.flags, flags...)
}

func (g *GccCompiler) AddLinks(links ...string) {
	g.links = append(g.links, links...)
}

func (g *GccCompiler) AddSources(sources ...string) {
	g.sources = append(g.sources, sources...)
}

func (g *GccCompiler) AddArgs(args ...string) {
	g.args = append(g.args, args...)
}

func (g *GccCompiler) createObjectFiles(source string) {
	defer g.waitGroup.Done()
	cmd := "gcc -c " + source + " " + strings.Join(g.flags, " ")

	fmt.Println(cmd)
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil {
		if g.errorHandler != nil {
			g.errorHandler.errorOutput += string(out)
		} else {
			g.errorHandler = NewErrorHandler(CompileError, &err, string(out))
		}
	}

}

func (g *GccCompiler) CompileSources() {
	for _, source := range g.sources {
		g.waitGroup.Add(1)
		go g.createObjectFiles(source)
		g.objects = append(g.objects, strings.Replace(source, ".c", ".o", 1))
	}
	g.waitGroup.Wait()


	cmd := "gcc -o exec " +
		strings.Join(g.objects, " ") + " " +
		strings.Join(g.links, " ") + "\n"

	fmt.Println(cmd)
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil {
		if g.errorHandler != nil {
			g.errorHandler.errorOutput += string(out)
		} else {
			g.errorHandler = NewErrorHandler(CompileError, &err, string(out))
		}
		return
	}
}

func (g *GccCompiler) RunExec() string {
	cmd := "./exec " + strings.Join(g.args, " ")
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil {
		g.errorHandler = NewErrorHandler(RunTimeError, &err, string(out))
		return ""
	}

	return string(out)
}

func (g *GccCompiler) RaisedError() bool {
	return g.errorHandler != nil
}

func (g *GccCompiler) GetErrorHandler() *ErrorHandler {
	return g.errorHandler
}

func (g GccCompiler) String() string {
	outString := "VERSION: " + g.version
	outString += "   PATH: " + g.path
	outString += "SOURCES: " + strings.Join(g.sources, ";") + "\n"
	outString += "  FLAGS: " + strings.Join(g.flags, " ") + "\n"
	outString += "  LINKS: " + strings.Join(g.links, " ") + "\n"
	outString += "   ARGS: " + strings.Join(g.args, " ")

	return outString
}
