package version

import (
	"fmt"
	"runtime"
)

// Application build information.
var (
	Branch    string
	BuildDate string
	GitSHA1   string
	Version   = "v0.0.3-dev"
)

// Print writes application version details to standard output.
func Print() {
	// TODO remove hard coded "gsd" string here
	fmt.Printf("gsd, version %v (branch: %v, revision: %v)\n", Version, Branch, GitSHA1)
	fmt.Println("build date:", BuildDate)
	fmt.Println("go version:", runtime.Version())
}
