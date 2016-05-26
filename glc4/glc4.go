package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"github.com/imel96/glc/survey"
	"github.com/streadway/amqp"
)

/*
echo -n "5#3,3" | nc localhost 8080
echo -n "100#97" | nc localhost 8080
echo -n "10#9,9,9,9,9" | nc localhost 8080
echo -n "7#1,2,3" | nc localhost 8080
echo -n "5#3,3,3" | nc localhost 8080
*/

func main() {
	l, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err);
		l, err := net.Listen("tcp", ":8080")

		if err != nil {
			fmt.Println(err);
			return
		}
	}
	defer l.Close()
	buf := make([]byte, 1500)

	for {
		conn, err := l.Accept()

		if err != nil {
			return
		}
		n, _ := conn.Read(buf)
		conn.Close()
		go handleRequest(string(buf[:n]))
	}
}

func handleRequest(req string) {
	N := 0
	tmp := ""
	var s []int

	_, err := fmt.Sscanf(req, "%d#%s", &N, &tmp)

	if err != nil {
		fmt.Println("Sscanf", err)
	}
	for _, a := range strings.Split(tmp, ",") {
		i, _ := strconv.Atoi(a)
		s = append(s, i)
	}
	fmt.Println("N: ", N, " data: ", tmp, "result: ", survey.CountBigShoppers(N, s))
}
