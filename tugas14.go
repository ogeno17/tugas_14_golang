package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var hasil, _ = exec.Command("ls").Output()
	fmt.Println(string(hasil))

	var hasil2, _ = exec.Command("date").Output()
	fmt.Println(string(hasil2))

	var hasil3, _ = exec.Command("who").Output()
	fmt.Println(string(hasil3))
}
