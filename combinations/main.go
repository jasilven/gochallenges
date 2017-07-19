package main

import "fmt"
import "os"

// Prints out all the possible values (combinations) that can be created by
// concatenating the given arguments together
func main() {
	prefix := []string{}
	recur(prefix, os.Args[1:])
}

func recur(pref, suf []string) {
	if len(suf) == 1 {
		for _, v := range pref {
			fmt.Print(v)
		}
		fmt.Println(suf[0])
	} else {
		s := make([]string, len(suf), len(suf))
		copy(s, suf)
		for i := range s {
			if i != 0 {
				tmp := s[0]
				s[0] = s[i]
				s[i] = tmp
			}
			recur(append(pref, s[0]), s[1:])
		}
	}
}
