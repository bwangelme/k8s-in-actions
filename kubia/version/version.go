package version

import "fmt"

const Version = "0.1.0"

var (
	Name         = "QAE Kubia"
	HumanVersion = fmt.Sprintf("%s v%s", Name, Version)
)
