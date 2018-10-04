package main

import "github.com/yicat/p2s/router"

func main() {
	r := router.Load()
	r.Run(":9021")
}
