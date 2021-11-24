package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fptr := flag.String("fpath", "sonnets.txt", "file to read from")
	flag.Parse()

	// data, err := ioutil.ReadFile(*fptr)
	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }
	// fmt.Println("Contents", string(data))

	// readchunks(*fptr)
	readlines(*fptr)

}

func readchunks(path string) {
	buff, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		if err = buff.Close(); err != nil {
			log.Fatalln(err)
		}

	}()

	r := bufio.NewReader(buff)
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}

}

func readlines(path string) {

	fileOpen, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer fileOpen.Close()
	snl := bufio.NewScanner(fileOpen)
	//ignore first line
	if snl.Scan() {
		snl.Text()
	}

	for snl.Scan() {
		fmt.Println(snl.Text())
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
}
