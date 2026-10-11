package modelos

// DadosAutenticacao representa os dados de autenticação retornados pelo servidor após o login
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
