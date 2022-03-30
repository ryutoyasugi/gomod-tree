package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const packageSeparator = "@"

type Package struct {
	name    string
	version string
}

// ex. github.com/spf13/cobra@v1.4.0
func newPackage(s string) Package {
	split := strings.Split(s, packageSeparator)
	if len(split) == 1 {
		return Package{split[0], ""}
	}
	return Package{split[0], split[1]}
}

func (p Package) isMain() bool {
	if p.version == "" {
		return true
	}
	return false
}

func (p Package) print() {
	var fullName = p.name
	if p.version != "" {
		fullName = fullName + packageSeparator + p.version
	}
	fmt.Println("-", fullName)
}

// ex. github.com/spf13/cobra@v1.4.0 github.com/spf13/pflag@v1.0.5
func newPackagePair(s string) (parent, child Package) {
	split := strings.Split(s, " ")
	return newPackage(split[0]), newPackage(split[1])
}

type Dependencies = map[Package][]Package

func readGoModGraph() (Dependencies, Package) {
	out, err := exec.Command("go", "mod", "graph").Output()
	if err != nil {
		cobra.CheckErr(errors.New("failed to run `go mod graph`"))
	}

	trim := strings.TrimSuffix(string(out), "\n")
	lines := strings.Split(trim, "\n")
	var d = Dependencies{}
	var main Package
	for _, line := range lines {
		parent, child := newPackagePair(line)
		if parent.isMain() {
			main = parent
		}
		if children, ok := d[parent]; ok {
			d[parent] = append(children, child)
		} else {
			d[parent] = []Package{child}
		}
	}
	return d, main
}

func printIndent(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("    ")
	}
}

func printDependencyTree(d Dependencies, parent Package, depth, indent int) {
	if parent.isMain() {
		parent.print()
	}
	if indent > depth {
		return
	}
	for _, child := range d[parent] {
		printIndent(indent)
		child.print()

		if len(d[child]) != 0 {
			printDependencyTree(d, child, depth, indent+1)
		}
	}
}

func PrintDependencies(depth int) {
	d, main := readGoModGraph()
	printDependencyTree(d, main, depth, 1)
}
