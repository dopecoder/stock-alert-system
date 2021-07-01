package core

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type MockRepo struct {
}

func (mr *MockRepo) CreateTrigger(trigger *Trigger) error {
	return nil
}

func (mr *MockRepo) UpdateTrigger(trigger *Trigger) error {
	return nil
}

func (mr *MockRepo) GetTriggers(GetTriggerOptions) ([]Trigger, error) {
	return []Trigger{}, nil
}

func (mr *MockRepo) OnTrigger(trigger *Trigger) error {
	endPoint := os.Getenv("UPDATES_POST_ENDPOINT")
	status, err := trigger.GetStatusString()
	if err != nil {
		return err
	}
	message := map[string]interface{}{
		"id":     trigger.Id,
		"status": status,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])

	return nil
}

func (mr *MockRepo) DeleteTrigger(trigger *Trigger) error {
	return nil
}
