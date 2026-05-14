package helpersWorkers

import "encoding/json"

func ParseMessage[T any](body []byte) (T, error) {
	var payload T

	err := json.Unmarshal(body, &payload)
	if err != nil {
		var zero T
		return zero, err
	}

	return payload, nil
}
