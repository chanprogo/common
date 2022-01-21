
允许 2181 端口 通过 防火墙：  

`vi /etc/sysconfig/iptables`  

`-A INPUT -m state--state NEW -m tcp -p tcp --dport 2181 -j ACCEPT`  
