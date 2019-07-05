package evaluators

import (
	"fmt"
	"os/exec"
	"strings"
)

type GccCompiler struct {
	Compiler
	exists  bool
	version string
	path    string

	flags   []string
	sources []string
	links   []string
	args    []string

	output string

	errorHandler ErrorHandler
}

func NewGccCompiler() *GccCompiler {
	g := new(GccCompiler)
	g.checkVersion()
	g.checkPath()
	return g
}

func (g *GccCompiler) checkPath() string {
	cmd := "which gcc"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		//TODO:
	}
	g.path = string(out)
	return string(out)
}

func (g *GccCompiler) checkVersion() string {
	cmd := "gcc --version | head -n1"
	out, _ := exec.Command("bash", "-c", cmd).Output()
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

func (g *GccCompiler) CompileSources() {
	cmd := "gcc -o exec " +
		strings.Join(g.flags, " ") + " " +
		strings.Join(g.sources, " ") + " " +
		strings.Join(g.links, " ")

	fmt.Println(cmd)
	out, err := exec.Command("bash", "-c", cmd).Output()

	fmt.Println(out)
	if err != nil {

	}
}

func (g *GccCompiler) RunExec() {
	cmd := "./exec " + strings.Join(g.args, " ")
	out, err := exec.Command("bash", "-c", cmd).Output()

	fmt.Println(string(out))
	if err != nil {

	}
}

func (g GccCompiler) String() string {
	outString := "VERSION: " + g.version
	outString += "   PATH: " + g.path
	outString += "SOURCES: " + strings.Join(g.sources, ";") + "\n"
	outString += "  FLAGS: " + strings.Join(g.flags, " ") + "\n"
	outString += "  LINKS: " + strings.Join(g.links, " ") + "\n"

	return outString
}