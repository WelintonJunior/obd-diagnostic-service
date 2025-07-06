package api

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/WelintonJunior/obd-diagnostic-service/cmd"
	"github.com/WelintonJunior/obd-diagnostic-service/internal/grpcservice"
	pb "github.com/WelintonJunior/obd-diagnostic-service/proto"
	"github.com/WelintonJunior/obd-diagnostic-service/utils"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var apiPort int

var initApiServerCmd = &cobra.Command{
	Use:   "initApiServer",
	Short: "Inicializa o serviço web de API",
	Long:  "Comando para realizar a execução da aplicação",
	Run: func(cmd *cobra.Command, args []string) {

		portStr := os.Getenv("API_PORT")
		apiPort, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Variavel de ambiente API_PORT inválido: %v", err)
		}

		go func() {
			// runPort := fmt.Sprintf(":%d", apiPort)
			// if err := app.Listen(runPort); err != nil {
			// 	log.Panic(err)
			// }
			StartGrpcServer(apiPort)
		}()

		cancelChan := make(chan os.Signal, 1)
		signal.Notify(cancelChan, os.Interrupt, syscall.SIGTERM)

		<-cancelChan
		fmt.Println("Gracefully shutdown")
		// _ = app.Shutdown()
	},
}

func init() {
	cmd.RootCmd.AddCommand(initApiServerCmd)
	initApiServerCmd.Flags().IntVarP(&apiPort, "port", "p", apiPort, "Http server execution port")
}

func StartGrpcServer(port int) {
	err := utils.LoadEnvMem()
	if err != nil {
		log.Fatal("No .env file found")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Erro ao escutar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPrismaDiagnosticsServer(s, &grpcservice.DiagnosticService{})

	log.Printf("gRPC server iniciado na porta %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Erro ao iniciar gRPC: %v", err)
	}
}
