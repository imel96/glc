package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"glc/survey"
	"github.com/streadway/amqp"
)

/*
		fmt.Println("buyer: ", survey.CountBigShoppers(100, []int{97}));
		fmt.Println("buyer: ", survey.CountBigShoppers(10, []int{9, 9, 9, 9, 9}));
		fmt.Println("buyer: ", survey.CountBigShoppers(7, []int{1, 2, 3}));
		fmt.Println("buyer: ", survey.CountBigShoppers(5, []int{3, 3, 3}));
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

	for {
		conn, err := l.Accept()

		if err != nil {
			return
		}
		go handleRequest(conn)
	}
}

func handleRequest(c net.Conn) {
	buf := make([]byte, 1500)
	N := 0
	tmp := ""
	var s []int
	n, _ := c.Read(buf)

	_, err := fmt.Sscanf(string(buf[:n]), "%d#%s", &N, &tmp)
	c.Close()

	if err != nil {
		fmt.Println("Sscanf", err)
	}
	for _, a := range strings.Split(tmp, ",") {
		i, _ := strconv.Atoi(a)
		s = append(s, i)
	}
	fmt.Println("N: ", N, " data: ", tmp, "result: ", survey.CountBigShoppers(N, s))
}
