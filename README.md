# go-srv

练手用的，bug随缘了

简易版goserver 目前所有handler都写在msghandler下 若要拆分则需将其他文件的interface写到inet中

消息转发由proto定义 目前source字段有服务器附带发送，可以考虑由本服发送 但会导致重新编码

TODO Logger

    file writter/接入open search

TODO config.json
    docker 替换json


### 加入docker支持
- 构建 sh bin/build_docker.sh
- 运行 sh run_in_docker.sh port (port 默认为9999)