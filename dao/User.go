package dao

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"` //用户名
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

func (User) TableName() string {
	return "go_user"
}

type ResultUser struct {
	Code string `json:"code"`
	Data []User `json:"data"`
	Msg  string `json:"msg"`
}
