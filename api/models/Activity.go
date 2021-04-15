package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Activity struct {
	IdActivity   string    `json:"id"`
	DescActivity string    `json:"activity"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Activities []Activity

func (a *Activity) FindAllActivity(db *gorm.DB) (*[]Activity, error) {
	var err error
	act := []Activity{}
	err = db.Debug().Model(&Activity{}).Find(&act).Error
	if err != nil {
		return &[]Activity{}, err
	}
	return &act, err
}

func (a *Activity) FindActivityById(db *gorm.DB, id string) (*Activity, error) {
	err := db.Debug().Model(Roles{}).Where("id = ?", id).Take(&a).Error
	if err != nil {
		return &Activity{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Activity{}, errors.New("Activity Nt Found")
	}
	return a, err
}

func (a *Activity) SaveActivity(db *gorm.DB) (*Activity, error) {
	err := db.Debug().Create(&a).Error
	if err != nil {
		return &Activity{}, err
	}
	return a, nil
}

func (a *Activity) UpdateActivity(db *gorm.DB, uid uint32) (*Activity, error) {

	// To hash the password
	// err := u.BeforeSave()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	db = db.Debug().Model(&Activity{}).Where("id = ?", uid).Take(&Activity{}).UpdateColumns(
		map[string]interface{}{
			"activity":  a.DescActivity,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Activity{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&a).Error
	if err != nil {
		return &Activity{}, err
	}
	return a, nil
}

func (a *Activity) DeleteActivity(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Activity{}).Where("id = ?", uid).Take(&Activity{}).Delete(&Activity{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
