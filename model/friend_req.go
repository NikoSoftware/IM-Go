package model

type FriendReq struct {
	Id         int    `gorm:"id" json:"id"`                   // id
	FromUserId int    `gorm:"from_user_id" json:"fromUserId"` // 请求好友用户id
	ToUserId   int    `gorm:"to_user_id" json:"toUserId"`     // 被请求好友用户id
	CreateTime string `gorm:"create_time" json:"createTime"`  // 请求时间
	Message    string `gorm:"message" json:"message"`         // 发送的消息
	Status     int    `gorm:"status" json:"status"`           // 是否已处理
}

func (FriendReq) TableName() string {
	return "tb_friend_req"
}
