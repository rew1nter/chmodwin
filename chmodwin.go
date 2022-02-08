package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type chmod struct {
	path string
	args []string
}

func main() {
	flag.Usage = func() {
		fmt.Println("This can 'chmod 600' files in Windows\nusage:\n\tchmodwin key.pem\n\tchmodwin.exe key.pem")
	}
	flag.Parse()
	if flag.NArg() > 1 {
		fmt.Println("ERR: You have to provide only one argument")
		os.Exit(1)
	} else if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filePath := flag.Args()[0]
	
	c1 := chmod{
		path: "icacls.exe",
		args: []string{filePath, "/reset"},
	}

	u := username()
	c2 := chmod{
		path: "icacls.exe",
		args: []string{filePath, "/GRANT:R", fmt.Sprintf("%v:(R)", u)},
	}

	c3 := chmod{
		path: "icacls.exe",
		args: []string{filePath, "/inheritance:r"},
	}
	execute(c1.path, c1.args)
	execute(c2.path, c2.args)
	execute(c3.path, c3.args)

}

func username() (x string) {
	for _, j := range os.Environ() {
		s := strings.Split(j, "=")
		if s[0] == "USERNAME" {
			x = s[1]
		}
	}
	return x
}

func execute(name string, args []string) {
	cmd := exec.Command(name, args...)

	cmd.Stdin = os.Stdin
	//cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
