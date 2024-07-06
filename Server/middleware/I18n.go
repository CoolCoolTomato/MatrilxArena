package middleware

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/bundle"
	"github.com/CoolCoolTomato/MatrilxArena/Server/config"
	"github.com/gin-gonic/gin"
	"sync"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type LocalizerCache struct {
	cache map[string]*CachedLocalizer
	mutex sync.RWMutex
}

type CachedLocalizer struct {
	localizer *i18n.Localizer
	lastUsed  time.Time
}

var (
	localizerCache *LocalizerCache
	once           sync.Once
)

const (
	cleanupInterval = 10 * time.Minute
	maxIdleTime     = 30 * time.Minute
)

func GetLocalizerCache() *LocalizerCache {
	once.Do(func() {
		localizerCache = &LocalizerCache{
			cache: make(map[string]*CachedLocalizer),
		}
		go localizerCache.startCleanupTask()
	})
	return localizerCache
}

func (lc *LocalizerCache) GetLocalizer(lang string) *i18n.Localizer {
	lc.mutex.RLock()
	cached, ok := lc.cache[lang]
	lc.mutex.RUnlock()

	if ok {
		cached.lastUsed = time.Now()
		return cached.localizer
	}

	localizer := i18n.NewLocalizer(bundle.GetBundle(), lang)

	lc.mutex.Lock()
	lc.cache[lang] = &CachedLocalizer{
		localizer: localizer,
		lastUsed:  time.Now(),
	}
	lc.mutex.Unlock()

	return localizer
}

func (lc *LocalizerCache) startCleanupTask() {
	ticker := time.NewTicker(cleanupInterval)
	for range ticker.C {
		lc.cleanup()
	}
}

func (lc *LocalizerCache) cleanup() {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	for lang, cached := range lc.cache {
		if time.Since(cached.lastUsed) > maxIdleTime {
			delete(lc.cache, lang)
		}
	}
}

func I18n() gin.HandlerFunc {
	localizerCache := GetLocalizerCache()

	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = config.GetConfig().GetString("Lang")
		}

		localizer := localizerCache.GetLocalizer(lang)
		c.Set("localizer", localizer)
		c.Next()
	}
}
