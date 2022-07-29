
`yum install python3-pip`  













## 方法一

`wget https://bootstrap.pypa.io/get-pip.py --no-check-certificate`  
`python get-pip.py`  




## 方法二

`wget https://pypi.python.org/packages/11/b6/abcb525026a4be042b486df43905d6893fb04f05aac21c32c638e939e447/pip-9.0.1.tar.gz#md5=35f01da33009719497f01a4ba69d63c9`  
`tar -xzvf pip-1.5.4.tar.gz`  
`cd pip-1.5.4`  
`python setup.py install`  

先下载安装 setuptools 包  
`wget http://pypi.python.org/packages/source/s/setuptools/setuptools-2.0.tar.gz`  
`tar zxvf setuptools-2.0.tar.gz`  
`cd setuptools-2.0`  
`python setup.py build`  
`python setup.py install`  

现在，setuptools 已经安装好，我们再次进入 pip-1.5.4 目录，使用 `python setup.py install` 命令安装 pip

若安装后 pip 命令 无法使用，搜索系统中 pip 文件，创建命令链接  
`ln -s /usr/local/python27/bin/pip  /usr/bin/pip`   