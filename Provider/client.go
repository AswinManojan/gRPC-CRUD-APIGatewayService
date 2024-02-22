package provider

import (
	"log"

	pb "github.com/grpc/gobus/apigw/Provider/PB"
	"github.com/grpc/gobus/apigw/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ProviderClientDial(cfg *config.Config) (pb.ProviderServicesClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCSERVICEPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing client on port 8082")
		return nil, err
	}
	log.Printf("Successfully connected to client on port %s", cfg.GRPCSERVICEPORT)
	return pb.NewProviderServicesClient(grpc), nil
}
