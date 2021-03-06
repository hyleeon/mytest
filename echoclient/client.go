package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"unicode/utf8"
)

// StartClient ....
func StartClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")

	if err != nil {
		//
	}

	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read user input err", err)
			continue
		}
		if len(line) == 1 {
			continue
		}

		fmt.Printf("Read str len %d\n", utf8.RuneCountInString(line)-1)

		// for i, c := range line {
		// 	fmt.Printf("1 Index %d %d\n", i, c)
		// }
		line, err = substring(&line, 0, len(line)-1, true)
		if err != nil {
			//
			fmt.Println("substring msg error", err)
			continue
		}
		// for i, c := range line {
		// 	fmt.Printf("2 Index %d %d\n", i, c)
		// }
		b := ([]byte)(line)
		n, err := conn.Write(b)
		if err != nil {
			//
		}
		fmt.Printf("Write %d bytes\n", n)
	}

}

func substring(source *string, start int, end int, usedStrLen bool) (string, error) {

	if source == nil {
		return "", errors.New("Source is nil")
	}

	if start < 0 || start > end || end < 0 {
		msg := fmt.Sprintf("Invalide index: start=%d, end=%d", start, end)
		return "", errors.New(msg)
	}

	if usedStrLen && start == 0 && len(*source) == end {
		return *source, nil
	}

	if usedStrLen {
		return substringWithLenCalIndex(source, start, end), nil
	}
	return substringWithSpecifyIndex(source, start, end), nil
}

func substringWithLenCalIndex(source *string, start int, end int) string {

	// fmt.Printf("Param %d, %d\n", start, end)

	s := *source

	count := 0

	pre := 0

	var buf bytes.Buffer

	for i, c := range s {
		count += (i - pre)
		// fmt.Printf("3 Index %d %d %d %s\n", i, count, c, string(c))
		if count >= start && count < end {
			buf.WriteRune(c)
		}
		if count >= end {
			break
		}
		pre = i
	}
	return buf.String()
}

func substringWithSpecifyIndex(source *string, start int, end int) string {
	s := *source
	var r []rune
	r = ([]rune)(s)
	return string(r[start:end])
}

// func substring0(source *string, start int, end int) string {

// 	fmt.Printf("Param %d, %d\n", start, end)

// 	var r []rune
// 	r = ([]rune)(*source)
// 	for i, c := range r {
// 		fmt.Printf("3 Index %d %d %s\n", i, c, string(c))
// 	}
// 	l := len(r)
// 	fmt.Println("rune len:", l)
// 	return string(r[start:end])
// }

// func substring3(source *string, start int, end int) string {
// 	s := *source
// 	fmt.Printf("Param %d, %d\n", start, end)
// 	return string(s[start:end])
// }

// func substring2(source *string, start int, end int) string {

// 	fmt.Printf("Param %d, %d\n", start, end)

// 	var r []rune
// 	r = ([]rune)(*source)
// 	l := len(r)
// 	fmt.Println("rune len:", l)

// 	var buf bytes.Buffer

// 	index := 0

// 	for i, c := range r {
// 		if index >= start && index < end {
// 			fmt.Printf("Append %d %d %s\n", i, c, string(c))
// 			buf.WriteRune(c)
// 		} else {
// 			fmt.Printf("Skip %d %d %s\n", i, c, string(c))
// 		}
// 		index++
// 	}
// 	return buf.String()
// 	//return string(r[start:end]);
// }
