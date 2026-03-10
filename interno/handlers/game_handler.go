package handlers

import (
	"go-http/interno/storage"
	"go-http/interno/utils"
	"net/http"
	"strconv"
	"strings"
)

type GameHandler struct {
	store *storage.JSONStore
}

func NewGameHandler(store *storage.JSONStore) *GameHandler {
	return &GameHandler{store: store}
}

func (h *GameHandler) HandleGames(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	query := r.URL.Query()
	idParam := query.Get("id")

	// Si no viene id → devolver todos
	if idParam == "" {
		utils.WriteJSON(w, http.StatusOK, h.store.Games)
		return
	}

	// Convertir id a int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	// Buscar juego
	for _, game := range h.store.Games {
		if game.ID == id {
			utils.WriteJSON(w, http.StatusOK, game)
			return
		}
	}

	utils.WriteError(w, http.StatusNotFound, "game not found")
}

func (h *GameHandler) HandleGameByID(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Extraer el ID de la URL
	path := strings.TrimPrefix(r.URL.Path, "/api/games/")

	id, err := strconv.Atoi(path)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid game id")
		return
	}

	// Buscar juego
	for _, game := range h.store.Games {
		if game.ID == id {
			utils.WriteJSON(w, http.StatusOK, game)
			return
		}
	}

	utils.WriteError(w, http.StatusNotFound, "game not found")
}
