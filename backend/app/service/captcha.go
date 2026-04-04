package service

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// CaptchaStore 验证码存储（内存级）
type CaptchaStore struct {
	mu      sync.RWMutex
	codes   map[string]*captchaEntry
	expiry  time.Duration
	codeLen int
}

type captchaEntry struct {
	code      string
	expiresAt time.Time
}

var defaultCaptchaStore = NewCaptchaStore(5*time.Minute, 4)

// NewCaptchaStore 创建验证码存储
func NewCaptchaStore(expiry time.Duration, codeLen int) *CaptchaStore {
	cs := &CaptchaStore{
		codes:   make(map[string]*captchaEntry),
		expiry:  expiry,
		codeLen: codeLen,
	}
	// 启动定时清理过期验证码
	go cs.cleanup()
	return cs
}

// GetCaptchaStore 获取全局验证码存储
func GetCaptchaStore() *CaptchaStore {
	return defaultCaptchaStore
}

// Generate 生成验证码
// captchaId 为验证码标识（如请求ID或客户端IP）
func (cs *CaptchaStore) Generate(captchaId string) string {
	code := cs.generateCode()
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.codes[captchaId] = &captchaEntry{
		code:      code,
		expiresAt: time.Now().Add(cs.expiry),
	}
	return code
}

// Verify 校验验证码
func (cs *CaptchaStore) Verify(captchaId string, code string) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	entry, ok := cs.codes[captchaId]
	if !ok {
		return false
	}
	// 验证码使用后立即删除（一次性）
	delete(cs.codes, captchaId)
	if time.Now().After(entry.expiresAt) {
		return false
	}
	return entry.code == code
}

// generateCode 生成随机数字验证码
func (cs *CaptchaStore) generateCode() string {
	code := ""
	for i := 0; i < cs.codeLen; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}

// cleanup 定期清理过期验证码
func (cs *CaptchaStore) cleanup() {
	ticker := time.NewTicker(cs.expiry)
	defer ticker.Stop()
	for range ticker.C {
		cs.mu.Lock()
		now := time.Now()
		for id, entry := range cs.codes {
			if now.After(entry.expiresAt) {
				delete(cs.codes, id)
			}
		}
		cs.mu.Unlock()
	}
}
