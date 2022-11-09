package model

type friend struct {
	Id         int    `gorm:"id" json:"id"`
	UserId     int    `gorm:"user_id" json:"userId"`         // 用户id
	FriendId   int    `gorm:"friend_id" json:"friendId"`     // 好友id
	Comment    string `gorm:"comment" json:"comment"`        // 备注
	CreateTime string `gorm:"create_time" json:"createTime"` // 添加好友日期
	Del        int    `gorm:"del" json:"del"`
}

func (friend) TableName() string {
	return "tb_friend"
}
