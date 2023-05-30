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

	names := make(map[string][]string)

	f, err := os.Open("nginx.cfg")
	check(err)
	scanner := bufio.NewScanner(f)
	inOpt := false
	for scanner.Scan() {
		inOpt = extractOptionValue(scanner.Text(), "server_name", inOpt, names)
	}
	//fmt.Printf("Domain,Number of Occurrences\n")
	keys := make([]string, 0, len(names))

	for k, _ := range names {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%v\n", k)
		for v := range names[k] {
			fmt.Printf("\t\t%v\n", names[k][v])
		}

	}
}

func extractOptionValue(ln, option string, inOpt bool, ns map[string][]string) bool {
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
					sld := ""
					subs := strings.Split(v, ".")
					if len(subs) > 1 {
						sld = strings.Join(subs[len(subs)-2:], ".")
					} else {
						sld = strings.Join(subs, ".")
					}
					isSeen := false
					for j := range ns[sld] {
						if ns[sld][j] == v {
							isSeen = true
						}
					}
					if !isSeen {
						ns[sld] = append(ns[sld], v)
					}
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
