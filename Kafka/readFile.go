package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	//absPath, _ := filepath.Abs("./Data/order.fulfillment.txt")
	//file ,err := os.Open(absPath)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//defer file.Close()
	//reader := bufio.NewReader(file)
	//
	//for {
	//	line ,err := reader.ReadString('\n')
	//	if err != nil {
	//		if err ==io.EOF {
	//			break
	//		} else {
	//			fmt.Print(err)
	//		}
	//	}
	//	row := strings.Split(line, ":")
	//	//key := row[0]
	//	value :=  row[1]
	//	fmt.Printf(" kafka  value %v ", value )
	//}
	var lines = readFromFile1()
			for _, line:= range lines {
				a, _ := regexp.Compile(`\d*?:`)
				//row := a.Split(line,2)
				//strings.Split(line ,":")
				key :=(strings.Split(line ,":"))[0]
				value := a.Split(line,2)[1]
				//fmt.Printf("row %s ,%v" , key)
				fmt.Printf("key %v \n, value %v" ,key,value)

			}
	//fmt.Println(lines)
}


func readFromFile1() []string {
	absPath, _ := filepath.Abs("./Data/test.txt")
	file ,err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	//kafka:= make([]string,600)
	t := []string{}
	for {
		line ,err := reader.ReadString('\n')
		t = append(t,line)
		if err != nil {
			if err ==io.EOF {
				break
			} else {
				fmt.Print(err)
			}
		}
		row := strings.Split(line, ":")
		//fmt.Println(line)

		value :=  row[1]
		fmt.Printf(" kafka  value %v ", value )

	}

	return t
}