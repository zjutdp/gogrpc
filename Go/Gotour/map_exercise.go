package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	sa := strings.Fields(s)
	fmt.Println(sa)
	for _, w := range sa {
		v, ok := m[w]
		if ok {
			m[w] = v + 1
		}else{
			m[w] = 1
		}
	}
	return m
}
