package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func help() {
	he := `
	IP         : ip of the target
	-p         : range to scan (default: 1-1000)
	-v         : verbrose, shows ports that are closed as well as other info
	-t         : scan type (default: tcp)
	example: go-scan xxx.xxx.xxx.xxx -p 1-1000 -t tcp -v`
	fmt.Println("USAGE: go-scan [IP] [ARGS]")
	fmt.Println(he)

}

func main() {
	verb := false
	typ := "tcp"
	ran := "1-1000"

	fmt.Println("go-scan!")

	args := os.Args[1:]

	for i, v := range args {
		if v == "-v" {
			verb = true
		} else if v == "-t" {
			typ = args[i+1]
		} else if v == "-p" {
			ran = args[i+1]
		}
	}

	if len(args) <= 0 {
		help()
	} else if strings.Contains(args[0], ".") {
		fmt.Println("[*] scanning ports...")

		num_base, err := strconv.Atoi(strings.Split(ran, "-")[0])
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		num_max, err := strconv.Atoi(strings.Split(ran, "-")[1])
		if err != nil {
			fmt.Println("ERROR: ", err)
		}

		for i := num_base; i != num_max; i++ {
			data := scanPort(typ, args[0], i)
			if data {
				fmt.Println(i, " : open")
			} else {
				if verb {
					fmt.Println(i, " : closed")
				}
			}

		}
	} else {
		help()
	}

}
