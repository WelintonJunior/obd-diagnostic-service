package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/WelintonJunior/obd-diagnostic-service/internal/grpcservice"
	pb "github.com/WelintonJunior/obd-diagnostic-service/proto"
	"github.com/WelintonJunior/obd-diagnostic-service/utils"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var grpcPort int

var initGrpcServerCmd = &cobra.Command{
	Use:   "initGrpcServer",
	Short: "Inicializa o serviço web de Grpc",
	Long:  "Comando para realizar a execução da aplicação",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.LoadEnvMem()
		if err != nil {
			log.Fatal("No .env file found")
		}

		portStr := os.Getenv("GRPC_PORT")
		grpcPort, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Variavel de ambiente GRPC_PORT inválido: %v", err)
		}

		go func() {
			StartGrpcServer(grpcPort)
		}()

		cancelChan := make(chan os.Signal, 1)
		signal.Notify(cancelChan, os.Interrupt, syscall.SIGTERM)

		<-cancelChan
		fmt.Println("Gracefully shutdown")
	},
}

func init() {
	rootCmd.AddCommand(initGrpcServerCmd)
	initGrpcServerCmd.Flags().IntVarP(&grpcPort, "port", "p", grpcPort, "Http server execution port")
}

func StartGrpcServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Erro ao escutar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDiagnosticsServer(s, &grpcservice.DiagnosticService{})

	log.Printf("gRPC server iniciado na porta %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Erro ao iniciar gRPC: %v", err)
	}
}
