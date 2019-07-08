package evaluators

type Compiler interface {
	checkPath() string
	checkVersion() string

	CompileSources()
	RaisedError() bool
	GetErrorHandler() *ErrorHandler
	RunExec() string

	AddFlags(flags ...string)
	AddLinks(links ...string)
	AddSources(sources ...string)
	AddArgs(args ...string)
	String() string
}
