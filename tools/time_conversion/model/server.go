package model

type Server struct {
	ID              int64  `json:"id"`
	ClientVer       string `json:"client_ver"`
	Name            string `json:"name"`
	Alias           string `json:"alias"`
	GameAddr        string `json:"game_addr"`
	GamePort        int64  `json:"game_port"`
	LogDBConfig     string `json:"log_db_config"`
	ReportURL       string `json:"report_url"`
	Status          int64  `json:"status"`
	CreateTime      string `json:"create_time"`
	RequireVer      int64  `json:"require_ver"`
	Remark          string `json:"remark"`
	JsonData        string `json:"json_data"`
	Order           int64  `json:"order"`
	Commend         int64  `json:"commend"`
	IsIOS           int64  `json:"is_ios"`
	LastTime        string `json:"last_time"`
	TabID           int64  `json:"tabId"`
	GameData        string `json:"game_data"`
	BattleplanID    int64  `json:"battleplan_id"`
	RealNumber      int64  `json:"real_number"`
	ExpectNumber    int64  `json:"expect_number"`
	ExpectTime      int64  `json:"expect_time"`
	SwitchStatus    int64  `json:"switch_status"`
	ServerVer       string `json:"server_ver"`
	Bitflag         uint   `json:"bitflag"`
	SeasonGroupID   int64  `json:"season_group_id"`
	SeasonStartTime string `json:"season_start_time"`
	SeasonEndTime   string `json:"season_end_time"`
}
