package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/WelintonJunior/obd-diagnostic-service/cmd"
	"github.com/WelintonJunior/obd-diagnostic-service/utils"
	"github.com/spf13/cobra"
)

var apiPort int

var initApiServerCmd = &cobra.Command{
	Use:   "initApiServer",
	Short: "Inicializa o serviço web de API",
	Long:  "Comando para realizar a execução da aplicação",
	Run: func(cmd *cobra.Command, args []string) {
		// app := SetupApp(context.Background())

		// portStr := os.Getenv("API_PORT")
		// apiPort, err := strconv.Atoi(portStr)
		// if err != nil {
		// 	log.Fatalf("Variavel de ambiente API_PORT inválido: %v", err)
		// }

		go func() {
			// runPort := fmt.Sprintf(":%d", apiPort)
			// if err := app.Listen(runPort); err != nil {
			// 	log.Panic(err)
			// }
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

func SetupApp(ctx context.Context) {
	err := utils.LoadEnvMem()
	if err != nil {
		log.Fatal("No .env file found")
	}

	return
}
