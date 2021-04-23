package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Timesheet struct {
	ID          int       `json:"id"`
	Date        string    `gorm:"default:CURRENT_TIMESTAMP" json:"Tanggal"`
	Activity    string    `json:"Activity"`
	ProjectName string    `json:"ProjectName"`
	WorkDes     string    `json:"WorkDes"`
	StartTime   time.Time `gorm:"default:TIME" json:"StartTime"`
	EndTime     time.Time `gorm:"default:TIME" json:"EndTime"`
	BreakHours  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"BreakHours"`
	NickName    string    `json:"NickName"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Timesheet) GetTimeSheet(db *gorm.DB) (*[]Timesheet, error) {
	var err error
	ts := []Timesheet{}
	err = db.Debug().Model(&Timesheet{}).Find(&ts).Error
	if err != nil {
		return &[]Timesheet{}, err
	}
	return &ts, err
}

func (t *Timesheet) Prepare() {
	t.ID = 0
	t.Date = html.EscapeString(strings.TrimSpace(t.Date))
	t.Activity = html.EscapeString(strings.TrimSpace(t.Activity))
	t.ProjectName = html.EscapeString(strings.TrimSpace(t.ProjectName))
	t.WorkDes = html.EscapeString(strings.TrimSpace(t.WorkDes))
}

func (t *Timesheet) SaveTimesheet(db *gorm.DB) (*Timesheet, error) {
	var err error
	err = db.Debug().Create(&t).Error
	if err != nil {
		return &Timesheet{}, err
	}
	return t, nil
}

func (t *Timesheet) UpdateTimesheet(db *gorm.DB, uid uint32) (*Timesheet, error) {
	ts := Timesheet{}
	db = db.First(&ts, uid)
	if db.Error != nil {
		return &Timesheet{}, db.Error
	}

	ts.Activity = t.Activity

	db.Save(&ts)
	err := db.Debug().Model(&Timesheet{}).Where("id = ?", uid).Take(&ts).Error
	if err != nil {
		return &Timesheet{}, err
	}
	return t, nil
}
