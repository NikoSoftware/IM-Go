package model

const (
	NORMAL  = 0
	PICTURE = 1
	MEDIA   = 2
)

type ChatRecord struct {
	Id         int    `gorm:"id" json:"id"`                  // id
	UserId     int    `gorm:"user_id" json:"userId"`         // 用户id
	FriendId   int    `gorm:"friend_id" json:"friendId"`     // 好友id
	HasRead    int    `gorm:"has_read" json:"hasRead"`       // 是否已读
	CreateTime string `gorm:"create_time" json:"createTime"` // 聊天时间
	Del        int    `gorm:"del" json:"del"`                // 是否删除
	Message    string `gorm:"message" json:"message"`        // 消息
	Type       int    `gorm:"type" json:"type"`
	Url        string `gorm:"url" json:"url"`
}

func (ChatRecord) TableName() string {
	return "tb_chat_record"
}
