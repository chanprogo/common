
Minikube is a lightweight Kubernetes implementation 
that creates a VM on your local machine 
and deploys a simple cluster containing only one node.   

Minikube is available for Linux, macOS, and Windows systems.   

The Minikube CLI provides basic bootstrapping operations for working with your cluster, including start, stop, status, and delete.



minikube is local Kubernetes,   
focusing on making it easy to learn and develop for Kubernetes.   

All you need is Docker (or similarly compatible) container   
or a Virtual Machine environment,    
and Kubernetes is a single command away: `minikube start`     




Hello Minikube   
This tutorial shows you how to run a sample app on Kubernetes using minikube and Katacoda. Katacoda provides a free, in-browser Kubernetes environment.   
Note: You can also follow this tutorial if you've installed minikube locally. See minikube start for installation instructions.    

Objectives   
1. Deploy a sample application to minikube.   
2. Run the app.   
3. View application logs.   

  

The dashboard command enables the dashboard add-on and opens the proxy in the default web browser.  
You can create Kubernetes resources on the dashboard such as Deployment and Service.  

If you are running in an environment as root, see Open Dashboard with URL.   
Open Dashboard with URL   
If you don't want to open a web browser, run the dashboard command with the --url flag to emit a URL:   
`minikube dashboard --url`   

By default, the dashboard is only accessible from within the internal Kubernetes virtual network.   
The dashboard command creates a temporary proxy to make the dashboard accessible from outside the Kubernetes virtual network.    




To stop the proxy, run Ctrl+C to exit the process. After the command exits, the dashboard remains running in the Kubernetes cluster.  
You can run the dashboard command again to create another proxy to access the dashboard.  

Create a Deployment  
A Kubernetes Pod is a group of one or more Containers, tied together for the purposes of administration and networking.   
The Pod in this tutorial has only one Container.  
A Kubernetes Deployment checks on the health of your Pod and restarts the Pod's Container if it terminates.   

Deployments are the recommended way to manage the creation and scaling of Pods.   

Use the `kubectl create` command to create a Deployment that manages a Pod.   
The Pod runs a Container based on the provided Docker image.   

`kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4`   

View the Deployment:  
`kubectl get deployments`    

View the Pod:  
`kubectl get pods`    

View cluster events:  
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













