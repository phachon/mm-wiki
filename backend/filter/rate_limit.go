package filter

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phachon/mm-wiki/app/entity"
	"github.com/phachon/mm-wiki/gopkg/errors"
)

// rateLimiter 简单的令牌桶速率限制器
type rateLimiter struct {
	mu       sync.Mutex
	buckets  map[string]*tokenBucket
	rate     int           // 每秒允许的请求数
	capacity int           // 桶容量
	expiry   time.Duration // 桶过期时间
}

type tokenBucket struct {
	tokens    float64
	lastTime  time.Time
	expiresAt time.Time
}

var defaultRateLimiter *rateLimiter

func init() {
	defaultRateLimiter = &rateLimiter{
		buckets:  make(map[string]*tokenBucket),
		rate:     100, // 每秒 100 个请求
		capacity: 200, // 桶容量 200
		expiry:   10 * time.Minute,
	}
	go defaultRateLimiter.cleanup()
}

// RateLimit API 速率限制中间件
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		if !defaultRateLimiter.allow(clientIP) {
			res := &entity.Response{
				Code:    int32(errors.ClientUnknownError),
				Message: "请求过于频繁，请稍后重试",
				Data:    nil,
			}
			c.JSON(http.StatusTooManyRequests, res)
			c.Abort()
			return
		}
		c.Next()
	}
}

func (rl *rateLimiter) allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	bucket, ok := rl.buckets[key]
	if !ok {
		rl.buckets[key] = &tokenBucket{
			tokens:    float64(rl.capacity) - 1,
			lastTime:  now,
			expiresAt: now.Add(rl.expiry),
		}
		return true
	}

	// 根据时间差补充令牌
	elapsed := now.Sub(bucket.lastTime).Seconds()
	bucket.tokens += elapsed * float64(rl.rate)
	if bucket.tokens > float64(rl.capacity) {
		bucket.tokens = float64(rl.capacity)
	}
	bucket.lastTime = now
	bucket.expiresAt = now.Add(rl.expiry)

	if bucket.tokens < 1 {
		return false
	}
	bucket.tokens--
	return true
}

func (rl *rateLimiter) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for key, bucket := range rl.buckets {
			if now.After(bucket.expiresAt) {
				delete(rl.buckets, key)
			}
		}
		rl.mu.Unlock()
	}
}
