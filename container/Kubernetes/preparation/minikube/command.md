
Start your cluster:  
`minikube start`  

Interact with your cluster  
If you already have kubectl installed, you can now use it to access your shiny new cluster:  
`kubectl get po -A`  


Initially, some services such as the storage-provisioner, 
may not yet be in a Running state.   
This is a normal condition during cluster bring-up, and will resolve itself momentarily.   
For additional insight into your cluster state,   
minikube bundles the Kubernetes Dashboard,   
allowing you to get easily acclimated to your new environment:   

`minikube dashboard`  
