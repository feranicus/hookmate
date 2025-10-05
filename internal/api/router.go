package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/feranicus/hookmate/internal/config"
    "github.com/feranicus/hookmate/internal/services"
    "go.uber.org/zap"
)

func webhookHandler(slackService *services.SlackService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var payload map[string]interface{}
        if err := c.ShouldBindJSON(&payload); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json payload"})
            return
        }
        go func() {
            if err := slackService.SendNotification(payload); err != nil {
                // Logging is handled inside the service
            }
        }()
        c.JSON(http.StatusAccepted, gin.H{"status": "webhook accepted"})
    }
}

func Run(cfg *config.Config, logger *zap.SugaredLogger) error {
    router := gin.Default()
    slackService := services.NewSlackService(cfg.Integrations.Slack.WebhookURL, logger)
    router.GET("/healthz", func(c * gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })
    router.POST("/webhook/slack", webhookHandler(slackService))
    return router.Run(":" + cfg.Server.Port)
}