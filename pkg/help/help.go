package help

import (
	_ "embed"
)

//go:embed README.md
var HelpText string
