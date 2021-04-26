package api

import (
	"github.com/flashguru-git/node-monitor-server/controllers/handlers"
	"github.com/flashguru-git/node-monitor-server/models"
)

var Routes = models.RoutePrefix{
	"/api",
	[]models.Route{
		{
			"HealthCheck",
			"GET",
			"/healthCheck",
			handlers.HealthCheck,
			false,
		},
		{
			"GetLatest",
			"GET",
			"/latest",
			handlers.GetLatest,
			true,
		},
		{
			"GetNode",
			"GET",
			"/{nodeId}",
			handlers.GetNodeById,
			true,
		},
		{
			"PostNodeInfo",
			"POST",
			"/nodes",
			handlers.PostNodeInfo,
			true,
		},
	},
}
