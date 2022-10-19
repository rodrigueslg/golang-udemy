package servidor

import (
	"crud/bd"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	var usuario usuario

	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT id, nome, email FROM usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao executar query"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}

		usuarios = append(usuarios, usuario)
	}

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro para inteiro"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	linha, erro := db.Query("SELECT id, nome, email FROM usuarios WHERE id = ?", id)
	if erro != nil {
		w.Write([]byte("Erro ao executar query"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			http.Error(w, erro.Error(), http.StatusInternalServerError)
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro para inteiro"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, id); erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	id, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		w.Write([]byte("Erro ao converter parâmetro para inteiro"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	db, erro := bd.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(id); erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
