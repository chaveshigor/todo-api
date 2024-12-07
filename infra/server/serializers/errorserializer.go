package serializers

type SerializedError struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

func SerializeError(err error) SerializedError {
	return SerializedError{
		Error: struct {
			Message string `json:"message"`
		}{Message: err.Error()},
	}
}
