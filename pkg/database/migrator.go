package database

import (
	"git.solsynth.dev/hydrogen/identity/pkg/models"
	"gorm.io/gorm"
)

var DatabaseAutoActionRange = []any{
	&models.Account{},
	&models.AuthFactor{},
	&models.AccountProfile{},
	&models.AccountPage{},
	&models.AccountContact{},
	&models.AccountFriendship{},
	&models.AuthSession{},
	&models.AuthChallenge{},
	&models.MagicToken{},
	&models.ThirdClient{},
	&models.ActionEvent{},
	&models.Notification{},
	&models.NotificationSubscriber{},
}

func RunMigration(source *gorm.DB) error {
	if err := source.AutoMigrate(DatabaseAutoActionRange...); err != nil {
		return err
	}

	return nil
}
