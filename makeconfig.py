import json
import sys
conf = {}
with open("mihoyo/config/config.json.example", mode='r') as f:
    exam_conf = json.loads(f.read())
with open("config/config.json", mode='r') as f:
    user_conf = json.loads(f.read())
conf.update(exam_conf)
conf.update(user_conf)
conf['account']['cookie'] = sys.argv[1]
with open("mihoyo/config/config.json", mode='w') as f:
    f.write(json.dumps(conf, indent=4, separators=(',', ':')))
