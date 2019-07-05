package evaluators

type Compiler interface {
	checkPath() string
	checkVersion() string

	AddFlags(flags ...string)
	AddLinks(links ...string)
	AddSources(sources ...string)

	String() string
}
