
`docker-compose down [options]` 停止和删除容器、网络、卷、镜像。  

选项包括：
```
    --rmi type  删除镜像，类型必须是：all，删除compose文件中定义的所有镜像；local，删除镜像名为空的镜像
-v, --volumes         删除已经在 compose 文件中 定义的 和匿名的 附在容器上的 数据卷
    --remove-orphans 删除服务中 没有在 compose 中定义的容器
```
