package main

import (
  "os"
  "fmt"
  "net"
  "strings"
  "time"
  "io/ioutil"
)

const (
  CONN_TYPE = "tcp"
  CONN_HOST = "0.0.0.0"
  CONN_PORT = "5555"
  ECHO = "ECHO"
  TIME = "TIME"
  CLOSE = "CLOSE"
  UPLOAD = "UPLOAD"
  CRLF = "\r\n"
  SERVER_FOLDER = "server_data"
)

func main() {
  fmt.Printf("Server launched on %s:%s\n", CONN_HOST, CONN_PORT)
  listener, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)

  if err != nil {
    fmt.Println("Listen error")
    os.Exit(1)
  }

  defer listener.Close(); // what is that gigle

  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println("Accept error")
      os.Exit(1)
    }
    fmt.Println("New connection created.")
    handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  for {
    buf := make([]byte, 1024)
    reqLen, err := conn.Read(buf)

    if err != nil {
      fmt.Println("Connection closed.")
      return
    }
    /*  print without CR LF (13 - carriage return, 10 - new line)
        fmt.Println(buf[0:reqLen])
    */
    if reqLen >= 2 {
      //fmt.Println(string(buf))
      reqStr := string(buf[:reqLen - 2]) // to string and remove CRLF
      handleMessage(conn, reqStr)
    } else {
      fmt.Println("Cannot convert byte array to string")
    }
  }
}

func handleMessage(conn net.Conn, msg string) {
  if strings.HasPrefix(msg, ECHO) {
    echo(conn, msg)
  } else if strings.HasPrefix(msg, TIME) {
    showTime(conn, msg)
  } else if strings.HasPrefix(msg, CLOSE) {
    conn.Close()
  } else if strings.HasPrefix(msg, UPLOAD) {
    uploadFile(conn, msg)
  }
}

func uploadFile(conn net.Conn, msg string) {
  fileName := msg[len(UPLOAD) + 1:] + "yo"
  err := ioutil.WriteFile(SERVER_FOLDER + fileName, []byte("yo nigga yo"), 0644)
  if err != nil {
    fmt.Printf("Cannot create %s file", fileName)
  } else {
    conn.Write([]byte(fileName + CRLF))
    fmt.Println("Uploading ", fileName)
  }
}

func echo(conn net.Conn, msg string) {
  echoMsg := msg[len(ECHO) + 1:] // fix slice out of bound error here
  fmt.Printf("Received msg: %s (length: %d)%s", echoMsg, len(echoMsg), CRLF)
  conn.Write([]byte(echoMsg + CRLF))
}

func showTime(conn net.Conn, msg string) {
  time :=  time.Now().String()
  conn.Write([]byte(time + CRLF))
  fmt.Println("Current cerver time: ", time)
}
