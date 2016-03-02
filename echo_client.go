package main

import(
  "fmt"
  "os"
  "os/exec"
  "strings"
  "net"
  "bufio"
)

const(
  SERVER_IP = "0.0.0.0"
  SERVER_PORT = "5555"
  CRLF = "\r\n"
  BUF_SIZE = 1024
)

func main() {
  conn, err := net.Dial("tcp", SERVER_IP + ":" + SERVER_PORT)
  if err != nil {
    fmt.Println("Cannot connect to ", SERVER_IP + ":" + SERVER_PORT)
  } else {
    fmt.Println("Connected to ", SERVER_IP + ":" + SERVER_PORT)
  }

  echoCommands()

  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter command: ")
    command, _ := reader.ReadString('\n')

    // fmt.Printf("Entered command: %s (length: %s)\n", command, len(command))
    result := handleCommand(conn, command)
    if result == 1 {
      break
    }
  }
}

func handleCommand(conn net.Conn, command string) int {
  if strings.HasPrefix(command, "CLEAR") {
    clearScreen()
  } else if strings.HasPrefix(command, "ECHO") {
    sendMessage(conn, command)
  } else if strings.HasPrefix(command, "TIME") {
    sendMessage(conn, command)
  } else if strings.HasPrefix(command, "UPLOAD") {
    return uploadFile(conn, command)
  } else if strings.HasPrefix(command, "EXIT") {
    return 1
  }
  return 0
}

func sendMessage(conn net.Conn, command string) {
  msg := command[:len(command) - 1] + CRLF
  _, err := conn.Write([]byte(msg))
  if err != nil {
    fmt.Println("Cannot send data to server. Error: ", err.Error())
  }
  buf := make([]byte, BUF_SIZE)
  _, error := conn.Read(buf)
  if error != nil {
    fmt.Println("Connection closed.")
    return
  }
  fmt.Print(string(buf))
}

func uploadFile(conn net.Conn, command string) int {
  msg := command[:len(command) - 1] + CRLF
  _, err := conn.Write([]byte(msg))
  if err != nil {
    fmt.Println("Cannot send data to server. Error: ", err.Error())
  }
  buf := make([]byte, BUF_SIZE)
  _, error := conn.Read(buf)
  if error != nil {
    fmt.Println("Connection closed.")
  }
  fmt.Print(string(buf))
  for i := 0; i < 15; i++ {
    _, e := conn.Write(buf)
    if e != nil {
      fmt.Println("Cannot send chunk of file")
    }
  }

  _, er := conn.Write([]byte("eof_command"))
  if er != nil {
    fmt.Println("Cannot send eof command.")
  }
  return 0
}

func clearScreen() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func echoCommands() {
  fmt.Println("Availible commads: ")
  fmt.Println("1. ECHO <text>")
  fmt.Println("2. TIME")
  fmt.Println("3. CLOSE")
  fmt.Println("4. UPLOAD <filename>")
  fmt.Println("5. DOWNLOAD ")
  fmt.Println("6. CLEAR (clear screen)")
  fmt.Println("7. EXIT (close client)")
}
