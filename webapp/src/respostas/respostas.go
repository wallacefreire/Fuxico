package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroAPI é uma estrutura que representa uma mensagem de erro em formato JSON.
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON é uma função auxiliar para enviar respostas JSON para o cliente.
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// TratarStatusCodeDeErro é uma função auxiliar para lidar com códigos de status de erro HTTP.
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
