# webhook

```bash
## route 定义在webhook配置的Payload URL
## path 项目路径执行git pull的地方
## port 监听的端口
## secret 定义在webhook配置的secret
nohup ./webhook_linux_amd64 -route=your_route -path=your_path -port=your_port -secret="your secret" > ./webhook.log 2>&1& echo $! > ./webhook.pid
```
