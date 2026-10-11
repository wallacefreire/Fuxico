package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL é a URL base da API que será utilizada para fazer requisições
	APIURL = ""
	// Porta é a porta que o servidor web irá escutar
	Porta = 0
	// HashKey é utilizada para autenticar de cookies
	HashKey []byte
	// BlockKey é utilizada para criptografia de cookies
	BlockKey []byte
)

// Carregar inicializa as variáveis de configuração da aplicação
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
