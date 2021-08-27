## August 3nd 2021

DNAT
        Public ip -> private ip
        Somebody hitting doesn’t know the destination -> they just think they are hitting public ip
        
        Ip in n/w 1 ——— Ip in n/w 2
        10.7.50.21 —8081				192.168.17.10:443
        —8082				192.168.17.10:80
        —8083				192.168.17.11:443
        —8084				192.168.17.12:80
        —8085				192.168.17.12:8080

        * Website hosted inside data center behind Firewall 
            and needs to be accessible to users over Internet

        * done generally for packets coming inside the n/w

Reverse in SNAT

        192.168.17.10		* any port		-
        192.168.17.11								-			x.y.z.w  ———————> ip in n/w 3 will think all the request came from x.y.z.w
                                                                                    [n/w 3 don’t know who the real source is]
        192.168.17.12								-			
        192.168.17.13								-
        
        So when the response comes back the SNATter knows who originally requested 
        * Usecase -> pods in k8s server will generally have dynamic IPs --> 
                                in order to regulate request to external systems

        * Client inside LAN and behind Firewall needs to browse Internet.
        
        * done generally for packets leaving the n/w