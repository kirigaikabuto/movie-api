package main

import (
	"fmt"
	"github.com/djumanoff/amqp"
	movie_store "github.com/kirigaikabuto/movie-store"
	"log"
)

var cfg = amqp.Config{
	Host:        "localhost",
	VirtualHost: "",
	User:        "",
	Password:    "",
	Port:        5672,
	LogLevel:    5,
}

var srvCfg = amqp.ServerConfig{
	ResponseX: "response",
	RequestX:  "request",
}

func main() {
	sess := amqp.NewSession(cfg)

	if err := sess.Connect(); err != nil {
		fmt.Println(err)
		return
	}
	defer sess.Close()

	srv, err := sess.Server(srvCfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	postgreConf := movie_store.Config{
		Host:             "localhost",
		User:             "kirito",
		Password:         "passanya",
		Port:             5432,
		Database:         "movie_store",
		ConnectionString: "",
		Params:           "sslmode=disable",
	}
	movieStore, err := movie_store.NewPostgreStore(postgreConf)
	if err != nil {
		log.Fatal(err)
	}
	movieService := movie_store.NewMovieService(movieStore)
	moviesAmqpEndpoints := movie_store.NewAMQPEndpointFactory(movieService)
	srv.Endpoint("movie.get", moviesAmqpEndpoints.GetMovieByIdAMQPEndpoint())
	srv.Endpoint("movie.create", moviesAmqpEndpoints.CreateMovieAMQPEndpoint())
	srv.Endpoint("movie.list", moviesAmqpEndpoints.ListMoviesAMQPEndpoint())
	srv.Endpoint("movie.update", moviesAmqpEndpoints.UpdateProductAMQPEndpoint())
	srv.Endpoint("movie.delete", moviesAmqpEndpoints.DeleteMovieAMQPEndpoint())
	srv.Endpoint("movie.getByName", moviesAmqpEndpoints.GetMovieByNameAMQPEndpoint())
	fmt.Println("Start server")
	if err := srv.Start(); err != nil {
		fmt.Println(err)
		return
	}
}
