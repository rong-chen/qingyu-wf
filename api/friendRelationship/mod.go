package friendRelationship

import (
	"github.com/google/uuid"
	"qinyu-wf/api/user"
	"qinyu-wf/global"
)

type FriendRelationship struct {
	global.Model
	UserId     uuid.UUID `json:"userId" gorm:"userId"`
	UserInfo   user.User `json:"userInfo"  gorm:"foreignKey:UserId;references:ID"`
	FriendId   uuid.UUID `json:"friendId" gorm:"friendId"`
	FriendInfo user.User `json:"friendInfo"  gorm:"foreignKey:FriendId;references:ID"`
	Status     string    `json:"status" gorm:"status"` // 1好友，2申请中
}

type RelationshipList struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"userId"`
	FriendId   uuid.UUID `json:"friendId"`
	FriendInfo user.User `json:"friendInfo" `
	Status     string    `json:"status" ` // 1好友，2申请中
}

type CreateParams struct {
	UserId   uuid.UUID `json:"userId" gorm:"userId" binding:"required"`
	FriendId uuid.UUID `json:"friendId" gorm:"friendId" binding:"required"`
}

func (FriendRelationship) TableName() string {
	return "friend_relationship"
}
