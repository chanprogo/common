"error while loading shared libraries: xxx.so.x" 错误的原因和解决办法  

一般 我们 在 Linux 下 执行 某些 外部程序 的 时候 可能 会提示 找不到 共享库 的错误, 比如:  

tmux: error while loading shared libraries: libevent-1.4.so.2: cannot open shared object file: No such file or directory  


原因 一般 有两个:  
1. 一个 是操作系统里 确实没有 包含 该共享库(lib*.so.* 文件) 或者 共享库 版本 不对, 遇到这种情况 那就去 网上下载 并安装上即可.  
2. 另外 一个 原因 就是 已经 安装了 该共享库, 但 执行 需要 调用 该共享库 的程序的时候, 程序 按照 默认 共享库 路径 找不到 该共享库文件.  








所以 安装共享库后 要注意 共享库路径 设置问题, 如下:

### 如果共享库文件 安装到 了 /lib 或 /usr/lib 目录下, 那么需执行一下 ldconfig 命令

`ldconfig` 命令 的用途, 主要是 在  
默认搜寻目录(`/lib` 和 `/usr/lib`)  
以及 动态库 配置文件 `/etc/ld.so.conf` 内 所列的目录 下,   
搜索出 可共享的 动态 链接库(格式如 lib*.so*),   
进而 创建出 动态 装入程序(ld.so) 所需的 连接 和 缓存文件.   

缓存文件 默认为 `/etc/ld.so.cache`, 此文件 保存 已排好序的 动态 链接库 名字 列表.  













### 如果 共享库 文件 安装到了 /usr/local/lib (很多开源的 共享库 都会安装到 该目录下) 或 其它 "非 /lib 或 /usr/lib" 目录下

那么 在执行 ldconfig 命令前,   
还要把 新共享库 目录 加入到 共享库 配置文件 `/etc/ld.so.conf` 中, 如下:

`cat /etc/ld.so.conf`  其内容应该是 include ld.so.conf.d/*.conf
`echo "/usr/local/lib" >> /etc/ld.so.conf`  
`ldconfig`  








### 如果 共享库 文件 安装到了 其它 "非 /lib 或 /usr/lib" 目录下,  但是 又不想在 /etc/ld.so.conf 中加路径(或者是没有权限加路径).   

那可以 export 一个全局变量 `LD_LIBRARY_PATH`, 然后 运行程序的时候 就会去 这个目录中 找共享库.  

`LD_LIBRARY_PATH` 的 意思 是告诉 loader 在哪些目录中 可以找到 共享库.   
可以设置 多个 搜索目录, 这些目录之间 用冒号 分隔开.   
比如安装了一个 mysql 到 /usr/local/mysql 目录下, 其中 有一大堆库文件在 /usr/local/mysql/lib 下面,   
则可以在 `.bashrc` 或 `.bash_profile` 或 shell 里 加入 以下语句 即可:  

`export LD_LIBRARY_PATH=/usr/local/mysql/lib:$LD_LIBRARY_PATH`   

一般来讲 这只是 一种临时的解决方案, 在没有权限或临时需要的时候使用.