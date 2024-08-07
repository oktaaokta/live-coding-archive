package internal

import "fmt"

func StringContainsWord() {
	s := "makan"
	p := "*aa*"
	fmt.Println(ContainsWord(s, p))
}

func ContainsWord(s, p string) bool {
	for len(p) > 0 && p[0] == '*' {
		p = p[1:]
	}

	for i := range s {
		if s[i] == p[0] {
			j, pointer := i+1, 1
			for pointer < len(p) {
				if p[pointer] == '*' {
					pointer++
					continue
				}
				if pointer == len(p) {
					return true
				}
				if s[j] == p[pointer] {
					j++
					pointer++
					continue
				}
				break
			}
			if pointer == len(p) {
				return true
			}
		}
	}

	return false
}
