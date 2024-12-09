package task

import (
	"encoding/json"
	"log"
)

type Task struct {
	Type   string                 `json:"type"`   // Тип задачи
	Params map[string]interface{} `json:"params"` // Параметры задачи
}

// ProcessTask обрабатывает задачу на основе её типа.
func ProcessTask(taskData []byte) error {
	var task Task
	if err := json.Unmarshal(taskData, &task); err != nil {
		return err
	}

	log.Printf("Processing task: %+v", task)

	switch task.Type {
	case "iscsi-connect":
		return handleISCSIConnect(task.Params)
	case "nvmeof-connect":
		return handleNVMeOFConnect(task.Params)
	case "fc-connect":
		return handleFibreChannelConnect(task.Params)

	default:
		log.Printf("Unknown task type: %s", task.Type)
		return nil
	}
}
