#!/usr/bin/env python3
#-*- coding:utf-8 -*-


import yaml 
import os
import sys
from Crypto.Cipher import AES

def decrypt(text):
    iv = b'oruuejasn49w5ren'
    mode = AES.MODE_CBC
    cryptos = AES.new(passwd, mode, iv)
    text = cryptos.decrypt(text)
    return bytes.decode(text).rstrip('\0')

def makeconfig(filename, confname):
    with open(f"config/user/{filename}", mode='rb') as f:
        cookie = decrypt(f.read())
    with open("mihoyo/config/config.yaml.example", mode='r', encoding='utf-8') as f:
        conf = yaml.load(f.read(), Loader=yaml.FullLoader)
    with open("config/config.yaml", mode='r', encoding='utf-8') as f:
        user_conf = yaml.load(f.read(), Loader=yaml.FullLoader)
    dicMeg(conf, user_conf)
    conf['account'] = eval(cookie)
    with open(f"mihoyo/config/{confname}.yaml", mode='w', encoding='utf-8') as f:
        f.write(yaml.dump(conf))

def dicMeg(dic1,dic2):
    for i in dic2:
        if i in dic1:
            if type(dic1[i]) is dict and type(dic2[i]) is dict:
                dicMeg(dic1[i],dic2[i])
            else:
                dic1[i] = dic2[i]
        else:
            dic1[i] = dic2[i]

def get_encrypted():
    files = os.listdir('config/user/')
    print(f'发现了{len(files)}个配置文件')
    if len(files) == 0:
        print('请添加一个账号后再次运行！')
    if len(files) > 1:
        for i in files:
            makeconfig(i, i)
        with open(f"run.sh", mode='w', encoding='utf-8') as f:
            f.write('python3 mihoyo/main_multi.py autorun')
    else:
        makeconfig(files[0], 'config')
        with open(f"run.sh", mode='w', encoding='utf-8') as f:
            f.write('python3 mihoyo/main.py')

def add_to_16(text):
    if len(text.encode('utf-8')) % 16:
        add = 16 - (len(text.encode('utf-8')) % 16)
    else:
        add = 0
    text = text + ('\0' * add)
    return text.encode('utf-8')



passwd = add_to_16(os.environ["PASSWORD"])

if __name__ == '__main__':
    get_encrypted()
