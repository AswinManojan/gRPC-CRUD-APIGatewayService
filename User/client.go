package user

import (
	"log"

	pb "github.com/grpc/gobus/apigw/User/PB"
	"github.com/grpc/gobus/apigw/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClientDial(cfg *config.Config) (pb.UserServicesClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCUSERPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing client on port 8082")
		return nil, err
	}
	log.Printf("Successfully connected to client on port %s", cfg.GRPCUSERPORT)
	return pb.NewUserServicesClient(grpc), nil
}
