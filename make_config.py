#!/usr/bin/env python3
#-*- coding:utf-8 -*-


import yaml 
import os
import time
from Crypto.Cipher import AES

def decrypt(text):
    iv = b'oruuejasn49w5ren'
    mode = AES.MODE_CBC
    cryptos = AES.new(passwd, mode, iv)
    text = cryptos.decrypt(text)
    try:
        decode = bytes.decode(text).rstrip('\0')
        return decode
    except:
        return False

def makeconfig(filename, confname):
    with open(f"config/user/{filename}", mode='rb') as f:
        cookie = decrypt(f.read())
    if not cookie:
        print(f'发现密码错误的配置文件{filename}，已经重命名为{filename}.bak，下次运行将会自动删除')
        os.rename(f"config/user/{filename}", f"config/user/{filename}.bak")
        return
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
    files.remove('cache')
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

def load_finish():
    list = os.listdir('mihoyo/config')
    list.remove('cache')
    num = -1
    for i in list:
        if '.yaml' in i:
            num += 1
    if num > 0:
        print(f'配置加载完成！共加载了{num}个配置文件')
    else:
        print(f'配置加载完成！未发现可用的配置文件！')
        exit(1)

def del_useless():
    list = os.listdir('config/user')
    for i in list:
        print(i)
        if 'bak' in i:
            print(i)
            os.remove(f'config/user/{i}')
    
def update_runtime():
    with open(f"config/user/cache", mode='w', encoding='utf-8') as f:
            f.write(time.asctime( time.localtime(time.time()) ))

def add_to_16(text):
    if len(text.encode('utf-8')) % 16:
        add = 16 - (len(text.encode('utf-8')) % 16)
    else:
        add = 0
    text = text + ('\0' * add)
    return text.encode('utf-8')



passwd = add_to_16(os.environ["PASSWORD"])

if __name__ == '__main__':
    update_runtime()
    try:
        del_useless()
    finally:
        get_encrypted()
        load_finish()
