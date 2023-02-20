package resolver

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type ModelType string

const (
	ModelTypeUser ModelType = "User"
	ModelTypeTodo ModelType = "Todo"
)

func encodeNodeId(model ModelType, id int) (string, error) {
	encoded := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", model, id)))
	return encoded, nil
}

func parseNodeID(id string) (ModelType, int, error) {
	decoded, err := base64.URLEncoding.DecodeString(id)
	if err != nil {
		return "", 0, err
	}

	parts := strings.Split(string(decoded), ":")

	modelType := ModelType(parts[0])
	switch modelType {
	case ModelTypeUser, ModelTypeTodo: // do nothing
	default:
		return "", 0, fmt.Errorf("unknown model type: %s", modelType)
	}

	modelID, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, err
	}
	return modelType, modelID, nil
}
