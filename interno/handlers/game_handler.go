package handlers

import (
	"go-http/interno/modelos"
	"go-http/interno/storage"
	"go-http/interno/utils"

	"encoding/json"
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

	switch r.Method {

	case http.MethodGet:
		h.handleGetGames(w, r)

	case http.MethodPost:
		h.handleCreateGame(w, r)

	default:
		utils.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *GameHandler) handleGetGames(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	// 1️⃣ Prioridad: búsqueda por id
	if idParam := query.Get("id"); idParam != "" {

		id, err := strconv.Atoi(idParam)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
			return
		}

		for _, game := range h.store.Games {
			if game.ID == id {
				utils.WriteJSON(w, http.StatusOK, game)
				return
			}
		}

		utils.WriteError(w, http.StatusNotFound, "game not found")
		return
	}

	// 2️⃣ Filtros combinados
	filtered := h.store.Games

	if genre := query.Get("genre"); genre != "" {
		filtered = filterByGenre(filtered, genre)
	}

	if platform := query.Get("platform"); platform != "" {
		filtered = filterByPlatform(filtered, platform)
	}

	if difficultyParam := query.Get("difficulty"); difficultyParam != "" {
		difficulty, err := strconv.Atoi(difficultyParam)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, "invalid difficulty parameter")
			return
		}
		filtered = filterByDifficulty(filtered, difficulty)
	}

	utils.WriteJSON(w, http.StatusOK, filtered)
}

func filterByGenre(games []modelos.Game, genre string) []modelos.Game {
	var result []modelos.Game
	for _, game := range games {
		if game.Genre == genre {
			result = append(result, game)
		}
	}
	return result
}

func filterByPlatform(games []modelos.Game, platform string) []modelos.Game {
	var result []modelos.Game
	for _, game := range games {
		if game.Platform == platform {
			result = append(result, game)
		}
	}
	return result
}

func filterByDifficulty(games []modelos.Game, difficulty int) []modelos.Game {
	var result []modelos.Game
	for _, game := range games {
		if game.Difficulty == difficulty {
			result = append(result, game)
		}
	}
	return result
}

func (h *GameHandler) handleCreateGame(w http.ResponseWriter, r *http.Request) {

	var newGame modelos.Game

	// Decodificar JSON
	err := json.NewDecoder(r.Body).Decode(&newGame)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	// Validaciones básicas
	if newGame.Title == "" || newGame.Developer == "" {
		utils.WriteError(w, http.StatusBadRequest, "title and developer are required")
		return
	}

	if newGame.Difficulty < 1 || newGame.Difficulty > 10 {
		utils.WriteError(w, http.StatusBadRequest, "difficulty must be between 1 and 10")
		return
	}

	// Generar ID
	newGame.ID = h.generateNextID()

	// Agregar a memoria
	h.store.Games = append(h.store.Games, newGame)

	// Guardar en archivo
	if err := h.store.Save(); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to save game")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, newGame)
}

func (h *GameHandler) generateNextID() int {

	maxID := 0

	for _, game := range h.store.Games {
		if game.ID > maxID {
			maxID = game.ID
		}
	}

	return maxID + 1
}

func (h *GameHandler) HandleGameByID(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/api/games/")
	id, err := strconv.Atoi(path)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid game id")
		return
	}

	switch r.Method {

	case http.MethodGet:
		h.handleGetGameByID(w, id)

	case http.MethodPut:
		h.handleUpdateGame(w, r, id)

	case http.MethodPatch:
		h.handlePatchGame(w, r, id)

	case http.MethodDelete:
		h.handleDeleteGame(w, id)

	default:
		utils.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *GameHandler) handleGetGameByID(w http.ResponseWriter, id int) {

	for _, game := range h.store.Games {
		if game.ID == id {
			utils.WriteJSON(w, http.StatusOK, game)
			return
		}
	}

	utils.WriteError(w, http.StatusNotFound, "game not found")
}
func (h *GameHandler) handleUpdateGame(w http.ResponseWriter, r *http.Request, id int) {

	var updatedGame modelos.Game

	err := json.NewDecoder(r.Body).Decode(&updatedGame)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	// Validaciones básicas
	if updatedGame.Title == "" || updatedGame.Developer == "" {
		utils.WriteError(w, http.StatusBadRequest, "title and developer are required")
		return
	}

	if updatedGame.Difficulty < 1 || updatedGame.Difficulty > 10 {
		utils.WriteError(w, http.StatusBadRequest, "difficulty must be between 1 and 10")
		return
	}

	// Buscar y actualizar
	for i, game := range h.store.Games {
		if game.ID == id {

			updatedGame.ID = id // aseguramos que no cambien el ID

			h.store.Games[i] = updatedGame

			if err := h.store.Save(); err != nil {
				utils.WriteError(w, http.StatusInternalServerError, "failed to save data")
				return
			}

			utils.WriteJSON(w, http.StatusOK, updatedGame)
			return
		}
	}

	utils.WriteError(w, http.StatusNotFound, "game not found")
}

func (h *GameHandler) handlePatchGame(w http.ResponseWriter, r *http.Request, id int) {

	var updates map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	for i, game := range h.store.Games {
		if game.ID == id {

			// Aplicar cambios dinámicamente
			if title, ok := updates["title"].(string); ok {
				game.Title = title
			}

			if developer, ok := updates["developer"].(string); ok {
				game.Developer = developer
			}

			if genre, ok := updates["genre"].(string); ok {
				game.Genre = genre
			}

			if releaseYear, ok := updates["release_year"].(float64); ok {
				game.ReleaseYear = int(releaseYear)
			}

			if difficulty, ok := updates["difficulty"].(float64); ok {
				if difficulty < 1 || difficulty > 10 {
					utils.WriteError(w, http.StatusBadRequest, "difficulty must be between 1 and 10")
					return
				}
				game.Difficulty = int(difficulty)
			}

			if platform, ok := updates["platform"].(string); ok {
				game.Platform = platform
			}

			if bossCount, ok := updates["boss_count"].(float64); ok {
				game.Boss_count = int(bossCount)
			}

			h.store.Games[i] = game

			if err := h.store.Save(); err != nil {
				utils.WriteError(w, http.StatusInternalServerError, "failed to save data")
				return
			}

			utils.WriteJSON(w, http.StatusOK, game)
			return
		}
	}

	utils.WriteError(w, http.StatusNotFound, "game not found")
}

func (h *GameHandler) handleDeleteGame(w http.ResponseWriter, id int) {

	for i, game := range h.store.Games {
		if game.ID == id {

			// Eliminar del slice
			h.store.Games = append(h.store.Games[:i], h.store.Games[i+1:]...)

			if err := h.store.Save(); err != nil {
				utils.WriteError(w, http.StatusInternalServerError, "failed to save data")
				return
			}

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	utils.WriteError(w, http.StatusNotFound, "game not found")
}
