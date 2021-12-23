package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/antsla/pow-client/internal"
	"github.com/antsla/pow-client/vars"
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
	defer func() {
		errCls := conn.Close()
		if errCls != nil {
			fmt.Println(errCls)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	for {
		go handle(conn)
		time.Sleep(time.Second * 3)
	}
}

func handle(conn net.Conn) {
	res := vars.Response{
		Type: vars.ChooseType,
	}

	mes, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	_, err = fmt.Fprintf(conn, string(mes)+"\n")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	message, err := bufio.NewReader(conn).ReadBytes(byte('\n'))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	req := vars.Request{}
	err = json.Unmarshal(message, &req)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if req.Type == vars.ChallengeType {
		fmt.Println("challenge accepted")

		nonce, hash := internal.DoWork(req.Complexity, req.Data)
		fmt.Printf("work completed: [%d] %x \n", nonce, hash)

		res := vars.Response{
			Type:  vars.VerifyType,
			Nonce: nonce,
			Data:  req.Data,
		}
		mes, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		_, err = fmt.Fprintf(conn, string(mes)+"\n")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		message, err := bufio.NewReader(conn).ReadBytes(byte('\n'))
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		req := vars.Request{}
		err = json.Unmarshal(message, &req)
		if len(req.WordOfWisdom) != 0 {
			fmt.Printf("verify success: %s \n\n", req.WordOfWisdom)
		}
	}
}
