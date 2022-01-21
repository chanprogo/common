
A Kubernetes Pod  is    
a group of one or more Containers, 
tied together    
for the purposes of administration and networking.   


A Kubernetes Deployment 
checks on      
the health of your Pod 
and 
restarts the Pod's Container if it terminates.  


Deployments are the recommended way 
to manage      
the creation and scaling of Pods.     



By default,  
the Pod    is     only accessible by 
its internal IP address 
within the Kubernetes cluster.    


To make 
the hello-node Container 
accessible from outside the Kubernetes virtual network,   
you have to     
expose the Pod 
as a Kubernetes Service.  
