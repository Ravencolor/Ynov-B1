package main

import (
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		} else {
			os.Stdout.Write(data)
		}
	} else {
		for _, s := range args {
			file, err := os.Open(s)
			if err != nil {
				os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
				break
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
					os.Exit(1)
					break
				} else {
					os.Stdout.Write(data)
				}
			}
		}
	}
}
