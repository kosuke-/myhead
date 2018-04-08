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

	if (number <= 0 ) {
		number = 0
	}
	m := &Myhead{
		number :  number,
		filepath: filepath,
	}
	return m
}

func (m *Myhead) Run(){

	headList := m.headCut()
	headPrint(headList)

	return
}

func (m *Myhead)headCut() map[string]string{

	if(0 == m.number){
		return nil
	}

	retData := make(map[string]string)

	for _,v := range m.filepath {

		head,err := headScanner(v,m.number);if err != nil{
			break;
		}
		retData[v] = head
	}
	return retData
}

func headScanner(path string,number int) (string,error){

	fp, err := os.Open(path);if err != nil{
		return "",err
	}

	scanner := bufio.NewScanner(fp)
	count := 0
	buf := ""
	for scanner.Scan() {
		if (count >= number) {
			break;
		}
		buf += scanner.Text()
		buf += "\n";
		count++
	}
	buf = strings.TrimRight(buf, "\n")
	fp.Close()
	return buf,nil
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