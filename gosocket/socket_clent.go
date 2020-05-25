package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    // 1. create connect to server
    conn, err := net.Dial("tcp", "127.0.1:9090")
    if err != nil {
        fmt.Printf("conn server failed, err:v%\n", err)
        return
    }

    // send data to server
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Println("Please enter somthing:")
        text, _ := reader.ReadString('\n')
        text = strings.TrimSpace(text)
        if strings.ToUpper(text) == "Q" {
            fmt.Println("Quit success")
            return
        }
    _, err = conn.Write([]byte(text))
    if err != nil {
        fmt.Printf("send failed, err: %v\n", err)
        return
    }
    // recive the data from server
    var buf [1024]byte
    n, err := conn.Read(buf[:])
    if err != nil {
        fmt.Printf("read failed:v%\n", err)
        return
    }
    fmt.Printf("recive the data from server: %v\n", string(buf[:n]))
    }
}
