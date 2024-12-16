package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/bhushan217/go-rnd/db"
	"log"
	"fmt"
	"encoding/json"
)

func canVote(name string, age int) string{

	// const username, age = "Bhushan", 37
	isAdult := If(age > 17, "Vote", "not Vote")
	return fmt.Sprintf("%s can %s", name, isAdult)
}

func main() {

	const username, age = "Bhushan", 37
	fmt.Println(canVote(username, age))

	ctx := context.Background()
	connStr := "postgresql://modulith:modulith@localhost:5532/modulith?sslmode=disable"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	defer conn.Close(ctx)
	queries := db.New(conn)
	permission, err := queries.CreatePermission(ctx, "CREATE_USER")
	if err != nil {
		log.Println(err)
		return
	}
	perm, err := json.Marshal(permission)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(perm)

	// list all permissions
	permissions, err := queries.ListPermission(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	// print all permissions
	for _, p := range permissions {
		perm, err := json.Marshal(p)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(perm))
	}

}

// If - the ternary function 
func If[T any](cond bool, vtrue T, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}
