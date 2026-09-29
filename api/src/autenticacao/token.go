package autenticacao

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// CriarToken cria um token JWT para o usuário com base no seu ID.
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = usuarioID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("Secret"))
}
