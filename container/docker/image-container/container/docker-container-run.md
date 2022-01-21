
Run a command in a new container:  
```
docker container run [OPTIONS] IMAGE [COMMAND] [ARG...]
```

  

`-it` 参数告诉 Docker 开启容器的交互模式，并将读者当前的 Shell 连接到 容器 终端。   
按 Ctrl - PQ 组合键，可以在退出容器的同时还保持容器运行。  


```
-i                  Keep STDIN open      even if not attached
-t	 	            Allocate a pseudo-TTY
    --name	        Assign a name to the container
-d  --detach 		Run container in background and print container ID
-p  --publish 		Publish a container's port(s) to the host
-e  --env	    	Set environment variables
    --rm	       	Automatically remove the container when it exits
```
