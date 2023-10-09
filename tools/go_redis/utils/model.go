package utils


type Player struct {
	ServerID   int    `json:"server_id"`
	PlayerID   int    `json:"player_id"`
	PlayerName string `json:"player_name"`
	RankType   string `json:"rank_type"`
	Power      int    `json:"power"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}