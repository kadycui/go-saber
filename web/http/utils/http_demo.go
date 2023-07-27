package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Server struct {
	Name          string   `json:"server_name"`
	MasterID      int      `json:"master_id"`
	ServerAliases []string `json:"server_aliases"`
}

var serverMap = map[int]Server{
	1: {
		Name:          "Server 1",
		MasterID:      101,
		ServerAliases: []string{"server1.example.com", "server1.internal.local"},
	},
	2: {
		Name:          "Server 2",
		MasterID:      102,
		ServerAliases: []string{"server2.example.com", "server2.internal.local"},
	},
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	serverID := r.URL.Query().Get("server_id")
	if serverID == "" {
		http.Error(w, "server_id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(serverID)
	if err != nil {
		http.Error(w, "server_id must be an integer", http.StatusBadRequest)
		return
	}

	var server Server
	var ok bool

	switch r.Method {
	case http.MethodGet:
		server, ok = serverMap[id]
		if !ok {
			http.NotFound(w, r)
			return
		}
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&server)
		if err != nil {
			http.Error(w, "Json数据格式错误", http.StatusBadRequest)
			return
		}

		server.MasterID = id
		serverMap[id] = server
	default:
		http.Error(w, "不支持的请求方式", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(server)
}
