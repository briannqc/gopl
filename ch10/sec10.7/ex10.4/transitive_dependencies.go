package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

func main() {
	cmd := exec.Command("go", "list", "./...")
	b, err := cmd.Output()
	if err != nil {
		log.Println("Run cmd failed", cmd)
		os.Exit(1)
	}

	var mu sync.Mutex
	pkgAndDeps := map[string][]string{}

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		pkg := scanner.Text()

		wg.Add(1)
		go func(pkg string) {
			defer wg.Done()

			cmd := exec.Command("go", "list", "-json", pkg)
			b, err := cmd.Output()
			if err != nil {
				log.Println("Run cmd failed", cmd)
				return
			}
			var output goList
			err = json.Unmarshal(b, &output)
			if err != nil {
				log.Printf("Unmarshal cmd output failed, pkg: %v, err: %v\n", pkg, err)
				return
			}

			mu.Lock()
			pkgAndDeps[pkg] = output.Deps
			mu.Unlock()
		}(pkg)
	}

	wg.Wait()
	for _, libPkg := range os.Args[1:] {
		fmt.Printf("Packages depend (incl. transitively) on %v are:\n", libPkg)
		for pkg, deps := range pkgAndDeps {
			if contains(deps, libPkg) {
				fmt.Printf("\t%s\n", pkg)
			}
		}
		fmt.Println()
	}
}

type goList struct {
	Deps []string
}

func contains(all []string, element string) bool {
	for _, s := range all {
		if s == element {
			return true
		}
	}
	return false
}
