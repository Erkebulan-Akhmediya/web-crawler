package main

import (
	"flag"
	"fmt"
)

type item struct {
	links []string
	depth int
}

type visitor struct {
	// seen keeps track of visited pages
	seen map[string]bool
	// n keeps track of number of sends to the worklist that are yet to occur
	n int
	// chanel to sends links between crawler and main goroutines
	worklist chan item
}

func (v *visitor) initialVisit() {
	v.worklist <- item{links: flag.Args(), depth: 1}
}

// the function keeps visiting until all links have been visited
func (v *visitor) visit() {
	for ; v.n > 0; v.n-- {
		v.visitItem()
	}
}

// the function visits each link from a list coming from the worklist
// it stops visiting once it reaches max depth
func (v *visitor) visitItem() {
	list := <-v.worklist
	currentDepth := list.depth
	for _, link := range list.links {
		if v.seen[link] {
			continue
		}

		v.seen[link] = true
		if currentDepth == *maxDepth {
			fmt.Println(link)
			continue
		}

		v.n++
		go func(link string) {
			v.worklist <- item{links: crawl(link), depth: currentDepth + 1}
		}(link)
	}
}
