# bdev
bdev
#### 站内信

##### 1: 未读消息 toUserId: 7003496620
```bash
/innerMsg/unreadDetail/7003496620?
```

##### 2: 全部消息 
```bash
/innerMsg/allDetail/7003714351?
```

##### 3: 已读消息
```bash
#/innerMsg/read/0
/innerMsg/readDetail/7003714351?categoryId=0

```
#####  读取消息
```sh
localhost:5008/v1/api/msg/read/to?pageSize=5&pageNumber=1&msgType=UNSEEN
localhost:5008/v1/api/msg/read/to?pageSize=5&pageNumber=1&msgType=SEEN
localhost:5008/v1/api/msg/read/to?pageSize=5&pageNumber=1&msgType=ALL
post
json:
{
    "pageSize": 5,
    "pageNumber": 1,
    "msgType": "SEEN",  #SEEN/UNSEEN/ALL/DETAIL
    "id": -1,           #当msgType为DETAIL是，id有效
    "FromUserId": 115,  #发信人
    "toUserId": 0       #收件信
}
```

##### 4.分页设计
```bash
https://notifications.console.aliyun.com/message/getMessageList.json?pageNumber=2&pageSize=10&status=1
```

##### 4.1 分页测试
```sh
POST localhost:5008/v1/api/msg/read/to?pageSize=5&pageNumber=2
{
    "from_user_id": 100,
    "to_user_id": 1
}

{
    "data": [
        {
            "Id": 6,
            "FromUserId": 105,
            "ToUserId": 0,
            "CreatedAt": "2018-04-03T10:43:32.580822+08:00",
            "UpdateAt": "2018-04-03T10:43:32.580822+08:00",
            "Title": "Welcome to join APKPURE developer",
            "Message": "1.you xxx 2.xxxxx 3.xxxx",
            "IsDelete": false,
            "Status": "UNSEEN"
        },
        {
            "Id": 7,
            "FromUserId": 106,
            "ToUserId": 0,
            "CreatedAt": "2018-04-03T10:43:32.580827+08:00",
            "UpdateAt": "2018-04-03T10:43:32.580827+08:00",
            "Title": "Welcome to join APKPURE developer",
            "Message": "1.you xxx 2.xxxxx 3.xxxx",
            "IsDelete": false,
            "Status": "UNSEEN"
        },
        ...
        {
            "Id": 10,
            "FromUserId": 109,
            "ToUserId": 0,
            "CreatedAt": "2018-04-03T10:43:32.580835+08:00",
            "UpdateAt": "2018-04-03T10:43:32.580836+08:00",
            "Title": "Welcome to join APKPURE developer",
            "Message": "1.you xxx 2.xxxxx 3.xxxx",
            "IsDelete": false,
            "Status": "UNSEEN"
        }
    ],
    "err": "",
    "nums": 5
}
```

#### API read  统一升级为post
```sh
localhost:5008/v1/api/msg/read
DETAIL:
{
	"pageSize": 5,
	"pageNumber": 1,
	"id": 18,
	"msgType": "DETAIL",
    "FromUserId": 116,
    "toUserId": 0
}
UNSEEN:
{
	"pageSize": 5,
	"pageNumber": 1,
	"id": 18,
	"msgType": "UNSEEN",
    "FromUserId": 116,
    "toUserId": 0
}
SEEN:
{
	"pageSize": 5,
	"pageNumber": 1,
	"id": 18,
	"msgType": "SEEN",
    "FromUserId": 116,
    "toUserId": 0
}
ALL:
{
	"pageSize": 5,
	"pageNumber": 1,
	"id": 18,
	"msgType": "ALL",
    "FromUserId": 116,
    "toUserId": 0
}
```
##### API create  统一升级为post
```sh
localhost:5008/v1/api/msg/create
{
    "FromUserId": 116,
    "toUserId": 1,
    "Title": "Welcome to join APKPURE developer 1",
    "Message": "1.you xxx 2.xxxxx 3.xxxx"
}
```
##### API update  统一升级为post
```sh
localhost:5008/v1/api/msg/update
{
	"id": 20,           #mid 唯一识别msg
    "status": "SEEN"    #状态为已读
}

// 用户只能更新status 为已读未读，不可以修改其他message 的原始信息，网站后台可以修改, 网站😨
	msg, err := models.UpdateMessage(ob.Id, &models.Message{
		Id: ob.Id,
		Status: ob.Status,
	})
```

##### API delete  统一升级为post
```sh
{
	"id": 20          # 删除
}
// 通过id检索，只跟新状态
models.DeleteMessage(ob.Id);
```
