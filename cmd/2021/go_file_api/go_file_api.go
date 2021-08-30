package main

import (
	"fmt"
	"os"
)

func main() {
	execShowFileStat()
}

// go run go_file_api.go /etc/hosts
func execShowFileStat() {
	arg := os.Args[1]
	fd, _ := os.Stat(arg)
	fmt.Printf("file %s mod is: %o", arg, fd.Mode()&os.ModePerm)
}
