package model

import (
	"gorm.io/gorm"
)

// User 定义用户模型
type User struct {
	gorm.Model
	UserID         string `json:"user_id" gorm:"column:user_id;type:varchar(36);uniqueIndex;not null"` // 用户ID，唯一且不能为空
	Username       string `json:"username" gorm:"column:username;type:varchar(100);unique;not null"`   // 用户名，唯一且不能为空
	Role           string `json:"role,omitempty" gorm:"column:role;type:varchar(50);"`                 // 角色，可选，长度为50
	PasswordHashed string `json:"-" gorm:"column:password_hashed;type:varchar(255);not null"`          // 存储哈希后的密码，长度为255
	Email          string `json:"email,omitempty" gorm:"column:email;type:varchar(100);"`              // 邮箱，可选，长度为100
	Phone          string `json:"phone,omitempty" gorm:"column:phone;type:varchar(20);"`               // 电话号码，可选，长度为20
	Avatar         string `json:"avatar,omitempty" gorm:"column:avatar;type:varchar(255);"`            // 头像URL，可选，长度为255
}

type QuesBank struct {
	gorm.Model
	QuesBankID   string `json:"ques_bank_id" gorm:"column:ques_bank_id;type:varchar(36);uniqueIndex;not null"` // 题库ID，唯一且不能为空
	QuesBankName string `json:"ques_bank_name" gorm:"column:ques_bank_name;type:varchar(100);unique;not null"` // 题库名称，唯一且不能为空
	OwnerID      string `json:"owner_id" gorm:"column:owner_id;index;not null"`                                // 拥有者的用户ID，增加索引
	Description  string `json:"description,omitempty" gorm:"column:description;type:text"`                     // 题库描述，文本类型
	IsPublic     bool   `json:"is_public" gorm:"column:is_public;type:boolean;not null;"`                      // 是否公开
}

type Question struct {
	gorm.Model
	Type       string `json:"type" gorm:"column:type;type:varchar(50);not null"`                                 // 题目类型，长度为50，不能为空
	GenID      string `json:"gen_id" gorm:"column:gen_id;type:varchar(36);not null"`                             // 生成ID，唯一且不能为空
	OwerID     string `json:"ower_id" gorm:"column:ower_id;index;"`                                              // 拥有者的用户ID，增加索引
	OwerName   string `json:"ower_name,omitempty" gorm:"column:ower_name;type:varchar(100);"`                    // 拥有者的用户名，长度为100
	Subject    string `json:"subject,omitempty" gorm:"column:subject;type:varchar(50);"`                         // 学科，可选，长度为50
	Grade      string `json:"grade,omitempty" gorm:"column:grade;type:varchar(50);"`                             // 年级，可选，长度为50
	Difficulty string `json:"difficulty,omitempty" gorm:"column:difficulty;type:varchar(50);"`                   // 难度，可选，长度为50
	QuesID     string `json:"ques_id" gorm:"column:ques_id;type:varchar(36);not null;index:idx_ques_id_bank_id"` // 题目ID，不能为空，联合索引
	QuesBankID string `json:"ques_bank_id" gorm:"column:ques_bank_id;index:idx_ques_id_bank_id"`                 // 题库ID，联合索引
	Title      string `json:"title,omitempty" gorm:"column:title;type:varchar(100);"`                            // 题目标题，可选，长度为100
	Question   string `json:"question" gorm:"column:question;type:text"`                                         // 题目内容，文本类型，不能为空
	Answer     string `json:"answer" gorm:"column:answer;type:text"`                                             // 答案，文本类型，不能为空
	Tip        string `json:"tip,omitempty" gorm:"column:prompt;type:text"`                                      // 提示，文本类型
	Parse      string `json:"parse,omitempty" gorm:"column:parse;type:text"`                                     // 解析，文本类型
}
