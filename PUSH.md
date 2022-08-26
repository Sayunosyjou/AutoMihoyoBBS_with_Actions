# 推送生成
### 本项目同样使用`Secrets`动态生成推送文件
## 使用方法
1. 新建一个名为`PUSH`的Secret，并填入以下推送方式中的一个：`cqhttp`, `server酱`, `pushplus`, `pushdeer`, `telegram`, `wecom`, `dingrobot`, `bark`, `gotify`

2. 新建一个名为`PUSH_CONFIG`的Secret:

3. - 如果你的推送方式为`server酱`/`pushplus`/`pushdeer(官方)`/`bark(官方)`，直接填写对应的`token`
   - 如果你的推送方式为`telegram`/`wecom`/`dingrobot`/`cqhttp`/`gotify`/`pushdeer(自架)`/`bark(自架)`，请根据[此文件](https://github.com/Womsxd/AutoMihoyoBBS/blob/master/config/push.ini.example)编写对应的配置信息

### 举个栗子：
- **pushdeer(官方)**

PUSH:
```
pushdeer
```
PUSH_CONFIG:
```
PDU*******************
```

- **pushdeer(自架)【查看[此文件](https://github.com/Womsxd/AutoMihoyoBBS/blob/master/config/push.ini.example)】**

PUSH:
```
pushdeer
```
PUSH_CONFIG:
```
api_url=https://push.****.cn
token=PDU*******************
```
