package main
import (
    "bufio"
    "net"
    "fmt"
    "io"
    "os"
    "strings"
    "time"
)

func dnsG(host string) {
	ns, err := net.LookupHost(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err[%s]: %s\n", host, err.Error())
		return
	}

	for _, n := range ns {
		fmt.Fprintf(os.Stdout, "[%s]--[%s]\n", host,n) 
	}
}
func main() {
    //inputFile, inputError := os.Open("malice_domains.txt")
    inputFile, inputError := os.Open("china_domains.txt")
    if inputError != nil {
        fmt.Printf("An error occurred on opening the inputfile\n" +
            "Does the file exist?\n" +
            "Have you got acces to it?\n")
        return // exit the function on error
    }
    defer inputFile.Close()

    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerError := inputReader.ReadString('\n')
    	fmt.Printf("The input was: %s", inputString)
        result := strings.Replace(string(inputString),"\n","",1)
        go dnsG(result)
        if readerError == io.EOF {
            time.Sleep(60 * 1e9) // sleep for 60 seconds
            return
        }      
    }
}
