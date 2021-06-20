
`mkdir /usr/local/qconf`  // 应该不用执行，此目录会自动创建  

`git clone https://github.com/Qihoo360/QConf.git`  

`cd QConf && mkdir build && cd build`  

`cmake ..`  

`make`  

`make install`  






























`yum install *gdbm*`  







## make 报错

缺少 libgdbm.a 静态库  

```
make[2]: *** No rule tomake target `../agent/../deps/gdbm/_install/lib/libgdbm.a', needed by`agent/qconf_agent'.  Stop.
```

自己上网（https://ftp.gnu.org/gnu/gdbm/）下载 gdbm 自己编译出 libgdbm.a 静态库  

`./configure` 执行报错

```
checking host systemtype... Invalid configuration `x86_64-unknown-linux-gnu': machine`x86_64-unknown' not recognized
```

`make`  
`make install`  

 

`cp /usr/share/libtool/config/config.guess .`  
`cp /usr/share/libtool/config/config.sub .`  



## 参数检查不通过

修改 `/usr/local/src/QConf/agent/qconf_dump.cc:63` 增加（char *）进行强转  
