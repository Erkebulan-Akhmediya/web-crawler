package main

import "flag"

var maxDepth = flag.Int("depth", 1, "maximum depth of crawling")

func main() {
	flag.Parse()
	v := visitor{seen: make(map[string]bool), worklist: make(chan item)}

	v.n++
	go func() { v.initialVisit() }()
	v.visit()
}
