package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_subscriptions "github.com/apimanagementclient/mcp-server/tools/subscriptions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_subscriptions.CreateSubscription_regenerateprimarykeyTool(cfg),
		tools_subscriptions.CreateSubscription_regeneratesecondarykeyTool(cfg),
		tools_subscriptions.CreateSubscription_listTool(cfg),
		tools_subscriptions.CreateSubscription_getTool(cfg),
		tools_subscriptions.CreateSubscription_getentitytagTool(cfg),
		tools_subscriptions.CreateSubscription_updateTool(cfg),
		tools_subscriptions.CreateSubscription_createorupdateTool(cfg),
		tools_subscriptions.CreateSubscription_deleteTool(cfg),
	}
}
