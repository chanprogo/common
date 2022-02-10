 
`kubectl get po -A` 


`kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4`   
 

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


Create a Service   
By default, the Pod is only accessible by its internal IP address within the Kubernetes cluster.    
To make the hello-node Container accessible from outside the Kubernetes virtual network, you have to expose the Pod as a Kubernetes Service.  

Expose the Pod to the public internet using the kubectl expose command:  
`kubectl expose deployment hello-node --type=LoadBalancer --port=8080`   

The `--type=LoadBalancer` flag indicates that you want to expose your Service outside of the cluster.   

The application code inside the image k8s.gcr.io/echoserver only listens on TCP port 8080.   
If you used kubectl expose to expose a different port, clients could not connect to that other port.  

View the Service you created:  
`kubectl get services`   

On cloud providers that support load balancers, an external IP address would be provisioned to access the Service.   

On minikube, the LoadBalancer type makes the Service accessible through the minikube service command.   

Run the following command:  
`minikube service hello-node`   


Katacoda environment only: 
Click the plus sign, and then click Select port to view on Host 1.   


Katacoda environment only: Note the 5-digit port number displayed opposite to 8080 in services output. This port number is randomly generated and it can be different for you.   













