Pivotal Container Service 
    
    K8s -> container orchestrator
    When u create k8s cluster 
        -> u create master nodes and worker nodes
        -> ETCD <-> k8s master <-> k8s worker
        -> master node has 
            -> scheduler
            -> distributes load
            -> Controller-Manager
            -> API server  
            -> etcd -> cluster state db
        -> worker node has 
            -> pods has appication container
            -> kubeproxy
            -> kubelet -> the component that runs on all of the machines in your cluster and 
                            does things like starting PODs and containers

    Challenges with k8s
        -> if pod goes down -> a component from master will bring it up
            but what if master goes down || worker goes down ?? 
            --> k8s does not handle this out of the box
            -> PKS handles this

        -> all pods in k8s cluster reside in single flat shared network address space 
                -> meaning each pod can talk to another pod directly with IP
                -> there is no NAT gateway -> security limitation

        -> k8s handles scaling of pods. 
            But what if u want to scale master || etcd
            -> PKS handles this

        -> Health check and healing of master node -> pks

        -> patch and upgrade of fleet of k8s cluster with 0 downtime ?

        -> Watch https://youtu.be/3PLq2vTIjVE
            -> k8s releases so frequently and supports only N-2 releases api
            -> Does your platform that scale upgrade and patch -> to be consistent with this timeline ?
            -> How many clusters will u need ? Does ur platform to setup these cluster fast ?
            -> Does ur platform do layer7 load balancing ?
            ---------> PKS solves this

        -> u ll need conatiner image repository
        -> monitor app and platform
        -> log management soln
        -> load balancing and routing soln
        -> Challenges u need to solve to run enterprise level k8s

