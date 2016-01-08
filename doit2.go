package main

import (
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("/bin/echo", "Why hallo thar")

	outfile, err := os.Create("./out.txt")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	cmd.Stdout = outfile

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	cmd.Wait()
}
