package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type Roles struct {
	ID        uint32    `gorm:"primaryKey;auto_increment;unique" json:"id"`
	UserEmail string    `gorm:"size:100;not null;foreignkey:Email;" json:"email"`
	Access    string    `gorm:"size:100;not null;" json:"acc"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	CreatedBy string    `gorm:"size:100;" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedBy string    `gorm:"size:100;" json:"-"`
	User      User      `json:"-"`
}

//function get all role
func (r *Roles) FindAllRole(db *gorm.DB) (*[]Roles, error) {
	var err error
	roles := []Roles{}
	err = db.Debug().Model(&Roles{}).Find(&roles).Error
	if err != nil {
		return &[]Roles{}, err
	}
	return &roles, err
}

// function get role by email
func (r *Roles) FindAllRoleBymail(db *gorm.DB, username string) (*Roles, error) {
	var err error
	err = db.Debug().Model(Roles{}).Where("user_email = ?", username).Take(&r).Error
	if err != nil {
		return &Roles{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Roles{}, errors.New("User Nt Found")
	}
	return r, err
}

func (r *Roles) Prepare() {
	r.ID = 0
	r.UserEmail = html.EscapeString(strings.TrimSpace(r.UserEmail))
	r.Access = html.EscapeString(strings.TrimSpace(r.Access))
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *Roles) SaveRole(db *gorm.DB) (*Roles, error) {
	var err error
	err = db.Debug().Create(&r).Error
	if err != nil {
		return &Roles{}, err
	}
	return r, err
}

func (r *Roles) UpdateRole(db *gorm.DB, uid uint32) (*Roles, error) {
	role := Roles{}
	db = db.First(&role, uid)
	if db.Error != nil {
		return &Roles{}, db.Error
	}

	role.Access = r.Access

	db.Save(&role)
	err := db.Debug().Model(&Roles{}).Where("id = ?", uid).Take(&r).Error
	if err != nil {
		return &Roles{}, err
	}
	return r, nil
}

func (r *Roles) DeleteRole(db *gorm.DB, uid uint32) (*Roles, error) {
	role := Roles{}
	db.Delete(&role, uid)
	if db.Error != nil {
		return &Roles{}, db.Error
	}

	return r, nil
}

func (r *Roles) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if r.UserEmail == "" {
			return errors.New("Required Email")
		}
		if r.Access == "" {
			return errors.New("Required Access")
		}
		if err := checkmail.ValidateFormat(r.UserEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if r.UserEmail == "" {
			return errors.New("Required Email")
		}
		if r.Access == "" {
			return errors.New("Required Access")
		}
		if err := checkmail.ValidateFormat(r.UserEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}

}
