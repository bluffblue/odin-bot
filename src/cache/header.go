package cache

import (
	"sync"
	"time"
	"math/rand"
	"github.com/bwmarrin/discordgo"
)

var (
	mutex sync.RWMutex
	userCache = make(map[string]interface{})
	guildCache = make(map[string]interface{})
	messageCache = make(map[string]interface{})
	
	cleanupMessages = []string{
		"Odin has initialized the cache, preparing for the next battle",
		"The AllFather has cleansed the sacred halls of Valhalla",
		"By Odin's wisdom, the cache has been purified",
		"The ravens of knowledge have cleared the ancient memories",
		"Behold, Odin's power flows through the system anew",
	}
)

func SetUserData(key string, value interface{}) {
	mutex.Lock()
	userCache[key] = value
	mutex.Unlock()
}

func GetUserData(key string) interface{} {
	mutex.RLock()
	defer mutex.RUnlock()
	return userCache[key]
}

func SetGuildData(key string, value interface{}) {
	mutex.Lock()
	guildCache[key] = value
	mutex.Unlock()
}

func GetGuildData(key string) interface{} {
	mutex.RLock()
	defer mutex.RUnlock()
	return guildCache[key]
}

func SetMessageData(key string, value interface{}) {
	mutex.Lock()
	messageCache[key] = value
	mutex.Unlock()
}

func GetMessageData(key string) interface{} {
	mutex.RLock()
	defer mutex.RUnlock()
	return messageCache[key]
}

func InitializeAutoCacheCleanup(session *discordgo.Session, guildID string, channels []string) {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			ClearCache()
			
			if len(channels) > 0 {
				randomChannel := channels[rand.Intn(len(channels))]
				randomMessage := cleanupMessages[rand.Intn(len(cleanupMessages))]
				
				session.ChannelMessageSend(randomChannel, randomMessage)
			}
		}
	}()
}

func ClearCache() {
	mutex.Lock()
	userCache = make(map[string]interface{})
	guildCache = make(map[string]interface{})
	messageCache = make(map[string]interface{})
	mutex.Unlock()
}
