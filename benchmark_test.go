package main

import (
	"testing"
	"fmt"
	"strings"
)

func BenchmarkParse(b *testing.B) { 
	commands := []string{
		"print",
		"palindrom",
	}
	const baseLen = 5

	for i := 0; i < 20; i++ { 
		l := baseLen * 1<<i 
		in := make([]string, l) 
		in[0] = commands[i % 2]
		for k := 1; k < l; k++ { 
		 in[k] = "a" 
		} 
		b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) { 
		  parse(strings.Join(in, "")) 
		}) 
	 } 
}
