package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Command struct {
	Fields []string
	Result chan string
}

func redisServer(commands chan Command) {
	var data = make(map[string]string)
	for cmd := range commands {

		if len(cmd.Fields) < 2 {
			cmd.Result <- "Expect at least 2 arguements"
			continue
		}

		fmt.Println("PROCESSING COMMAND", cmd)

		switch cmd.Fields[0] {

		case "SET":
			if len(cmd.Fields) != 3 {
				cmd.Result <- "EXPECT VALUE"
				continue
			}
			key := cmd.Fields[1]
			value := cmd.Fields[2]
			data[key] = value
			cmd.Result <- "VALUE SET"

		case "GET":
			key := cmd.Fields[1]
			value := data[key]
			cmd.Result <- value

		case "DEL":
			key := cmd.Fields[1]
			delete(data, key)
			cmd.Result <- "KEY VALUE DELETED"

		default:
			cmd.Result <- "INVALID COMMAND " + cmd.Fields[0] + "\n"
		}
	}
}

func handle(commands chan Command, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		text := strings.Fields(line)

		result := make(chan string)
		commands <- Command{
			Fields: text,
			Result: result,
		}

		io.WriteString(conn, <-result+"\n")
	}

}

func main() {
	Listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer Listener.Close()

	commands := make(chan Command)
	go redisServer(commands)

	for {
		conn, err := Listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		handle(commands, conn)
	}
}
