#!/usr/bin/env python3
#-*- coding:utf-8 -*-


import os
import configparser
import yaml

push = os.environ["PUSH"]
push_config = os.environ["PUSH_CONFIG"]

def copy_info(conf):
    push_yaml = push_config.replace(' ', '').replace('api_url=', 'api_url: ').replace('token=', 'token: ')
    yaml_info = yaml.load(push_yaml, Loader=yaml.FullLoader)
    if type(yaml_info) is dict:
        for i in yaml_info:
            conf.set(push, i, str(yaml_info[i]))
    else:
        print('请输入正确的配置文件（yaml类型/复制原始ini）')

def make_push_conf():
    if len(push) and len(push_config):
        conf = configparser.ConfigParser()
        conf.read('mihoyo/config/push.ini.example', encoding="utf-8")
        conf.set('setting', 'push_server', push)
        conf.set('setting', 'enable', 'true')
        if push == 'server酱' or  push == 'pushplus':
            conf.set('setting', 'push_token', push_config)
        
        elif push == 'pushdeer' or push == 'bark':
            if 'api_url' in push_config:
                copy_info(conf)
            else:
                conf.set(push, 'token', push_config)

        elif push == 'telegram' or  push == 'wecom' or  push == 'cqhttp' or  push == 'gotify' or  push == 'dingrobot':
            copy_info(conf)   
        else:
            print('推送方式有误！目前配置生成只支持"cqhttp" "server酱" "pushplus" "pushdeer" "telegram" "wecom" "dingrobot" "bark" "gotify"')
            conf.set('setting', 'enable', 'false')
        with open('mihoyo/config/push.ini', 'w', encoding='utf-8') as f:
            conf.write(f)
    else:
        print('请设置PUSH和PUSH_CONFIG')

if __name__ == '__main__':
    make_push_conf()
