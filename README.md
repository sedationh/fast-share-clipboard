# 快速分享剪贴板

# 设计思路

管理端
- 接收消息、向创建链接的客户端发送消息

客户端
- 接收消息、同步消息到剪贴板
- 监听剪贴板变化、向服务器发送消息

提供 GUI 页面
选择 管理端 or 客户端

# 实现核心

- 通讯
- 剪贴板