package helpers

func Response(message string, data any) map[string]any {
	var response = map[string]any{
		"message": message,
	}

	if data != nil {
		response["data"] = data
	}

	return response
}