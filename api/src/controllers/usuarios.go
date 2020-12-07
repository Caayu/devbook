package controllers

import "net/http"

// CriarUsuario insere um usuario no baco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando um usuario!"))
}

// BuscarUsuarios busca todos os usuarios salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios!"))
}

// BuscarUsuario um usuario salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario!"))
}

// AtualizarUsuario altera as informacoes de um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuario!"))
}

// DeletarUsuario exclua as informacoes de um usuario no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuario!"))
}
