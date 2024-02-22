package admin

import (
	"log"

	adminpb "github.com/grpc/gobus/apigw/Admin/pb"
	"github.com/grpc/gobus/apigw/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AdminClientDial(cfg *config.Config) (adminpb.AdminServiceClient, error) {
	grpc, err := grpc.Dial(":"+cfg.GRPCADMINPORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc client: 8081 ")
		return nil, err
	}
	log.Printf("succefully connected to admin server at port :8081")
	return adminpb.NewAdminServiceClient(grpc), nil
}
