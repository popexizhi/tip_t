package main

import (
	"net"
	"fmt"
	"os"
)

func test(host string) {
	ns, err := net.LookupHost(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	for _, n := range ns {
		fmt.Fprintf(os.Stdout, "--%s\n", n) 
	}
}
func main(){
	test("qzonecn.com")
}
