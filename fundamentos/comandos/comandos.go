package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!", "Go")
}

/*
comandos Go...

godoc -http=:6060 --> Abre a documentção da linguagem offline localmente (loccalhost:6060)
go env --> Oferece informações sobre a instalação do Go (GOPATH e GOROOT)
go doc cmd/vet --> Examina o código
go vet arquivo.go --> aponta os problemas do código do arquivo

*/
