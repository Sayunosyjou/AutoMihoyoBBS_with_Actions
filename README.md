# AutoMihoyoBBS_with_Actions
依托于Github Actions和[AutoMihoyoBBS](https://github.com/Womsxd/AutoMihoyoBBS)的每日签到
## 使用方法
[Fork](https://github.com/ShanshanHY/AutoMihoyoBBS_with_Actions/fork)本项目

根据[原项目](https://github.com/Womsxd/AutoMihoyoBBS)编辑本项目的`DefaultExampleConfig.yaml`，但不要将你的Cookie填入 ***【千万不要在这里填写Cookie！】***

点击`Settings`-`Secrets`-`Actions`，新建一个名为`COOKIE`的`Secret`并 ***在此处填写Cookie***

再次新建一个名为`PUSH`的`Secret`，并在此处填写您的PushDeer Key

点击`Actions`并同意使用Workflow，点击`AutoMihoyoBBS_with_Actions`-`Run Workflow`检查是否运行正常

## 关于本项目的声明
本项目为个人制作，每次运行均为Fork仓库[Womsxd/AutoMihoyoBBS](https://github.com/Womsxd/AutoMihoyoBBS)并生成配置文件

上游仓库作者[不建议使用Actions运行](https://github.com/Womsxd/AutoMihoyoBBS#%E5%85%B3%E4%BA%8E%E4%BD%BF%E7%94%A8-github-actions-%E8%BF%90%E8%A1%8C)，遇到问题请勿向上游仓库反馈！！！

本项目仅通过CI生成配置文件，如因为[个人操作不当](https://github.com/ShanshanHY/AutoMihoyoBBS_with_Actions/#%E4%BD%BF%E7%94%A8%E6%96%B9%E6%B3%95)或其他原因导致Cookie泄露，本项目不承担任何责任！！！

## 关于本项目的一些细节说明
本项目使用GoLand编写，并且引入了一些第三方库包括但不限于[go-yaml](https://github.com/go-yaml/yaml)

## 本项目使用MIT许可证进行分发