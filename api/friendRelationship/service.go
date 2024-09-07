package friendRelationship

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"qinyu-wf/api/user"
	"qinyu-wf/global"
)

// ApplyFriend 添加等待同意表格数据
func ApplyFriend(fr AwaitingAgreeTable) error {
	return global.MySql.Create(&fr).Error
}

// FindAwaitingAgreeTable 寻找当前用户，等待同意的表格
func FindAwaitingAgreeTable(userId, FriendId uuid.UUID) (a AwaitingAgreeTable) {
	global.MySql.Where("user_id = ? and friend_id = ?", userId, FriendId).First(&a)
	return
}

// FindApplyList 查询等待同意的表
func FindApplyList(userId interface{}) []RelationshipList {
	var list []AwaitingAgreeTable
	global.MySql.Where("friend_id = ? and status = ?", userId, "2").Find(&list)
	var frl []RelationshipList
	for i := range list {
		frl = append(frl, RelationshipList{
			UserId:     list[i].FriendId,
			FriendId:   list[i].UserId,
			FriendInfo: user.SearchDb("id", list[i].UserId.String()),
			Id:         list[i].ID,
		})
	}
	return frl
}

// UpdateAwaitAgreeTableStatus 变更等待同意的表格状态
func UpdateAwaitAgreeTableStatus(id, newStatus string) error {
	result := global.MySql.Model(&AwaitingAgreeTable{}).
		Where("id = ?", id).
		Update("status", newStatus)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// FindAwaitAgreeTable 根据id查询等待同意的表格
func FindAwaitAgreeTable(id string) (a AwaitingAgreeTable) {
	global.MySql.Where("id = ?", id).First(&a)
	return
}

// CreateRelationshipList 创建用户关系表
func CreateRelationshipList(fr FriendRelationship) (e error) {
	e = global.MySql.Create(&fr).Error
	return
}
