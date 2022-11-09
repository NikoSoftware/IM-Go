package model

type User struct {
	Id         int    `gorm:"id" json:"id"`                  // ID
	UserName   string `gorm:"user_name" json:"userName"`     // 用户名
	Password   string `gorm:"password" json:"password"`      // 密码
	PicNormal  string `gorm:"pic_normal" json:"picNormal"`   // 头像（正常尺寸）
	NickName   string `gorm:"nick_name" json:"nickName"`     // 昵称
	Qrcode     string `gorm:"qrcode" json:"qrcode"`          // 二维码
	ClientId   string `gorm:"client_id" json:"clientId"`     // 手机端唯一ID
	Sign       string `gorm:"sign" json:"sign"`              // 签名
	CreateTime string `gorm:"create_time" json:"createTime"` // 注册日期
	Phone      string `gorm:"phone" json:"phone"`            // 绑定手机
}

func (User) TableName() string {
	return "tb_user"
}
