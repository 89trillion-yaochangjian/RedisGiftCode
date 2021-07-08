
#### 1、整体框架
整体功能的实现思路
礼品码的主要思路是管理人员输入的礼品码信息，以json数据传入，将json数据转存到礼品码信息结构体，并生成8位随机礼品码，将8位随机礼品码为key，礼品信息作为value存入Redis数据库。管理人员根据礼品码查看此礼品码的信息。用户可根据礼品码获取奖励，并看到奖励内容。

#### 2、目录结构
```
.
├── README.md
├── __pycache__
│   └── locustfile.cpython-39.pyc
├── app
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── ctrl
│   │   └── GiftCodeCtrl.go
│   ├── dao
│   │   └── GiftCodeDao.go
│   ├── handler
│   │   └── GiftCodeHandler.go
│   ├── router
│   │   └── GiftCodeRouter.go
│   ├── service
│   │   ├── GiftCodeService.go
│   │   └── GiftCodeService_test.go
│   ├── structInfo
│   │   ├── GiftCodeInfo.go
│   │   ├── GiftContentList.go
│   │   ├── ReceiveGiftList.go
│   │   └── ginResult.go
│   └── utils
│       ├── GetRandomString.go
│       └── initClient.go
├── locustFile.py
├── report_1625710932.148289.html
└── 流程图.png


```

#### 3. 代码逻辑分层



|层|文件夹|主要职责|调用关系|其他说明|
| ------------ | ------------ | ------------ | ------------ | ------------ |
|应用层 |app/http/main.go  |服务器启动 |调用路由层工具层   |不可同层调用
|路由层 |internal/router/GiftCodeRouter.go  |路由转发 | 调用工具层 控制层 被应用层   |不可同层调用
|控制层 |internal/ctrl/GiftCodeCtrl.go  |请求参数处理，响应 | 调用handler层 被路由层调用    |不可同层调用
|handler层 |internal/handler/GiftCodeHandler.go  |处理具体业务 | 调用路由层service层，被控制层调用    |不可同层调用
|service层   |internal/service/GiftCodeService.go  |处理业务逻辑 | 调用工具层，被handler层调用    |可同层调用
| 配置文件 |StructInfo  |结构体 | 被service层 、handler层、控制层调用   |不可同层调用

#### 4.存储设计


礼包码信息

|  内容 |数据库   | key  | 类型  |
| ------------ | ------------ | ------------ | ------------ |
|   礼品码描述| Redis  | GiftDes  |  string |
| 礼品码类型  | Redis  |  CodeType |   int|
|  礼品码类型描述 |  Redis | CodeTypeDesc  | string  |
|  已经领取次数 | Redis  |  ReceiveNum |int   |
| 可以领取的次数  |Redis   |  AvailableTimes |  int |
|  有效期 |Redis   |  ValidPeriod | int  |
|  礼包内容 |  Redis | Contents  |   int|
|  礼包码 |  Redis | Code  |  string |
|  创建人 | Redis  |  Creator |  string |
|  创建时间 |  Redis |CreatTime   |time.Time   |
|  礼品内容列表 | Redis  | ContentList  |  GiftContentList |
|  领取列表 |  Redis |  ReceiveList |[]ReceiveGiftList   |
|  指定用户 |  Redis |  User |  string |

礼包领取用户

|  内容 |数据库   | key  | 类型  |
| ------------ | ------------ | ------------ | ------------ |
| 领取用户名  | Redis  |ReceiveUser   | string  |
|  领取时间 | Redis  | ReceiveTime  |time.Time   |

礼包内容

|  内容 |数据库   | key  | 类型  |
| ------------ | ------------ | ------------ | ------------ |
| 金币  | Redis  |GoldCoins   | int  |
|钻石   | Redis  |  Diamonds | int  |
| 道具  |  Redis |Props   |  int |
| 英雄  |Redis   | Heroes  |int   |
|  小兵 | Redis  |  Creeps |  int |


#### 5. 接口设计

http get

   | 接口地址  |  请求参数 |  响应参数|
   | ------------ | ------------ | ------------ |
   |  http://127.0.0.1:8080/VerifyGiftCode |  code |   礼包信息|
   | http://127.0.0.1:8080/GetGiftCodeInfo  |  code，user | 礼品码信息  |

http post

   | 接口地址  |  请求参数 |  响应参数|
   | ------------ | ------------ | ------------ |
   |  http://127.0.0.1:8080/CreateGiftCode |  json=giftCodeInfo |   礼品码|
#### 6. 如何编译执行 

`
go build
`

`
./app
`

#### 7、todo

进一步划分代码，增加代码的可服用以及可扩展行


