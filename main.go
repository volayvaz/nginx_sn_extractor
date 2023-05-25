package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var ()

func main() {

	names := make(map[string]int)

	f, err := os.Open("nginx.cfg")
	check(err)
	scanner := bufio.NewScanner(f)
	inOpt := false
	for scanner.Scan() {
		inOpt = extractOptionValue(scanner.Text(), "server_name", inOpt, names)
	}
	fmt.Printf("Domain,Number of Occurrences\n")
	keys := make([]string, 0, len(names))

	for k, _ := range names {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%v,%v\n", k, names[k])

	}
}

func extractOptionValue(ln, option string, inOpt bool, ns map[string]int) bool {
	if !strings.HasPrefix(strings.TrimLeft(ln, " \t"), "#") {
		sub := fmt.Sprint(strings.Split(ln, "#")[0])
		s := strings.Fields(sub)
		for _, v := range s {
			if inOpt {

				if v[len(v)-1:] == ";" {
					v = v[:len(v)-1]
					inOpt = false
				}

				if len(v) > 0 {
					ns[v] = ns[v] + 1
				}
			}
			if v == option {
				inOpt = true
			}

		}
	}
	return inOpt
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
