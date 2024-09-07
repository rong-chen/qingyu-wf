package friendRelationship

import (
	"github.com/google/uuid"
	"qingyu-wf/api/user"
	"qingyu-wf/global"
)

// FriendRelationship 用户关系表
type FriendRelationship struct {
	global.Model
	UserId   uuid.UUID `json:"userId" gorm:"column:user_id"`
	FriendId uuid.UUID `json:"friendId" gorm:"column:friend_id"`
	Status   string    `json:"status" gorm:"column:status"` // 1好友，2情侣，3黑名单
}

// AwaitingAgreeTable 待通过申请好友表
type AwaitingAgreeTable struct {
	global.Model
	UserId   uuid.UUID `json:"userId" gorm:"userId"`
	FriendId uuid.UUID `json:"friendId" gorm:"friendId"`
	Status   string    `json:"status" gorm:"column:status"` // 1同意，2等待，3拒绝
}

type RelationshipList struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"userId"`
	FriendId   uuid.UUID `json:"friendId"`
	FriendInfo user.User `json:"friendInfo"`
}

type CreateParams struct {
	UserId   uuid.UUID `json:"userId" gorm:"userId" binding:"required"`
	FriendId uuid.UUID `json:"friendId" gorm:"friendId" binding:"required"`
}

func (FriendRelationship) TableName() string {
	return "friend_relationship"
}
func (AwaitingAgreeTable) TableName() string {
	return "awaiting_agree_relationship"
}
