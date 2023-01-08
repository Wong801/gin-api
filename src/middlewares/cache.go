package middleware

import (
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func replyCache() cache.BeforeReplyWithCacheCallback {
	return func(c *gin.Context, cached *cache.ResponseCache) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func (m middleware) Cache() gin.HandlerFunc {
	return cache.CacheByRequestURI(
		persist.NewMemoryStore(1*time.Minute),
		30*time.Second,
		cache.WithBeforeReplyWithCache(replyCache()),
	)
}
