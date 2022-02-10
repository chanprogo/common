 
`kubectl get po -A` 


`kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4`   
`kubectl expose deployment hello-node --type=LoadBalancer --port=8080`  
The `--type=LoadBalancer` flag indicates that you want to expose your Service outside of the cluster.  

Create a sample deployment and expose it on port 8080:  
`kubectl create deployment hello-minikube --image=k8s.gcr.io/echoserver:1.4`  
`kubectl expose deployment hello-minikube --type=NodePort --port=8080`  
`kubectl get services hello-minikube` 

Alternatively, use kubectl to forward the port:   
`kubectl port-forward service/hello-minikube 7080:8080`  



`kubectl get deployments`    

`kubectl get pods`    
 
`kubectl get events`   

View the kubectl configuration:  
`kubectl config view`  



`kubectl get services`   
 