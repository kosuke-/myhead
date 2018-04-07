package main

import (
	"fmt"
	"os"
	"bufio"
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
	for _,v := range m.filepath {
		fp, err := os.Open(v)
		if err != nil {
			return;
		}

		scanner := bufio.NewScanner(fp)
		count := 0
		for scanner.Scan() {
			if (count >= m.number) {
				break;
			}
			fmt.Println(scanner.Text())
			count++
		}
		if err := scanner.Err(); err != nil {
			return;
		}
		fp.Close()
	}
	return
}