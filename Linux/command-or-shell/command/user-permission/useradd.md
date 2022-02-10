
`useradd -m [用户帐号]`   
`usermod -s /bin/bash [用户帐号]`  
`passwd [用户帐号]`  
`chown -R myusername[:mygroup] /home/myusername`   






没有权限进行 sudo （`is not in the sudoers file`），解决方法：  
`su`  
`vim /etc/sudoers`  

修改文件内容：找到 `root    ALL=(ALL)       ALL` 一行，  
在下面插入新的一行，内容是 `username   ALL=(ALL)       ALL`，   
`:wq!`

`exit`  
