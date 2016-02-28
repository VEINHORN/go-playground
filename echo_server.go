package main

import (
  "os"
  "fmt"
  "net"
)

const (
  CONN_TYPE = "tcp"
  CONN_HOST = "0.0.0.0"
  CONN_PORT = "5555"
)

func main() {
  fmt.Println("Launching server...")
  ln, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)

  if err != nil {
    fmt.Println("Listen error")
    os.Exit(1)
  }
  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Accept error")
      os.Exit(1)
    }
    handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  buf := make([]byte, 1024)
  reqLen, err := conn.Read(buf)

  if err != nil {
    fmt.Println("Error reading: ", err.Error())
  }
  fmt.Println(buf[0:reqLen]) // print without CR LF (carriage return, new line)
  if reqLen >= 2 {
    reqStr := string(buf[:reqLen - 2])
    fmt.Printf("Received msg: %s (length: %d)\n", reqStr, len(reqStr))
  } else {
    fmt.Println("Cannot convert byte array to string")
  }

  conn.Close()
}
