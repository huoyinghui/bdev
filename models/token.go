package models

var (
	Users map[string]*UserAuth
)

type UserAuth struct {
	Token  string `json:"token"`
	Uid    string `json:"uid"`
	Secret string `json:"secret"`
}

func init() {
	Users = make(map[string]*UserAuth)
	Users["5ab4bef1c4cd748f32c6dff3"] = &UserAuth{"", "5ab4bef1c4cd748f32c6dff3", "6fd64e1644629c16e7c0f994bcd493b5"}
}
