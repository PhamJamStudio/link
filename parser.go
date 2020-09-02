/* DESIGN - define API, which func users like to call
// Open and read HTML file into string using io.reader
	// Use parse, get node, traverse child nodes DFS
// Extract all links <a href="">...</a>
// For each extracted link, return struct including href & text
	// Can strip HTML in link, extra whitespaces, newlines, etc
	// Ignore nested link
// Approx struct output
	Link {
		Href: "/dog",
		Text: "Something in a span Text not in a span Bold text!",
	}
// REF:
	// https://godoc.org/golang.org/x/net/html
	// Hint: See NodeType constants and look for the types that you can ignore.


*/

// Package link contains functions to parse and extract HTML links from an HTML document
package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...") in an HTML doc
type Link struct {
	Href string
	Text string
}

// Parse takes an HTML document, get parse tree, DFS to extract link nodes, return list of Links
func Parse(htmlDoc io.Reader) ([]Link, error) {
	// Get parse tree from HTML file to DFS thru
	doc, err := html.Parse(htmlDoc)
	if err != nil {
		return nil, err
	}
	// Get nodes that are links
	nodes := linkNodes(doc)
	// Create Links with link nodes
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, err
}

// DFS HTML parse tree to find and return a list of link nodes
func linkNodes(n *html.Node) []*html.Node {
	// base case, node is a link
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	// TODO: Don't add links that already exist
	// DFS to find link nodes
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Append individual links from linkNodes
		ret = append(ret, linkNodes(c)...) // '...' breaks out list into individual links
	}
	// Potentially return nothing
	return ret
}

// Assume <a> node, extracts href link val and text and creates a Link
func buildLink(n *html.Node) Link {
	var ret Link

	// we already know this is a link node, so we can just extract the link val directly
	// ret.Href = n.Attr[0].Val

	// get href out of node
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	// DFS inside link node to find all text nodes, ignoring non text
	ret.Text = text(n)

	return ret
}

// Extract text from html node, ignore comments, <strong> etc. NOTE: text/etc nodes within a link node are children node of link, hence we can DFS
func text(n *html.Node) string {
	// Base case
	if n.Type == html.TextNode {
		return n.Data
	}
	// DFS needed because formatting nodes can be nested inside text nodes e.g hello <strong>world</strong>!
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " " // TODO: Use byte buffer to be more efficient
	}
	return strings.Join(strings.Fields(ret), " ")
}
