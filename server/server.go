package main

import "fmt"
import "net"

func main() {
  listener, error := net.Listen("tcp", ":5000")
  
  if error != nil {
	fmt.Println("Erro ao iniciar o servidor: ", error)
	return
  }
  defer listener.Close()

  fmt.Println("Servidor rodando...")

  for {
	connection, error := listener.Accept()
	if error != nil {
		fmt.Println("Erro ao aceitar a conexão: ", error)
		continue
	}
	go handleConnection(connection)
  }
}

func handleConnection(connection net.Conn) {
  defer connection.Close()

  buffer := make([]byte, 1024)
  n, error := connection.Read(buffer)
  if error != nil {
      fmt.Println("Erro ao ler dados:", error)
      return
  }
  fmt.Println("Mensagem recebida:", string(buffer[:n]))

  response := []byte("Mensagem recebida!")
  connection.Write(response)
}