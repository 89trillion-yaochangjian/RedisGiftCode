
#### 1、整体框架
整体功能的实现思路
礼品码的主要思路是管理人员输入的礼品码信息，以json数据传入，将json数据转存到礼品码信息结构体，并生成8位随机礼品码，将8位随机礼品码为key，礼品信息作为value存入Redis数据库。管理人员根据礼品码查看此礼品码的信息。用户可根据礼品码获取奖励，并看到奖励内容。

#### 2、目录结构
```
.
├── README.md
├── StructInfo
│   ├── GiftCodeInfo.go
│   ├── GiftContentList.go
│   └── ReceiveGiftList.go
├── __pycache__
│   └── locustfile.cpython-39.pyc
├── app
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── ctrl
│   │   └── GiftCodeCtrl.go
│   ├── handler
│   │   └── GiftCodeHandler.go
│   ├── router
│   │   └── GiftCodeRouter.go
│   ├── service
│   │   ├── GiftCodeService.go
│   │   └── GiftCodeService_test.go
│   └── utils
│       ├── GetRandomString.go
│       └── initClient.go
├── locustFile.py
└── report_1625539390.4651678.html


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

使用redis存储数据，ID 为8位随机礼品码 string类型，value为礼品码信息

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

定义异常结构体，统一处理异常


