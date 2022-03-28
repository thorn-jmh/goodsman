## goodsman

### 技术栈

Database: MongoDB

Web框架: gin

Logger: logrus

cache:  go-cache

### 接口文档

[详见飞书文档](https://xn4zlkzg4p.feishu.cn/wiki/wikcnljIjh0Czj0VtsLD0wt45Od)

### 项目配置

请在主目录下添加`config.yml`文件，并在其中按如下设置：

```yaml
Base:
  RunMode: "debug"
  HttpPort: 1926

App:
  AppID: "id"
  AppSecret: "pwd"

  ManagerID: "3210100000"
	KeyWord:   "借用物品"
	MAX_MONEY: 100.0



Mongo: 
  User: "admin"
  Pwd: "123456"
  Host: "localhost"
  Port: 27017
  DBName: "goodsman"

```

然后在目录下运行`go run ./`启动服务



### 数据库配置

数据库结构文件在./data目录下，请使用`mongorestore`进行导入

数据库部署后需要在goodsman库中新建用户并添加readWrite权限，同时应该在mongoconfig中设置auth选项

字段解释：

```go
records:
  _id:         ObjectID
  employee_id: string
  goods_id:    string
  date:		   int64          #记录更新时间
  state:	   int            #记录状态
  del_number:  int            #变化数量（区分正负）
index{date : -1}    #0：借出，1：归还，2：入库，3：出库
  
goods:
  _id          ObjectID
  goods_id:    string
  goods_auth:  int     #权限 0：实习，1：正式，2：非外借
  number:	   int	   #库存数量
  state:	   int     #物品状态 0：不在库，1：正常，2：报修，3：报损
  owner:	   string  #目前所有者(employee_id)
  goods_msg:
    name:	   string  #名称
    type:	   string  #型号
    photo:	   string  #照片>>>???????
    cost:	   float64 #金额
    consumables:  int  #是否为耗材,0：不是，1：是
index{goods_id : 1,unique}
  
 
```
