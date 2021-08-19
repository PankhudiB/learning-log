If you have a remote host and you need to send certain files to it smartly after some interval.

Example you might want to have this behaviour with certificates
Today I used it while provisioning gocd on a VM

    Rsync

* feels like smart / polished way of coping files
* transfers only the differences between source & destination.
* replacement of cp, scp or sftp commands
* Usage explained here : https://linuxize.com/post/how-to-use-rsync-for-local-and-remote-data-transfer-and-synchronization/

