package main

import (
	"log"
	"net"
	p "n11/Firdavs/dars3.6/genproto/transportService"
	pb "n11/Firdavs/dars3.6/genproto/weatherService"
	"n11/Firdavs/dars3.6/service"
	"n11/Firdavs/dars3.6/storage"
	"n11/Firdavs/dars3.6/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := storage.ConnnectDb()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	s := service.NewTransportService(postgres.ConnectBook(db))
	s1 := service.NewWeatherService()
	grpc := grpc.NewServer()
	p.RegisterTransportServiceServer(grpc, s)
	pb.RegisterWeatherServiceServer(grpc, s1)
	err = grpc.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}