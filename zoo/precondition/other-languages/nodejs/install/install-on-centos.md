
`cd /usr/local/bin/`  

`wget https://npm.taobao.org/mirrors/node/v12.16.1/node-v12.16.1-linux-x64.tar.gz`  

`tar -xvf node-v12.16.1-linux-x64.tar.gz`  

`cd node-v12.16.1-linux-x64`  
`yum install gcc gcc-c++`  

`cd .. && mv node-v12.16.1-linux-x64 Node.js`  

`ln -s /usr/local/bin/Node.js/bin/node /usr/bin/node`  
`ln -s /usr/local/bin/Node.js/bin/npm /usr/bin/npm`  
`ln -s /usr/local/bin/Node.js/bin/npx /usr/bin/npx`  
