package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Nav struct {
	ID        string    `json:"id"`
	Name      string    `json:"Name"`
	ParentId  string    `json:"ParentId"`
	SortId    int       `json:"SortId"`
	NavIcon   string    `json:"NavIcon"`
	NavPath   string    `json:"NavPath"`
	NavAcc    string    `json:"NavAcc"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Navs []Nav

//function get all menu
func (n *Nav) FindAllMenu(db *gorm.DB) (*[]Nav, error) {
	var err error
	navs := []Nav{}
	err = db.Debug().Model(&Nav{}).Find(&navs).Error
	if err != nil {
		return &[]Nav{}, err
	}
	return &navs, err
}

//Save menu function
func (n *Nav) SaveNav(db *gorm.DB) (*Nav, error) {
	var err error
	err = db.Debug().Create(&n).Error
	if err != nil {
		return &Nav{}, err
	}
	return n, nil
}

//Delete menu function
func (n *Nav) DeleteNav(db *gorm.DB, uid uint32) (*Nav, error) {
	navs := Nav{}
	db.Delete(&navs, uid)
	if db.Error != nil {
		return &Nav{}, db.Error
	}

	return n, nil
}

func (n *Nav) Validate(action string) error {
	switch strings.ToLower(action) {
	default:
		if n.Name == "" {
			return errors.New("Required Name")
		}
		if n.SortId == 0 {
			return errors.New("Required Id Sorting")
		}
		if n.NavPath == "" {
			return errors.New("Required Menu Path")
		}
		return nil

	}
}
