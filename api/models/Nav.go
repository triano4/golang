package models

import "github.com/jinzhu/gorm"

type Nav struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	ParentId int    `json:"ParentId"`
	SortId   int    `json:"SortId"`
	NavIcon  string `json:"NavIcon"`
	NavPath  string `json:"NavPath"`
	NavAcc   string `json:"NavAcc"`
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
