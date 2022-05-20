package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
)

type RequestInt struct {
	Tipo string `json:"tipo"`
	Val  int    `json:"val"`
}

type RequestChar struct {
	Tipo string `json:"tipo"`
	Val  byte   `json:"val"`
}

type RequestString struct {
	Tipo string `json:"tipo"`
	Val  string `json:"val"`
}

func main() {
	server := "127.0.0.1:9922"
	log.Print("Conectando ao servidor: ", server)
	buf := make([]byte, 8192)
	conn, err := net.Dial("udp", server)
	if err != nil {
		log.Print("Erro ao se conectar com Servidor UDP")
		log.Fatal(err)
	}

	defer conn.Close()

	// request := &RequestInt{"int", 68}
	// request := &RequestString{"string", "Request"}
	request := &RequestChar{"char", 'R'}

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
