package main

import (
	"fmt"

	// HACK ALERT
	// The correct import is this:
	//   "github.com/laramiel/bazel-example-golang-remote/remote"
	// But the bazel build rules for subversion packages don't support alternate
	// go_prefix() commands.
	// We hack around it by importing where those dependencies end up:
	"github.com/laramiel/bazel-example-golang/submodule/src/remote/remote"
)

func main() {
	fmt.Println("Hello", remote.World())
}
