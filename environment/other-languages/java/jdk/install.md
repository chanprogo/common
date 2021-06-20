
`wget https://repo.huaweicloud.com/java/jdk/8u151-b12/jdk-8u151-linux-x64.tar.gz`

`mkdir /usr/java && tar -C /usr/java -zxvf jdk-8u151-linux-x64.tar.gz`  
`chmod -R 777 /usr/java/`  






`vim /etc/profile`  

```
export JAVA_HOME=/usr/java/jdk1.8.0_151/  
export CLASSPATH=$CLASSPATH:$JAVA_HOME/lib/tools.jar:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib  
export PATH=$PATH:$JAVA_HOME/bin  
```

`source /etc/profile`  














<br><br><br>

执行（如果是平时的开发环境，建议在 root 和普通账号下都执行）以下命令：  
`echo "export PATH=$PATH:$JAVA_HOME/bin" >> ~/.bashrc`   
`source ~/.bashrc`  











<br><br><br>

`java -version`  