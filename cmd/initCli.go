package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	obd "github.com/WelintonJunior/obd-diagnostic-service/infraestructure/bluetooth"
	infraestructure "github.com/WelintonJunior/obd-diagnostic-service/infraestructure/postgres"
	"github.com/WelintonJunior/obd-diagnostic-service/utils"
	"github.com/spf13/cobra"

	tea "github.com/charmbracelet/bubbletea"
)

var initCliCmd = &cobra.Command{
	Use:   "initCli",
	Short: "Inicia o CLI interativo para leitura OBD2",
	Long:  `Menu interativo para conexão Bluetooth com dispositivo OBD2 e ações como leitura ou visualização de DTCs.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.LoadEnvMem(); err != nil {
			log.Fatalf("Erro ao carregar variáveis de ambiente: %v", err)
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

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		go SetupCli(ctx, &wg)

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig

		log.Println("Encerrando CLI...")
		wg.Wait()
	},
}

type menuModel struct {
	options  []string
	cursor   int
	selected string
	devices  []string
}

func initialModel() menuModel {
	devices := obd.ScanBluetoothDevices()
	return menuModel{
		options: []string{
			"Iniciar leitura OBD2",
			"Ver códigos DTC",
			"Sair",
		},
		cursor:  0,
		devices: devices,
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.options[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m menuModel) View() string {
	s := "\nSelecione uma opção:\n\n"

	for i, choice := range m.options {
		cursor := "  "
		if m.cursor == i {
			cursor = "👉"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nDispositivos Bluetooth encontrados:\n"
	for _, d := range m.devices {
		s += fmt.Sprintf("🔹 %s\n", d)
	}

	s += "\nUse ↑ ↓ para navegar e Enter para selecionar.\n"

	return s
}

func SetupCli(ctx context.Context, wg *sync.WaitGroup) {
	p := tea.NewProgram(initialModel())
	model, err := p.Run()
	if err != nil {
		log.Fatalf("Erro ao executar menu CLI: %v", err)
	}

	m := model.(menuModel)

	log.Printf("Opção selecionada: %s", m.selected)

	switch m.selected {
	case "Iniciar leitura OBD2":
		startOBDReadingSession()
	case "Ver códigos DTC":
		log.Println("Funcionalidade de códigos DTC ainda não implementada.")
	case "Sair":
		log.Println("Encerrando programa...")
	}
}

// func runGrpcPing(ctx context.Context) {
// 	conn, err := grpc.NewClient(os.Getenv("GRPC_HOST")+os.Getenv("GRPC_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("Erro ao conectar: %v", err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewDiagnosticsClient(conn)

// 	resp, err := client.Ping(ctx, &pb.Empty{})
// 	if err != nil {
// 		log.Fatalf("Erro no Ping: %v", err)
// 	}

// 	log.Printf("Resposta do servidor: %s", resp.GetMessage())
// }

func init() {
	rootCmd.AddCommand(initCliCmd)
}

func startOBDReadingSession() {
	fmt.Println("Iniciando leitura OBD2...\n")

	// obd.InitELM327("/dev/rfcomm0") // ou "COM3" no Windows
	obd.InitELM327("COM3") // ou "COM3" no Windows

	rpm := utils.ParseRPM(obd.ReadAndParse("010C"))
	speed := utils.ParseSpeedKPH(obd.ReadAndParse("010D"))
	temp := utils.ParseCoolantTemp(obd.ReadAndParse("0105"))
	throttle := utils.ParseThrottlePosition(obd.ReadAndParse("0104"))
	voltage := utils.ParseBatteryVoltage("")

	fmt.Println("🔧 Leituras OBD2:")
	fmt.Printf("➡️  RPM: %d rpm\n", rpm)
	fmt.Printf("➡️  Velocidade: %d km/h\n", speed)
	fmt.Printf("➡️  Temp. Arrefecimento: %.1f °C\n", temp)
	fmt.Printf("➡️  Posição Acelerador: %.1f%%\n", throttle)
	fmt.Printf("➡️  Voltagem Bateria: %.1f V\n", voltage)
}
