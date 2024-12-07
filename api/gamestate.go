package api

import (
  "encoding/json"
  "net/http"
)

type GameState struct {
  LobyID int
  CurrentStage int
  HostID int
}

func GetGameStateHandle(w http.ResponseWriter, r *http.Request) {
  testGamestate := &GameState{1, 0, 12345}

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(testGamestate)
}
