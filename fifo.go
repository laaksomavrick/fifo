package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting named pipe test...")

	// setup
	pipe := "pipe.log"

	os.Remove(pipe)
	err := syscall.Mkfifo(pipe, 0666)
	if err != nil {
		log.Fatal("Couldn't make fifo, ", err)
	}

	file, err := os.OpenFile(pipe, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal("Couldn't open pipe, ", err)
	}

	// write to named pipe
	go write(file)

	// read from named pipe
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal("Couldn't read line, ", err)
		}
		fmt.Printf(string(line))
	}

}

func write(file *os.File) {
	i := 0
	for {
		file.WriteString(fmt.Sprintf("wrote %d times\n", i))
		i++
		time.Sleep(time.Second)
	}
}
