package model

type JoinIn struct {
	StuID         int `gorm:"primaryKey;column:stu_id;foreignKey:StuID;"`
	AssociationID int `gorm:"primaryKey;column:association_id;foreignKey:AssociationID"`
}

// 设置表名
func (JoinIn) TableName() string {
	return "join_in"
}
