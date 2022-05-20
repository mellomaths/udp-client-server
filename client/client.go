package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"os"
)

type Request struct {
	Tipo string `json:"tipo"`
	Val  string `json:"val"`
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Execute o Cliente com o comando 'go run client.go 127.0.0.1:9922 int 68'")
	}

	server := os.Args[1]
	tipo := os.Args[2]
	val := os.Args[3]

	log.Print("Conectando ao servidor: ", server)
	buf := make([]byte, 8192)
	conn, err := net.Dial("udp", server)
	if err != nil {
		log.Print("Erro ao se conectar com Servidor UDP")
		log.Fatal(err)
	}

	defer conn.Close()

	request := &Request{tipo, val}

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		log.Print("Erro ao construir o JSON.")
		log.Fatal(err)
	}

	log.Print(string(jsonRequest))

	_, err = conn.Write(jsonRequest)
	if err != nil {
		log.Print("Erro ao enviar JSON para o servidor")
		log.Fatal(err)
	}

	log.Print("Aguardando a resposta do Server")

	_, err = bufio.NewReader(conn).Read(buf)
	if err == nil {
		log.Print(string(buf))
	} else {
		log.Print("Erro ao ler resposta do servidor UDP")
		log.Fatal(err)
	}
}
