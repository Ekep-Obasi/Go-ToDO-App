package utils

import (
	"encoding/json"

	"github.com/google/uuid"
)

func ToJSON(encoded interface{}) ([]byte, error) {
	jsonReturn, err := json.Marshal(encoded)

	if err != nil {
		return jsonReturn, err
	}
	return jsonReturn, nil
}

func ToSTRUCT(jsonData []byte, result interface{}) error {
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return err
	}
	return nil
}

func generateUniqueID() string {
	id := uuid.New()

	return id.String()
}