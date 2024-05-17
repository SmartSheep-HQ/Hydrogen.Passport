package services

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/datatypes"
	"time"

	"git.solsynth.dev/hydrogen/passport/pkg/database"
	"git.solsynth.dev/hydrogen/passport/pkg/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func GetAccount(id uint) (models.Account, error) {
	var account models.Account
	if err := database.C.Where(models.Account{
		BaseModel: models.BaseModel{ID: id},
	}).First(&account).Error; err != nil {
		return account, err
	}

	return account, nil
}

func LookupAccount(probe string) (models.Account, error) {
	var account models.Account
	if err := database.C.Where(models.Account{Name: probe}).First(&account).Error; err == nil {
		return account, nil
	}

	var contact models.AccountContact
	if err := database.C.Where(models.AccountContact{Content: probe}).First(&contact).Error; err == nil {
		if err := database.C.
			Where(models.Account{
				BaseModel: models.BaseModel{ID: contact.AccountID},
			}).First(&account).Error; err == nil {
			return account, err
		}
	}

	return account, fmt.Errorf("account was not found")
}

func CreateAccount(name, nick, email, password string) (models.Account, error) {
	user := models.Account{
		Name: name,
		Nick: nick,
		Profile: models.AccountProfile{
			Experience: 100,
		},
		Factors: []models.AuthFactor{
			{
				Type:   models.PasswordAuthFactor,
				Secret: HashPassword(password),
			},
			{
				Type:   models.EmailPasswordFactor,
				Secret: uuid.NewString()[:8],
			},
		},
		Contacts: []models.AccountContact{
			{
				Type:       models.EmailAccountContact,
				Content:    email,
				IsPrimary:  true,
				VerifiedAt: nil,
			},
		},
		PermNodes:   datatypes.JSONMap(viper.GetStringMap("permissions.default")),
		ConfirmedAt: nil,
	}

	if err := database.C.Create(&user).Error; err != nil {
		return user, err
	}

	if tk, err := NewMagicToken(models.ConfirmMagicToken, &user, nil); err != nil {
		return user, err
	} else if err := NotifyMagicToken(tk); err != nil {
		return user, err
	}

	return user, nil
}

func ConfirmAccount(code string) error {
	token, err := ValidateMagicToken(code, models.ConfirmMagicToken)
	if err != nil {
		return err
	}

	var user models.Account
	if err := database.C.Where(&models.Account{
		BaseModel: models.BaseModel{ID: *token.AssignTo},
	}).First(&user).Error; err != nil {
		return err
	}

	return database.C.Transaction(func(tx *gorm.DB) error {
		user.ConfirmedAt = lo.ToPtr(time.Now())

		for k, v := range viper.GetStringMap("permissions.verified") {
			if val, ok := user.PermNodes[k]; !ok {
				user.PermNodes[k] = v
			} else if !ComparePermNode(val, v) {
				user.PermNodes[k] = v
			}
		}

		if err := database.C.Delete(&token).Error; err != nil {
			return err
		}
		if err := database.C.Save(&user).Error; err != nil {
			return err
		}

		InvalidAuthCacheWithUser(user.ID)

		return nil
	})
}
