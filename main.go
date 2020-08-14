package main

import (
	"fmt"
	"os"

	"cursmedia.com/rakuten/cmd"
)

func main() {
	fmt.Println("Starting rakuten import. Arguments {}", os.Args)

	switch os.Args[1] {
	case "--download":
		fmt.Println("Calling Download")
		cmd.Download()
		break
	case "--insert":
		fmt.Println("Calling Insert")
		cmd.Insert()
		break
	default:
		fmt.Println("Params should be either --download to connect and download, or --insert to insert to database")
	}
}
