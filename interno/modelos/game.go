package modelos

type Game struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Developer   string `json:"developer"`
	Genre       string `json:"genre"`
	ReleaseYear int    `json:"release_year"`
	Difficulty  int    `json:"difficulty"`
	Platform    string `json:"platforms"`
	Boss_count  int    `json:"boss_count"`
}
