package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Gabriel-Macedogmc/grpc-golang/src/di"
	"github.com/Gabriel-Macedogmc/grpc-golang/src/infra/gorm"
	pb "github.com/Gabriel-Macedogmc/grpc-golang/src/proto"
	"github.com/Gabriel-Macedogmc/grpc-golang/src/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Erro fatal no arquivo de configuração: %w \n", err))
	}
}

func main() {
	psqlConn := viper.GetString("PSQL_CONN")

	fmt.Println(psqlConn)

	db := gorm.Connect(psqlConn)

	l, err := net.Listen("tcp", "localhost:5432")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	conn := di.ConfigCategoryDI(db)

	pb.RegisterCategoryServiceServer(server, service.NewService(conn))

	fmt.Println("Server starting")

	server.Serve(l)
}
