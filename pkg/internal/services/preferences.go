package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	localCache "git.solsynth.dev/hydrogen/passport/pkg/internal/cache"
	"git.solsynth.dev/hydrogen/passport/pkg/internal/database"
	"git.solsynth.dev/hydrogen/passport/pkg/internal/models"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/marshaler"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/samber/lo"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func GetNotificationPreferenceCacheKey(accountId uint) string {
	return fmt.Sprintf("notification-preference#%d", accountId)
}

func GetNotificationPreference(account models.Account) (models.PreferenceNotification, error) {
	var notification models.PreferenceNotification
	cacheManager := cache.New[any](localCache.S)
	marshal := marshaler.New(cacheManager)
	contx := context.Background()

	if val, err := marshal.Get(contx, GetNotificationPreferenceCacheKey(account.ID), new(models.PreferenceNotification)); err == nil {
		notification = val.(models.PreferenceNotification)
	} else {
		if err := database.C.Where("account_id = ?", account.ID).First(&notification).Error; err != nil {
			return notification, err
		}
		CacheNotificationPreference(notification)
	}

	return notification, nil
}

func CacheNotificationPreference(prefs models.PreferenceNotification) {
	cacheManager := cache.New[any](localCache.S)
	marshal := marshaler.New(cacheManager)
	contx := context.Background()

	marshal.Set(
		contx,
		GetNotificationPreferenceCacheKey(prefs.AccountID),
		prefs,
		store.WithExpiration(60*time.Minute),
		store.WithTags([]string{"notification-preference", fmt.Sprintf("user#%d", prefs.AccountID)}),
	)
}

func UpdateNotificationPreference(account models.Account, config map[string]bool) (models.PreferenceNotification, error) {
	var notification models.PreferenceNotification
	var err error
	if notification, err = GetNotificationPreference(account); err != nil {
		notification = models.PreferenceNotification{
			AccountID: account.ID,
			Config: datatypes.JSONMap(
				lo.MapValues(config, func(v bool, k string) any { return v }),
			),
		}
	} else {
		notification.Config = datatypes.JSONMap(
			lo.MapValues(config, func(v bool, k string) any { return v }),
		)
	}

	err = database.C.Save(&notification).Error
	if err == nil {
		CacheNotificationPreference(notification)
	}

	return notification, err
}

func CheckNotificationNotifiable(account models.Account, topic string) bool {
	var notification models.PreferenceNotification
	cacheManager := cache.New[any](localCache.S)
	marshal := marshaler.New(cacheManager)
	contx := context.Background()

	if val, err := marshal.Get(contx, GetNotificationPreferenceCacheKey(account.ID), new(models.PreferenceNotification)); err == nil {
		notification = val.(models.PreferenceNotification)
	} else {
		if err := database.C.Where("account_id = ?", account.ID).First(&notification).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return true
			}
			return false
		}
		CacheNotificationPreference(notification)
	}

	if val, ok := notification.Config[topic]; ok {
		if status, ok := val.(bool); ok {
			return status
		}
	}
	return true
}

func CheckNotificationNotifiableBatch(accounts []models.Account, topic string) []bool {
	cacheManager := cache.New[any](localCache.S)
	marshal := marshaler.New(cacheManager)
	contx := context.Background()

	var notifiable = make([]bool, len(accounts))
	queryNeededIdx := []uint{}
	notificationMap := make(map[uint]models.PreferenceNotification)

	// Check cache for each account
	for idx, account := range accounts {
		cacheKey := GetNotificationPreferenceCacheKey(account.ID)
		if val, err := marshal.Get(contx, cacheKey, new(models.PreferenceNotification)); err == nil {
			notification := val.(models.PreferenceNotification)
			notificationMap[account.ID] = notification
			// Determine if the account is notifiable based on the topic config
			if val, ok := notification.Config[topic]; ok {
				if status, ok := val.(bool); ok {
					notifiable[idx] = status
					continue
				}
			}
			notifiable[idx] = true
		} else {
			// Add to the list of accounts that need to be queried
			queryNeededIdx = append(queryNeededIdx, account.ID)
		}
	}

	// Query the database for missing account IDs
	if len(queryNeededIdx) > 0 {
		var dbNotifications []models.PreferenceNotification
		if err := database.C.Where("account_id IN ?", queryNeededIdx).Find(&dbNotifications).Error; err != nil {
			// Handle error by returning false for accounts without cached notifications
			return lo.Map(accounts, func(item models.Account, index int) bool {
				return false
			})
		}

		// Cache the newly fetched notifications and add to the notificationMap
		for _, notification := range dbNotifications {
			notificationMap[notification.AccountID] = notification
			CacheNotificationPreference(notification) // Cache the result
		}

		// Process the notifiable status for the fetched notifications
		for idx, account := range accounts {
			if notification, exists := notificationMap[account.ID]; exists {
				if val, ok := notification.Config[topic]; ok {
					if status, ok := val.(bool); ok {
						notifiable[idx] = status
						continue
					}
				}
				notifiable[idx] = true
			}
		}
	}

	return notifiable
}
