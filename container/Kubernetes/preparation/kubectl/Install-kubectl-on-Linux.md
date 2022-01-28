## Install kubectl binary with curl on Linux

Download the latest release with the command:   
`curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"`   

Install kubectl  
`sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl`   

Test to ensure the version you installed is up-to-date:  
`kubectl version --client`   
To check if kubectl is installed you can run the `kubectl version` command.   
 
