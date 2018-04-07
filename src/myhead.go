package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

type Myhead struct {
	number   int
	filepath []string
}

func NewMyhead(number int, filepath []string) *Myhead{
	m := &Myhead{
		number :  number,
		filepath: filepath,
	}
	return m
}

func (m *Myhead) Run()  {

	if(0 >= m.number){
		return
	}

	textMap := make(map[string]string)

	for _,v := range m.filepath {
		fp, err := os.Open(v)
		if err != nil {
			return;
		}

		scanner := bufio.NewScanner(fp)
		count := 0
		buf := ""
		for scanner.Scan() {
			if (count >= m.number) {
				break;
			}
			buf += scanner.Text()
			buf += "\n";
			count++
		}
		buf = strings.TrimRight(buf, "\n")
		textMap[v] = buf
		fp.Close()
	}

	headPrint(textMap)
	return
}

func headPrint(data map[string]string){

	size := len(data)
	count := 0
	for k,v := range data{
		count ++
		fmt.Printf("==> %s <==\n",k)
		fmt.Println(v)
		if size != count{
			fmt.Println()
		}
	}
	return
}