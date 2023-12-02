package http

import "net/http"

func (app *App) handleAddToQueue(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
