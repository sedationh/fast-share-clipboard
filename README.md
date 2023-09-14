# 快速分享剪贴板

在两台机器上自动同步剪贴板

## 如何使用

```zsh
➜ ./mac-64-fast-share-clipboard --help
Usage of ./mac-64-fast-share-clipboard:
  -h, --host string   主机，非 admin 角色需要指定 host (default "127.0.0.1")
  -p, --port string   端口 (default "8899")
  -r, --role string   角色 (default "admin")
pflag: help requested


# 一台主机先以 admin 的身份开启
# A computer
➜ ./mac-64-fast-share-clipboard
admin 8899 127.0.0.1
config main.Config{Role:"admin", Port:"8899", Host:"127.0.0.1"}
2023/09/14 09:23:51 admin local host: 192.168.31.35

# 另一台再进行链接，链接 192.168.31.35, -r 的参数可以是任意的，只要不是 admin 就好
./mac-64-fast-share-clipboard -h 192.168.31.35 -r xxx
```

## 设计思路

管理端

- 接收消息、向创建链接的客户端发送消息

客户端

- 接收消息、同步消息到剪贴板
- 监听剪贴板变化、向服务器发送消息

提供 GUI 页面
选择 管理端 or 客户端

## 实现核心

- 通讯
- 剪贴板

# TODO

- [ ] 增加密码校验机制