package interfaces

import (
	"fmt"
	"net"

	"github.com/private-project-pp/pos-general-lib/infrastructure"
	"github.com/private-project-pp/pos-general-lib/logger"
	model "github.com/private-project-pp/pos-grpc-contract/model/product_service"
	"github.com/private-project-pp/product-rpc-service/handler"
	"github.com/private-project-pp/product-rpc-service/repository/postgre"
	"github.com/private-project-pp/product-rpc-service/shared/config"
	"github.com/private-project-pp/product-rpc-service/usecase/product_adm"
	"google.golang.org/grpc/reflection"
)

func Container() (err error) {
	fmt.Println("Start container")
	configs := config.SetupConfig()
	db, err := SetupDatabase(configs)
	if err != nil {
		return err
	}

	logging, err := logger.SetupLogger(configs.Service.LogFile)
	if err != nil {
		return err
	}

	mwConn, err := infrastructure.SetupMiddlewareClientConnection(configs.Internal.MiddlewareRpcService.Address)
	if err != nil {
		return err
	}
	defer mwConn.CloseConnection()

	server := infrastructure.GrpcInstanceServer(logging, mwConn)
	reflection.Register(server)

	// setup repository
	productsRepo := postgre.SetupProductsRepo(db)
	productUnitRepo := postgre.SetupProductUnitRepo(db)

	// setup usecase
	productAdministration := product_adm.SetupProductAdministration(productsRepo, productUnitRepo)

	//setup RPC handler
	rpcHandler := handler.SetupProductService(productAdministration)

	model.RegisterProductServiceServer(server, rpcHandler)

	listen, err := net.Listen("tcp", configs.Service.Port)
	if err != nil {
		return err
	}

	fmt.Println("RUNNING GRPC")
	fmt.Printf("Running on PORT [%s] \n", config.Service.Port)
	if err = server.Serve(listen); err != nil {
		return err
	}

	return nil
}
