package api

import (
	"github.com/flashguru-git/node-monitor-server/controllers/handlers"
	"github.com/flashguru-git/node-monitor-server/models"
)

var Routes = models.RoutePrefix{
	"/",
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
			"/nodes/latest",
			handlers.GetLatest,
			true,
		},
		{
			"GetNode",
			"GET",
			"/nodes/{nodeId}",
			handlers.GetNodeById,
			true,
		},
		{
			"GetAllNodes",
			"GET",
			"/nodes",
			handlers.GetAllNodes,
			true,
		},
		{
			"CreateNodeMetric",
			"POST",
			"/nodes",
			handlers.CreateNodeMetric,
			true,
		},
	},
}
