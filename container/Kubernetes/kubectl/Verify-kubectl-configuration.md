
In order for kubectl to find and access a Kubernetes cluster, 
it needs a kubeconfig file, 
which is created automatically 
when you create a cluster using kube-up.sh 
or successfully deploy a Minikube cluster.   

By default, kubectl configuration is located at `~/.kube/config`.    

Check that kubectl is properly configured by getting the cluster state:   
`kubectl cluster-info`   

If kubectl cluster-info returns the url response 
but you can't access your cluster, to check whether it is configured properly, use:  
`kubectl cluster-info dump`   


Cluster details
Let’s view the cluster details.   
We’ll do that by running `kubectl cluster-info`   

During this tutorial, we’ll be focusing on the command line for deploying and exploring our application.   
To view the nodes in the cluster, run the `kubectl get nodes` command.   

This command shows all nodes that can be used to host our applications.   
Now we have only one node, and we can see that its status is ready (it is ready to accept applications for deployment).   
   

