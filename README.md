# AutoMihoyoBBS_with_Actions
### 依托于Github Actions和[AutoMihoyoBBS](https://github.com/Womsxd/AutoMihoyoBBS)的每日签到

## 项目分支说明
- 该项目具有[GoLang](https://github.com/RebAltair/AutoMihoyoBBS_with_Actions/tree/Go)/Python两个版本，分别由[@Sayunosyjou(RebAltair)](https://github.com/RebAltair)/[Shan_shanHY](https://github.com/ShanshanHY)进行维护，当前分所在支为**Python**版
- 两个版本的设置有所不同，使用前请查看对应的`README.md`

## 使用方法
1. [Fork](https://github.com/ShanshanHY/AutoMihoyoBBS_with_Actions/fork)本项目
2. 根据[上游项目](https://github.com/Womsxd/AutoMihoyoBBS/blob/master/config/config.yaml.example)编辑本项目的`config/config.yaml`，但不要将你的Cookie填入 **【千万不要在这里填写Cookie！】**
3. 点击`Settings`-`Secrets`-`Actions`，新建一个名为`PASSWORD`的`Secret`并**输入一个8位以上的密码**
4. 新建一个名为`NAME`的`Secret`，并填入任意信息作为配置文件名称
5. 新建一个名为`COOKIE`的`Secret`，并填入你的`Cookie信息`
6. 点击`Actions`并同意使用Workflow，点击`添加账号`-`Run Workflow`，用于生成配置文件
7. 点击`签到`-`Run Workflow`检查是否运行正常

## 进阶项目
### 多账号
- 本项目可支持多账号签到，可调用上游项目的多账号功能进行签到，只需要以不同名称重复添加账号即可

### 用户管理
- 所有的用户信息均保存在`config/user`内保存，并以添加配置时的名称命名，可自由删除
- 用户配置使用`PASSWORD`进行`AES加密`，具体过程请自行查看python脚本
- 用户配置内仅存储了`cookie`，`stoken`用于登录，其他信息通过读取本仓库的`config/config.yaml`与上游仓库的[config/config.yaml.example](https://github.com/Womsxd/AutoMihoyoBBS/blob/master/config/config.yaml.example)进行生成

### 推送相关
- [点此查看](https://github.com/RebAltair/AutoMihoyoBBS_with_Actions/blob/python/PUSH.md)

## 使用声明&注意事项
- 本项目为个人制作，每次运行均为Fork仓库[Womsxd/AutoMihoyoBBS](https://github.com/Womsxd/AutoMihoyoBBS)并生成配置文件
- 上游仓库作者[不建议使用Actions运行](https://github.com/Womsxd/AutoMihoyoBBS#%E5%85%B3%E4%BA%8E%E4%BD%BF%E7%94%A8-github-actions-%E8%BF%90%E8%A1%8C)，遇到问题请勿向上游仓库反馈！！！
- 本项目仅在workflow内使用python生成明文配置文件，如因为个人操作不当或其他原因导致Cookie泄露，本项目不承担任何责任！！！
- 加密后的配置文件保存在公开仓库中，如果对账号安全性担忧，请将仓库设置成Private！
- 验证码问题非本项目bug，请自行寻找解决方法

## 关于许可证
本项目使用[MIT](https://spdx.org/licenses/MIT)许可证进行分发
