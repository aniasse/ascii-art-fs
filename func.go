package main

import "fmt"

func Match(r rune, i int, ascii map[byte][]string) {
	for ind, v := range ascii {
		if rune(ind) == r {
			fmt.Print(v[i])
		}
	}
}

func NewLine(tab []string) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] != "" {
			return false
		}
	}
	return true
}

func Printable(tab []rune) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] < 32 || tab[i] > 126 {
			return false
		}
	}
	return true
}

func Banner(s string) string {
	return "./" + s + ".txt"
}
