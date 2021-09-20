
Centos 6 yum 源 centos/6.6/os/x86_64/repodata/repomd.xml: [Errno 14] PYCURL ERRO 报错解决方法 及 新的源   



`cd /etc/yum.repos.d/`  
`ll`  

replace CentOS-Base.repo  

`yum clean all`  
`yum makecache`  
`yum repolist`   

