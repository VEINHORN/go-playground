package main

import(
  "fmt"
  "os"
  "os/exec"
  "strings"
  //"net"
)

const(
  SERVER_IP = "localhost"
  SERVER_PORT = "5555"
)

func main() {
  echoCommands()

  for {
    var command string
    fmt.Print("Enter command: ")
    fmt.Scanf("%s", &command)
    handleCommand(command)
  }
}

func handleCommand(command string) {
  if strings.HasPrefix(command, "CLEAR") {
    clearScreen()
  } else if strings.HasPrefix(command, "ECHO") {
    fmt.Println("yo")
  }
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
}
