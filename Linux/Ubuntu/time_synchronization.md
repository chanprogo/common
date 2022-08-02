Ubuntu 如何同步网络时间  
 

装完 Ubuntu 设置完时间，重启总是恢复设置前的时间。  

设定时区：`dpkg-reconfigure tzdata`  

           选择 Asia -> 再选择 Shanghai -> OK  

 

解决方法：  

1. 安装 ntpdate 工具  
`sudo apt-get install ntpdate`  

2. 将系统时间与网络同步  
`ntpdate cn.pool.ntp.org`  

3. 将时间写入硬件  
`hwclock --systohc`  

4. 查看时间是否已同步  
`date`   