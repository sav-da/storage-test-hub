package task

import (
	"fmt"
	"log"
	"os/exec"
)

func handleISCSIConnect(params map[string]interface{}) error {
	target, ok := params["target"].(string)
	if !ok || target == "" {
		return fmt.Errorf("invalid target parameter")
	}

	iqn, ok := params["iqn"].(string)
	if !ok || iqn == "" {
		return fmt.Errorf("invalid iqn parameter")
	}

	log.Printf("Connecting to iSCSI target: %s (IQN: %s)", target, iqn)

	// Выполняем iSCSI login через команду `iscsiadm`
	cmd := exec.Command("iscsiadm", "-m", "node", "-T", iqn, "-p", target, "--login")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to connect to iSCSI target: %s, output: %s", err, output)
	}

	log.Printf("Successfully connected to iSCSI target: %s", target)
	return nil
}
