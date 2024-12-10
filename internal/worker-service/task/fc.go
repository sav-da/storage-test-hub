package task

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// handleFibreChannelConnect настраивает подключение FC-таргета
func handleFibreChannelConnect(params map[string]interface{}) error {
	targetWWPN, ok := params["target_wwpn"].(string)
	if !ok || targetWWPN == "" {
		return fmt.Errorf("invalid target WWPN parameter")
	}

	lun, ok := params["lun"].(string)
	if !ok || lun == "" {
		return fmt.Errorf("invalid LUN parameter")
	}

	log.Printf("Connecting to FC target: WWPN=%s, LUN=%s", targetWWPN, lun)

	// Сканируем шины для подключения FC-таргетов
	cmd := exec.Command("rescan-scsi-bus.sh", "-w")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to rescan SCSI bus: %s, output: %s", err, output)
	}

	log.Println("SCSI bus rescan completed.")

	// Проверяем доступность устройства
	devicePath, err := findFCDevice(targetWWPN, lun)
	if err != nil {
		return fmt.Errorf("failed to find FC device: %s", err)
	}

	log.Printf("FC target connected successfully: %s", devicePath)
	return nil
}

// findFCDevice ищет устройство FC по WWPN и LUN
func findFCDevice(wwpn, lun string) (string, error) {
	cmd := exec.Command("lsscsi")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to list SCSI devices: %s", err)
	}

	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, wwpn) && strings.Contains(line, lun) {
			// Находим путь устройства
			fields := strings.Fields(line)
			if len(fields) >= 6 {
				return fields[5], nil
			}
		}
	}

	return "", fmt.Errorf("device not found for WWPN=%s, LUN=%s", wwpn, lun)
}
