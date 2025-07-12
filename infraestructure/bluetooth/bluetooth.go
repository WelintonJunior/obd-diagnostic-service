package infraestructure

import (
	"log"
	"strings"
	"time"

	"github.com/tarm/serial"
	s "go.bug.st/serial"
)

var elmPort *serial.Port

func ScanBluetoothDevices() []string {
	ports, err := s.GetPortsList()
	if err != nil {
		log.Printf("Erro ao listar portas seriais: %v", err)
		return []string{}
	}
	if len(ports) == 0 {
		return []string{"Nenhuma porta serial encontrada"}
	}
	return ports
}

func connectToELM327(portName string) *serial.Port {
	config := &serial.Config{
		Name:        portName, // ex: "/dev/rfcomm0" ou "COM3"
		Baud:        9600,
		ReadTimeout: time.Second * 2,
	}
	s, err := serial.OpenPort(config)
	if err != nil {
		log.Fatalf("Erro ao abrir porta serial: %v", err)
	}
	return s
}

func InitELM327(port string) {
	elmPort = connectToELM327(port)

	initCmds := []string{"ATZ", "ATE0", "ATL0", "ATH0"}
	for _, cmd := range initCmds {
		SendCommand(cmd)
		time.Sleep(300 * time.Millisecond)
	}
}

func SendCommand(cmd string) string {
	fullCmd := cmd + "\r"
	_, err := elmPort.Write([]byte(fullCmd))
	if err != nil {
		log.Printf("Erro ao enviar comando %s: %v", cmd, err)
		return ""
	}

	time.Sleep(500 * time.Millisecond)

	buf := make([]byte, 128)
	n, err := elmPort.Read(buf)
	if err != nil {
		log.Printf("Erro ao ler resposta: %v", err)
		return ""
	}

	resp := string(buf[:n])
	return cleanResponse(resp)
}

func cleanResponse(resp string) string {
	lines := strings.Split(resp, "\r")
	var cleaned []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ">") || strings.HasPrefix(line, "AT") {
			continue
		}
		cleaned = append(cleaned, line)
	}
	return strings.Join(cleaned, " ")
}

func ReadAndParse(pid string) string {
	return SendCommand(pid)
}
