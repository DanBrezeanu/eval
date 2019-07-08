package evaluators

type Compiler interface {
	checkPath() string
	checkVersion() string

	CompileSources()
	RunExec() string

	RaisedError() bool
	GetErrorHandler() *ErrorHandler
	EraseErrorHandler()

	GetName() string

	AddFlags(flags ...string)
	AddLinks(links ...string)
	AddSources(sources ...string)
	AddArgs(args ...string)
	String() string
}
