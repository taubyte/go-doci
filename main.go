package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	doci "github.com/taubyte/go-doci/lib"
)

func decode_and_display_data(r string, decodeShowTS bool) {
	data, ts, err := doci.Decode(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	if decodeShowTS {
		fmt.Println("Timestamp: ", ts)
	}

	fmt.Println(data)
}

type entry struct {
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
}

func process_dns_txt(fqdn string, decodeShowTS bool, decodeAsJSON bool) {
	txtrecords, _ := net.LookupTXT(fqdn)

	b := make([]interface{}, 256)
	i := -1

	for _, txt := range txtrecords {
		i++
		if i >= cap(b) {
			break
		}
		if !decodeAsJSON {
			decode_and_display_data(txt, decodeShowTS)
		} else {
			data, ts, err := doci.Decode(txt)
			if err != nil {
				continue
			}
			if decodeShowTS {
				b[i] = entry{
					Timestamp: ts,
					Data:      data,
				}
			} else {
				b[i] = data
			}
			//fmt.Println(b[i])
		}
	}

	if decodeAsJSON {
		bj, _ := json.Marshal(b[:i+1])
		fmt.Println(string(bj))
	}
}

func main() {

	encodeCmd := flag.NewFlagSet("encode", flag.ExitOnError)

	decodeCmd := flag.NewFlagSet("decode", flag.ExitOnError)
	decodeFromTXT := decodeCmd.Bool("dns", false, "dns")
	decodeAsJSON := decodeCmd.Bool("json", false, "json")
	decodeShowTS := decodeCmd.Bool("timestamp", false, "timestamp")

	if len(os.Args) < 2 {
		fmt.Println("expected subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encode":
		encodeCmd.Parse(os.Args[2:])
		if len(encodeCmd.Args()) < 1 {
			fmt.Println("Expecting string to encode!")
			os.Exit(1)
		}
		m, _ := doci.Encode(encodeCmd.Args()[0], time.Now().Unix())
		fmt.Println(m)
	case "decode":
		decodeCmd.Parse(os.Args[2:])
		if len(decodeCmd.Args()) < 1 {
			fmt.Println("Expecting string or fqdn!")
			os.Exit(1)
		}
		if *decodeFromTXT {
			process_dns_txt(decodeCmd.Args()[0], *decodeShowTS, *decodeAsJSON)
		} else {
			decode_and_display_data(decodeCmd.Args()[0], *decodeShowTS)
		}
	default:
		fmt.Println("expected subcommands")
		os.Exit(1)
	}
}
