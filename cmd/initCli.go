package cmd

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	infraestructure "github.com/WelintonJunior/obd-diagnostic-service/infraestructure/postgres"
	"github.com/WelintonJunior/obd-diagnostic-service/utils"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var initCliCmd = &cobra.Command{
	Use:   "initCli",
	Short: "Comando para iniciar o consume de mensagens de energia",
	Long:  `Utilize o comando para iniciar o consume de mensagens de energia`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.LoadEnvMem(); err != nil {
			panic(err)
		}

		db, err := infraestructure.NewSqlDbConnection(infraestructure.GetSqlConfig())
		if err != nil {
			log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		}
		defer func() {
			dbSQL, _ := db.DB()
			dbSQL.Close()
		}()

		var wg sync.WaitGroup
		cron := cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
		go cron.Start()
		defer cron.Stop()

		// go ConsumeEnergyMessage(ctx, mqttConn, &wg)

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		log.Println("Encerrando Cli...")

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(initCliCmd)
}
