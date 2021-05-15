package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"startern/companies"
	"startern/vacancies"
	"startern/generated"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	generated.RegisterCompaniesServer(grpcServer, &companies.CompaniesService{})
	generated.RegisterVacanciesServer(grpcServer, &vacancies.VacanciesService{})

	grpcServer.Serve(lis)
}
