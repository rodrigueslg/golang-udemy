package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	strConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", strConexao)
	if erro != nil {
		fmt.Println("Erro ao executar sql.Open()")
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		fmt.Println("Erro ao executar db.Ping()")
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("SELECT * FROM users")
	if erro != nil {
		fmt.Println("Erro ao executar db.Query()")
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
