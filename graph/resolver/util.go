package resolver

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

func encodeNodeId(model string, id int) (string, error) {
	encoded := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", model, id)))
	return encoded, nil
}
func parseNodeID(id string) (string, int, error) {
	decoded, err := base64.URLEncoding.DecodeString(id)
	if err != nil {
		return "", 0, err
	}
	parts := strings.Split(string(decoded), ":")
	modelID, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, err
	}
	return parts[0], modelID, nil
}
