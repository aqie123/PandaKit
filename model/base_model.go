package model

import (
	"github.com/XM-GO/PandaKit/time2"
	"gorm.io/gorm"
)

// BaseAutoModel 使用代码生成需要此，不能自由命名id
type BaseAutoModel struct {
	Id        int            `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
	CreatedAt time2.JSONTime `gorm:"column:create_time" json:"createTime" form:"createTime"`
	UpdatedAt time2.JSONTime `gorm:"column:update_time" json:"updateTime" form:"updateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:delete_time" sql:"index" json:"-"`
}

// BaseAutoModelD 表想要硬删除使用他
type BaseAutoModelD struct {
	Id        int            `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
	CreatedAt time2.JSONTime `gorm:"column:create_time" json:"createTime" form:"createTime"`
	UpdatedAt time2.JSONTime `gorm:"column:update_time" json:"updateTime" form:"updateTime"`
}

type BaseModel struct {
	CreatedAt time2.JSONTime `gorm:"column:create_time" json:"createTime" form:"createTime"`
	UpdatedAt time2.JSONTime `gorm:"column:update_time" json:"updateTime" form:"updateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:delete_time" sql:"index" json:"-"`
}

// BaseModelD 硬删除
type BaseModelD struct {
	CreatedAt time2.JSONTime `gorm:"column:create_time" json:"createTime" form:"create_time"`
	UpdatedAt time2.JSONTime `gorm:"column:update_time" json:"updateTime" form:"update_time"`
}
