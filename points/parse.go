package points

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

func Parse(){
	file,err := os.Open("test")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	m := make(map[string]struct{})

	for{
		line,_,err :=reader.ReadLine()
		if err != nil && err == io.EOF{
			break
		}
		parse_line := strings.Fields(string(line))
		m[parse_line[1]+" "+parse_line[2]+" "+parse_line[3]+" "+parse_line[4]] = struct{}{}
	}
	file.Close()

	file,err = os.Open("points_final.txt")
	if err != nil {
		fmt.Println(err)
	}
	reader1 := bufio.NewReader(file)
	m1 := make(map[string]struct{})

	for{
		line,_,err :=reader1.ReadLine()
		if err != nil && err == io.EOF{
			break
		}
		parse_line := strings.Split(string(line),"-->")
		m1[strings.TrimSpace(parse_line[0])] = struct{}{}
	}
	file.Close()

	for key,_ := range m{
		if _,ok := m1[key];!ok{
			fmt.Println(key)
		}
	}
}
