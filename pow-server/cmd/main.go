package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/antsla/pow-server/helpers"
	"github.com/antsla/pow-server/internal"
	"github.com/antsla/pow-server/vars"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("TCP_BIND")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		m, err := bufio.NewReader(conn).ReadBytes(byte('\n'))
		if err != nil {
			fmt.Print(err)
			continue
		}

		r := vars.Request{}
		err = json.Unmarshal(m, &r)
		if err != nil {
			fmt.Print(err)
			continue
		}

		switch r.Type {
		case vars.ChooseType:
			fmt.Print("challenge request... \n")
			res := vars.Response{
				Type:       vars.ChallengeType,
				Data:       helpers.RandStringBytes(32),
				Complexity: vars.Complexity,
			}
			mes, err := json.Marshal(res)
			if err != nil {
				fmt.Print(err)
				continue
			}
			_, err = conn.Write(bytes.Join([][]byte{mes, []byte("\n")}, []byte{}))
			if err != nil {
				fmt.Print(err)
				continue
			}
		case vars.VerifyType:
			fmt.Print("trying verify...\n")
			pow := internal.NewProof(vars.Complexity, []byte(r.Data), r.Nonce)
			if pow.Validate() {
				i := rand.Intn(len(vars.Quotes))
				quote := vars.Quotes[i]

				res := vars.Response{
					Type:          vars.GrantType,
					WordsOfWisdom: quote,
				}
				mes, err := json.Marshal(res)
				if err != nil {
					fmt.Print(err)
					continue
				}
				_, err = conn.Write(bytes.Join([][]byte{mes, []byte("\n")}, []byte{}))
				if err != nil {
					fmt.Println(err)
					continue
				}

				fmt.Print("verify success\n")
				continue
			}

			fmt.Print("verify forbidden\n")
		default:
			fmt.Print("unhandled type\n")
		}
	}
}
