package main

import (
	"bufio"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	//For writing ORIGINAL.txt initially
	out_file, err := os.Create("Original.txt")
	check(err)
	defer out_file.Close()

	w := bufio.NewWriter(out_file)
	_, erro := w.WriteString("text text text text\ntext text text text\ntext text word text\ntext text text text\ntext text text text\n")
	check(erro)
	w.Flush()

	//Manipulation Part
	file, err := os.OpenFile("Original.txt", os.O_RDWR, 0777)
	check(err)
	defer file.Close()
	wr := bufio.NewWriter(file)

	stat, err := file.Stat()
	check(err)

	bs := make([]byte, stat.Size())
	_, err2 := file.Read(bs)
	check(err2)

	str := string(bs)

	file.Seek(0, 0)
	_, er := wr.WriteString("")
	check(er)
	wr.Flush()

	for i, v := range str {

		if v == 'w' {
			str = str[0:i-1] + " inserted " + str[i:]
		}
	}
	_, e := wr.WriteString(str)
	check(e)
	wr.Flush()
}
