## 字节青训营
极简版抖音服务端大项目

#### 拉取项目
```git clone https://github.com/lemuzhi/douyin.git```

#### 1、使用Goland或vscode打开项目
#### 2、修改config/config.toml配置
    将gin的地址修改为自己的而是内网地址，不是127.0.0.1或者localhost，一般以192.168开头，这样在前端就能连接测试了
    mysql只需将账号密码改成自己的，之后会自动建库建表
    redis如有变动可修改，否则默认即可
#### 2、运行main.go文件里面的main函数或在终端输入如下命令：
```go run main,go```

