#!/usr/bin/env python3
#-*- coding:utf-8 -*-


import yaml 
import os
import sys
from Crypto.Cipher import AES

sys.path.append("mihoyo")

import login, config



def add_to_16(text):
    if len(text.encode('utf-8')) % 16:
        add = 16 - (len(text.encode('utf-8')) % 16)
    else:
        add = 0
    text = text + ('\0' * add)
    return text.encode('utf-8')

def check_passwd():
    if passwd == '':
        print('请设置用于加密配置文件的密码！')
        exit(1)
    elif len(passwd) < 8:
        print('密码长度需要大于8位！')
        exit(1)

def update_cookie():
    with open("mihoyo/config/config.yaml.example", mode='r', encoding='utf-8') as f:
        conf = yaml.load(f.read(), Loader=yaml.FullLoader)
    conf['account']['cookie'] = cookie
    with open("mihoyo/config/config.yaml", mode='w', encoding='utf-8') as f:
        f.write(yaml.dump(conf))
    config.load_config()
    login.login()
    with open("mihoyo/config/config.yaml", mode='r', encoding='utf-8') as f:
        conf = yaml.load(f.read(), Loader=yaml.FullLoader)
    if conf['account']['cookie'] == 'CookieError':
        exit(1)
    else :
        if not os.path.exists('config/user'):
            os.makedirs('config/user')
        with open(f"config/user/{name}", mode='wb') as f:
            f.write(encryption(conf['account']))
    print('配置保存成功！')

def encryption(text):
    iv = b'oruuejasn49w5ren'
    mode = AES.MODE_CBC
    aes = AES.new(passwd, mode, iv)
    code = aes.encrypt(add_to_16(str(text)))
    return code



passwd = add_to_16(os.environ["PASSWORD"])
cookie = os.environ["COOKIE"]
name = os.environ["NAME"]


if __name__ == '__main__':
    check_passwd()
    update_cookie()
