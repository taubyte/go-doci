package main

import (
	"fmt"
	"time"

	doci "github.com/taubyte/go-doci/lib"
)

func main() {
	m, _ := doci.Encode("/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN", time.Now().Unix())
	fmt.Println("encoded:", m)

	fmt.Println(doci.Decode(m))

	fmt.Println(doci.DecodeDnsTXT("__elders.net.taubyte.com", true))
}
