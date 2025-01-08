package apiproject

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendJson(w http.ResponseWriter, resp apiRespose, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to mashal json data", "error", err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}
