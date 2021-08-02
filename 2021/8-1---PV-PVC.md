PV and PVC

    PersistentVolume subsystem provides an API for users and administrators 
        that abstracts details of how storage is provided & 
        from how it is consumed.

    new API resources: PersistentVolume and PersistentVolumeClaim

    PersistentVolume (PV) is a piece of storage in the cluster
        provided either by 
            -> statically by admin
            -> or dynamically by Storage Class

    It is a resource just like a node
    - But lifecycle independent of any individual Pod

    This PV API object captures implementation details of storage,
        -> be that NFS, iSCSI, or a cloud-provider-specific storage system.
    

    Pod                            |            PVC
                                   |            request for storage by user 
    uses node resources            |            use pv resources
    request for CPU and memory     |            request specific size and access modes


    Use Protection to prevent Data loss :
    If a user deletes a PVC used by a Pod -> PVC not removed immediately -> Postponed until PVC is no longer actively used by any Pods.
    If admin deletes a PV bound to PVC -> the PV not removed immediately -> Postponed until PV is no longer bound to a PVC.

    Reclaiming the resources ?
    Once PVC object is deleted -> it tells cluster that the PV used can be
        -> recycled
        -> ratained
        -> delete

    Volume mode 
        -> block    -> fastest way to access volume -> although pod need to know how to access raw block
        -> file system -> 
        
    The access modes are:
        ReadWriteOnce -- volume - read-write by a single node
        ReadOnlyMany -- volume - read-only by many nodes
        ReadWriteMany -- volume - read-write by many nodes

            How is it for vpshere ?
            Volume Plugin	ReadWriteOnce	ReadOnlyMany    ReadWriteMany
            VsphereVolume	✓	            -	            - (works when Pods are collocated)