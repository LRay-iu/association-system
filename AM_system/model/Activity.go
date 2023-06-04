package model

import "time"

type Activity struct {
	ActivityID       int       `gorm:"primaryKey;column:activity_id;autoIncrement"`
	AssociationID    int       `gorm:"column:association_id;foreignKey:AssociationID"`
	ActivityTime     time.Time `gorm:"column:activity_time"`
	ActivityHost     string    `gorm:"column:activity_host"`
	ActivityContent  string    `gorm:"column:activity_content"`
	ActivityState    bool      `gorm:"column:activity_state"`
	ActivityRelease  time.Time `gorm:"column:activity_release"`
}

// 设置表名
func (Activity) TableName() string {
	return "activity"
}
