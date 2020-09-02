package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PhamJamStudio/link"
)

func main() {
	htmlFile := flag.String("html", "ex4.html", "Source file for HTML link parsing")
	flag.Parse()

	// Don't want to convert to string, having readers is common
	r, err := os.Open(*htmlFile)
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	for _, link := range links {
		fmt.Printf("%#v\n", link)
	}

}

/* LEARNING NOTES
// Breaking down larger functions into smaller ones
	// User facing Parse()
		// html.Parse() -> html -> node w/ parsed tree
		// Get link nodes
		// Build Links from link nodes
// html.Node
	// anything inside a link node is a child of that node
	// DFS into node children via recursion
		// Have a base case(s), recurse thru children, then loop thru siblings

// General
	// fmt.Printf("%#v\n", link) -> prints type & val
	// ... -> return list as individual elements, or accept 0 or more params

*/
