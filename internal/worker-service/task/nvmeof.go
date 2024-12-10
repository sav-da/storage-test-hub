package task

import (
	"fmt"
	"log"
	"os/exec"
)

// handleNVMeOFConnect подключает NVMe-oF таргет
func handleNVMeOFConnect(params map[string]interface{}) error {
	target, ok := params["target"].(string)
	if !ok || target == "" {
		return fmt.Errorf("invalid target parameter")
	}

	subsystem, ok := params["subsystem"].(string)
	if !ok || subsystem == "" {
		return fmt.Errorf("invalid subsystem parameter")
	}

	log.Printf("Connecting to NVMe-oF target: %s (Subsystem: %s)", target, subsystem)

	// Выполняем подключение через `nvme connect`
	cmd := exec.Command("nvme", "connect", "-t", "tcp", "-a", target, "-s", "4420", "-n", subsystem)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to connect to NVMe-oF target: %s, output: %s", err, output)
	}

	log.Printf("Successfully connected to NVMe-oF target: %s (Subsystem: %s)", target, subsystem)
	return nil
}
