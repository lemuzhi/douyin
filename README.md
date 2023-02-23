## 字节青训营
极简版抖音服务端大项目

#### 拉取项目
```git clone https://github.com/lemuzhi/douyin.git```

#### 1、使用Goland或vscode打开项目
#### 2、拷贝config.toml.example，命名为config.toml，并修改其中的配置为你当前的配置，（出于安全考虑，未将服务器和数据库地址账号密码等数据上传）
    将gin的地址修改为自己的而是内网地址，不是127.0.0.1或者localhost，一般以192.168开头，这样在前端就能连接测试了
    mysql只需将账号密码改成自己的，之后会自动建库建表
    redis如有变动可修改，否则默认即可
#### 2、运行main.go文件里面的main函数或在终端输入如下命令：
```go run main,go```

### 开发流程
1、打开终端，git pull更新到最新    

2、git branch -a查看所有分支

3、git checkout 分支名，切换到属于自己的分支

4、开始code

5、完成code后，git push origin 分支名，将代码提交到自己的分支

6、打开github，进入仓库，提交pull requests
