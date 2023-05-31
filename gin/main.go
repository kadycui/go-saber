package main

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

func main() {
	router := gin.Default()

	router.GET("/server", func(c *gin.Context) {
		serverID := c.Query("server_id")
		if serverID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server_id is required"})
			return
		}

		id, err := strconv.Atoi(serverID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server_id must be an integer"})
			return
		}

		server, ok := serverMap[id]
		if !ok {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, server)
	})

	router.POST("/server", func(c *gin.Context) {
		serverID := c.Query("server_id")
		if serverID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server_id is required"})
			return
		}

		id, err := strconv.Atoi(serverID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "server_id must be an integer"})
			return
		}

		var server Server

		err = c.BindJSON(&server)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid json format"})
			return
		}

		// check md5
		md5Hash := md5.Sum([]byte(server.Name + strconv.Itoa(server.MasterID) + strings.Join(server.ServerAliases, ",")))
		if hex.EncodeToString(md5Hash[:]) != c.GetHeader("X-Signature") {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid signature"})
			return
		}

		server.MasterID = id
		serverMap[id] = server

		c.JSON(http.StatusOK, server)
	})

	router.Run(":8080")
}
