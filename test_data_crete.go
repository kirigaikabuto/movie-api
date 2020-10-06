package main

//import (
//	"encoding/csv"
//	"fmt"
//	movie_store "github.com/kirigaikabuto/movie-store"
//	"log"
//	"os"
//)
//
//func main(){
//	postgreConf := movie_store.Config{
//		Host:             "localhost",
//		User:             "kirito",
//		Password:         "passanya",
//		Port:             5432,
//		Database:         "movie_store",
//		ConnectionString: "",
//		Params:           "sslmode=disable",
//	}
//	movieStore, err := movie_store.NewPostgreStore(postgreConf)
//	if err != nil {
//		log.Fatal(err)
//	}
//	movieService := movie_store.NewMovieService(movieStore)
//
//	f, _ := os.Open("data.csv")
//
//	defer f.Close()
//
//	csvReader := csv.NewReader(f)
//	records, _ := csvReader.ReadAll()
//	for _, v := range records {
//		cmdEl := movie_store.CreateMovieCommand{
//			v[1],
//			v[2],
//			v[7],
//			v[4],
//			v[3],
//			v[6],
//			9,
//		}
//		el,err := movieService.CreateMovie(&cmdEl)
//		if err != nil{
//			log.Fatal(err)
//		}
//		fmt.Println(el)
//		fmt.Println(v[0],v[1],v[2],v[3],v[4],v[6],v[7])
//	}
//
//}
