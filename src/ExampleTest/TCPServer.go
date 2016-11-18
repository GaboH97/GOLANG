package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {

  fmt.Println("Launching server...")

  // listen on all interfaces
 
  ln, _ := net.Listen("tcp", ":8081")

  // accept connection on port
  conn, _ := ln.Accept()

  
  for {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n');
    // output message received
   // messageAsString := string(message)
    
   // if messageAsString 
    fmt.Print("Message Received:", string(message))
    fmt.Print("Text to send: ",string(message))
   // fmt.Print("%t",strings.EqualFold(strings.TrimRight(string(message), "\n"), "bye"));
    // sample process for string receive
    
    
    if(!strings.EqualFold(strings.TrimRight(string(message), "\n"), "bye")){
    	conn.Close();
    	break;
    }
    newmessage := strings.ToUpper(message)
    // send new string back to client
    conn.Write([]byte(newmessage + "\n"))
  }
}