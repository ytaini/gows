package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int) //记录相应字符出现的次数

	var utflen [utf8.UTFMax + 1]int //记录utf8字符的编码字节数的次数

	invalid := 0 //记录非法字符

	in := bufio.NewReader(os.Stdin)

	for {
		r, s, e := in.ReadRune()
		if e == io.EOF {
			break
		}

		if e != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", e)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && s == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[s]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
