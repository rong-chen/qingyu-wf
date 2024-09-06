package friendRelationship

import (
	"gorm.io/gorm"
	"qinyu-wf/global"
)

func ApplyFriend(fr FriendRelationship) error {
	return global.MySql.Create(&fr).Error
}

func FindApplyList(userId interface{}) []RelationshipList {
	var list []FriendRelationship
	global.MySql.Preload("UserInfo").Preload("FriendInfo").Where("friend_id = ?", userId).Find(&list)
	var frl []RelationshipList
	for i := range list {
		frl = append(frl, RelationshipList{
			UserId:     list[i].FriendId,
			FriendId:   list[i].UserId,
			FriendInfo: list[i].UserInfo,
			Id:         list[i].ID,
		})
	}

	return frl
}

func UpdateFriendRelationshipStatus(id, userId, friendId string, newStatus string) error {
	// 使用 Update 更新单个字段
	result := global.MySql.Model(&FriendRelationship{}).
		Where("user_id = ? and friend_id = ? and id = ?", userId, friendId, id).
		Update("status", newStatus)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
