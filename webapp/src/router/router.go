package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar cria as rotas da aplicação
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
