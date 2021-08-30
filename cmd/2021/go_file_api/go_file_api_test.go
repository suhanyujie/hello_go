package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFile1(t *testing.T) {
	file, err := os.Stat("D:\\tech\\repo\\golang\\zinxDemo1\\notes\\2021\\go_file_api\\go_file_api.go")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%v", file)
}

func TestFileMode1(t *testing.T) {

}
