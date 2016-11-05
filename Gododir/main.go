package main

import (
	do "gopkg.in/godo.v2"
	"strings"
)

func tasks(p *do.Project) {
	p.Task("default", do.S{"guard"}, nil)

	p.Task("guard", nil, func(c *do.Context) {
		if c.FileEvent != nil {
			path := c.FileEvent.Path
			if strings.Contains(path, "_test.go") {
				c.Run("go test " + path)
			} else if strings.Contains(path, ".go") {
				path = strings.Replace(path, ".go", "_test.go", -1)
				c.Run("go test " + path)
			}
		}
	}).Src("**/*.go")

}

func main() {
	do.Godo(tasks)
}
