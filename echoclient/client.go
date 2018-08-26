package main

import (
	"bytes"
	"os"
	"bufio"
	"fmt"
	"net"
)







// StartClient ....
func StartClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081");

	if err != nil {
		//
	}

	var reader *bufio.Reader;
	reader = bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n');
		if err != nil {
			fmt.Println("Read user input err", err);
			continue;
		}
		if len(line) == 1 {
			continue;
		}
		for i, c := range line {
			fmt.Printf("1 Index %d %d\n", i, c);
		}
		line = substring3(&line, 2, 6);
		for i, c := range line {
			fmt.Printf("2 Index %d %d\n", i, c);
		}
		b := ([]byte)(line);
		n, err := conn.Write(b);
		if err != nil {
			//
		}
		fmt.Printf("Write %d bytes\n", n);
	}

}

func substring(source *string, start int, end int) string {

	fmt.Printf("Param %d, %d\n", start, end);

	var r []rune;
	r = ([]rune)(*source);
	for i, c := range r {
		fmt.Printf("3 Index %d %d %s\n", i, c, string(c));
	}
	l := len(r);
	fmt.Println("rune len:", l);
	return string(r[start:end]);
}

func substring4(source *string, start int, end int) string {

	fmt.Printf("Param %d, %d\n", start, end);
	
	s := *source;
	
	count := 0;

	pre := 0;
	
	var buf bytes.Buffer;

	for i, c := range s {
		count += (i-pre);
		fmt.Printf("3 Index %d %d %d %s\n", i, count, c, string(c));
		if count>=start && count<end {
			buf.WriteRune(c);
		}
		if count >= end {
			break;
		}
		pre = i;
	}
	return buf.String();
}

func substring3(source *string, start int, end int) string {
	s := *source;
	fmt.Printf("Param %d, %d\n", start, end);
	return string(s[start:end]);
}

func substring2(source *string, start int, end int) string {

	fmt.Printf("Param %d, %d\n", start, end);

	var r []rune;
	r = ([]rune)(*source);
	l := len(r);
	fmt.Println("rune len:", l);
	
	var buf bytes.Buffer;

	index := 0;

	for i, c := range r {
		if index >=start && index < end {
			fmt.Printf("Append %d %d %s\n", i, c, string(c));
			buf.WriteRune(c);
		} else {
			fmt.Printf("Skip %d %d %s\n", i, c, string(c));
		}
		index++;
	}
	return buf.String();
	//return string(r[start:end]);
}
