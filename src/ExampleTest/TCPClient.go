package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  for { 
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
    text, _ := reader.ReadString('\n')
    fmt.Println("Message to send",text)
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply
    message, err := bufio.NewReader(conn).ReadString('\n')
    bufio.NewReader(conn).re
    if(err != nil){
    	fmt.Println("Unresponsive server")
    }
    fmt.Print("Message from server: "+message)
  }
}