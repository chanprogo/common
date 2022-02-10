## Install kubectl binary with curl on Linux

Download the latest release with the command:   
`curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"`   

Install kubectl  
`sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl`   

 
`kubectl version --client`   
`kubectl version`    
 