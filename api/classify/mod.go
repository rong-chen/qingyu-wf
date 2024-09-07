package classify

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TableClassify 用户分类
type TableClassify struct {
	gorm.Model
	CId   uuid.UUID `json:"cId" gorm:"c_id"`    // 创建者
	Label string    `json:"label" gorm:"label"` // 名称
}

func (c TableClassify) TableName() string {
	return "classify_table"
}
