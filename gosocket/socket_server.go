package main

import (
    "bufio"
    "fmt"
    "net"
)

func process(conn net.Conn) {
    // process finish, then close
    defer conn.Close()

    // the operator about send and receive
    for {
        reader := bufio.NewReader(conn)
        var buf [128]byte
        n, err := reader.Read(buf[:])
        if err != nil {
            fmt.Printf("read from conn failed, err:%v\n", err)
            break
        }

        recv := string(buf[:n])
        fmt.Printf("received data from client: %v\n", recv)

        // return the data from client to client
        _, err = conn.Write([]byte("server recive success: ok"))
        if err != nil {
            fmt.Printf("write from conn failed, err: %v\n", err)
            break
        }
    }
}

func main(){
    // create TCP server
    ip_port := "127.0.0.1:9090"
    listen, err := net.Listen("tcp", ip_port)
    if err != nil {
        fmt.Printf("listen failed, err: %v\n", err)
        return
    }
    fmt.Printf("Server start on: %s\n...\n", ip_port)

    for {
        // waiting for client to connect
        conn, err := listen.Accept()
        if err != nil {
            fmt.Printf("Accept failed, err:%v\n", err)
            continue
        }
        // start a new goroutine to face the gale
        go process(conn)
    }
}
