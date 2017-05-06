package main

import (
	"fmt"
	"os"
)

var cos *COS

func main() {
	if len(os.Args) > 1 {
		cos = NewCosClient()
		switch os.Args[1] {
		case "get":
			get()
		case "put":
			put()
		case "version":
			version()
		default:
			usage()
		}
	} else {
		help()
	}
}

func put() {
	fmt.Println(cos)
}

func get() {
	fmt.Println(cos)

}

func version() {
	fmt.Println(os.Args[0] + ": " + Config.Name + "/" + Config.Version)
}

func help() {
	version()

}

func usage() {
	fmt.Println(os.Args[0] + " get file")
	fmt.Println(os.Args[0] + " put file")
}
