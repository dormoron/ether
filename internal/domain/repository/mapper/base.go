package mapper

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int64          `gorm:"primarykey;autoIncrement;comment:'id'"`
	CreatedAt time.Time      `gorm:"column:create_time;comment:'创建时间'"`
	UpdatedAt time.Time      `gorm:"column:update_time;comment:'修改时间'"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:'软删除'"`
}
