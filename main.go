package main

import (
	"context"
	// "database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/bhushan217/go-rnd/api/rest/invoice"
	api "github.com/bhushan217/go-rnd/api/rest/user"
	"github.com/bhushan217/go-rnd/middleware"

	db "github.com/bhushan217/go-rnd/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

const SERVER_PORT = 8080
const v1URI = "/sample/v1"
const invoiceURI = "/invoice"
const IdURI = "{id}"

func canVote(name string, age int) string {

	// const username, age = "Bhushan", 37
	isAdult := If(age > 17, "Vote", "not Vote")
	return fmt.Sprintf("%s can %s", name, isAdult)
}

func main() {
	const username, age = "Bhushan", 37
	fmt.Println(canVote(username, age))

	initEnv()
	DoDbOperation()
	initRestServer()

}

func handleV1(res http.ResponseWriter, req *http.Request) {
	log.Println("Request received")
	res.Write([]byte("Hello World"))
	log.Println("Response sent")
}

func loadInvoiceRoutesGlobal(router *http.ServeMux) {
	handler := &invoice.Handler{}
	// router.Handler{}
	justSlash := invoiceURI
	slashWithId := justSlash+"/"+IdURI
	const POST = "POST "
	const GET = "GET "
	const PUT = "PUT "
	const DELETE = "DELETE "
	// const PATCH = "PATCH "
	router.HandleFunc((POST+justSlash), handler.Create)
	router.HandleFunc((GET+slashWithId), handler.FindByID)
	router.HandleFunc((PUT+slashWithId), handler.UpdateByID)
	router.HandleFunc((DELETE+slashWithId), handler.DeleteByID)
	router.HandleFunc((GET+invoiceURI), handler.FindAll)
	// router.HandleFunc((PUT+slashWithId), middleware.IsAdmin(http.HandlerFunc(handler.UpdateByID)))
	// router.HandleFunc((DELETE+slashWithId), middleware.IsAdmin(http.HandlerFunc(handler.DeleteByID)))
	// router.HandleFunc((GET+invoiceURI), middleware.IsAdmin(http.HandlerFunc(handler.FindAll)))
}

// func loadInvoiceRoutesAdmin(router *http.ServeMux) {
// 	// handler := &invoice.Handler{}
// }

func initRestServer() {

	invoiceRouterGlobal := http.NewServeMux()
	loadInvoiceRoutesGlobal(invoiceRouterGlobal)
	// invoiceRouterAdmin := http.NewServeMux()
	// loadInvoiceRoutesAdmin(invoiceRouterAdmin)
	// invoiceRouter := http.NewServeMux()
	// invoiceRouter.Handle(invoiceURI+"/", http.StripPrefix(invoiceURI, invoiceRouterGlobal))


	// getting env variables DB.PORT
	// viper.Get() returns an empty interface{}
	// so we have to do the type assertion, to get the value
	DB_URL, ok := viper.Get("DB_SOURCE").(string)
	if !ok {
		log.Fatal("DB URL not found")
	}
	ctx:=context.Background()
	// connStr := "postgresql://root:secret@localhost:5643/simple_blog?sslmode=disable"
	conn, err := pgx.Connect(ctx, DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	defer conn.Close(ctx)
	queries := db.New(conn)

	userService := api.NewService(queries)
	v1 := http.NewServeMux()
	v1.Handle(v1URI+"/", http.StripPrefix(v1URI, invoiceRouterGlobal))
	userRouter := http.NewServeMux()
	userService.RegisterHandlers(userRouter)
	v1.Handle(v1URI+"/user/", http.StripPrefix(v1URI+"/user", userRouter))

	// v1.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	// 	t, err := route.GetPathTemplate()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println(t)
	// 	return nil
	// })

	congiguredRouter := middleware.Chain(
		middleware.Logging,
		middleware.ErrorHandler,
		middleware.AllowCors,
		// middleware.IsAuthenticated,
		// middleware.EnsureAdmin,
		middleware.LoadUser,
		middleware.CheckPermissions,
	)(v1)
	server := http.Server{
		Addr:     fmt.Sprintf(":%d", SERVER_PORT),
		Handler:  congiguredRouter,
		ErrorLog: log.Default(),
	}

	log.Printf(`Server started on port %d`, SERVER_PORT)
	log.Fatal(server.ListenAndServe().Error())
	log.Printf(`Server started on port %d`, SERVER_PORT)

}

// If - the ternary function
func If[T any](cond bool, vtrue T, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func initEnv() {
	// Set the file name of the configurations file
	viper.SetConfigName("app")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

}

func getQuerier(ctx context.Context) *db.Queries{

	// getting env variables DB.PORT
	// viper.Get() returns an empty interface{}
	// so we have to do the type assertion, to get the value
	DB_URL, ok := viper.Get("DB_SOURCE").(string)
	if !ok {
		log.Fatal("DB URL not found")
	}

	// connStr := "postgresql://root:secret@localhost:5643/simple_blog?sslmode=disable"
	conn, err := pgx.Connect(ctx, DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	defer conn.Close(ctx)
	queries := db.New(conn)
	return queries
}

func DoDbOperation() {

	// getting env variables DB.PORT
	// viper.Get() returns an empty interface{}
	// so we have to do the type assertion, to get the value
	DB_URL, ok := viper.Get("DB_SOURCE").(string)
	if !ok {
		log.Fatal("DB URL not found")
	}

	ctx := context.Background()
	// connStr := "postgresql://root:secret@localhost:5643/simple_blog?sslmode=disable"
	conn, err := pgx.Connect(ctx, DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	defer conn.Close(ctx)
	queries := db.New(conn)
	permission, err := queries.CreatePermission(ctx, fmt.Sprintf("CREATE_USER %s", strconv.Itoa(rand.Intn(10))))
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
