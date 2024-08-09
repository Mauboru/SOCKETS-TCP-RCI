package main

import "fmt"
import "net"
import "os"

func main(){
	connection, error := net.Dial("tcp", "localhost:5000")
	if error != nil {
		fmt.Println("Erro ao conectar ao servidor", error)
		os.Exit(1)
	}
	defer connection.Close()

	var message string
    fmt.Print("Digite sua mensagem: ")
    fmt.Scanln(&message)

    _, error = connection.Write([]byte(message))
    if error != nil {
        fmt.Println("Erro ao enviar mensagem:", error)
        return
    }

    buffer := make([]byte, 1024)
    n, err := connection.Read(buffer)
    if err != nil {
        fmt.Println("Erro ao receber resposta:", err)
        return
    }
    fmt.Println("Resposta do servidor:", string(buffer[:n]))
}