package main

import (
	"bufio"
	"io"
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

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(outfile)
	defer writer.Flush()

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go io.Copy(writer, stdoutPipe)
	cmd.Wait()
}
