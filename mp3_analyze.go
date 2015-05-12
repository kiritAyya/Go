package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	m_file, err := os.OpenFile("SmoothCriminal.mp3", os.O_RDONLY, 0777)
	check(err)
	defer m_file.Close()

	stat, er := m_file.Stat()
	check(er)

	buff := make([]byte, stat.Size())
	_, e := m_file.Read(buff)
	check(e)

	n_buff := ""

	str := string(buff)

	for i, v := range str {
		if v == 'G' {
			n_buff = str[i+1:]
		}
	}

	n_buff_ := strings.Fields(n_buff)

	fmt.Println(n_buff_)
}
