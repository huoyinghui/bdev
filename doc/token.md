#### token api

#### 1.new token
```sh
type UserAuth struct {
	Token  string `json:"token"`
	Uid    string `json:"uid"`
	Secret string `json:"secret"`
}
POST: localhost:4008/v1/api/token
{
    "uid": "123456",
    "secret": "123",
}

return:
{
    "secret": "4041810df91001619b7b92ed48cf7a7e",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxMjM0NTYiLCJleHAiOjE1MjI4MTI1MzIsImlzcyI6IjEyMzQ1NiJ9.3BRBEyfoie4hUMjhBvnEEn0oIa5YVJrPQTGwrOqCfQs"
}
```

#### 2.check token
```sh
POST: localhost:4008/v1/api/valid
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxMjM0NTYiLCJleHAiOjE1MjI4MTI1MzIsImlzcyI6IjEyMzQ1NiJ9.3BRBEyfoie4hUMjhBvnEEn0oIa5YVJrPQTGwrOqCfQs"
}

本质一样: token异常时，new_uid为空,uid 仍可以获取.
{
    "new_uid": "123456", //newUid, err := utils.AuthToken(ob.Token, secret) 验证成功后，返回的uid
    "uid": "123456"      // 根据token 中payload 解出来uid
}
```