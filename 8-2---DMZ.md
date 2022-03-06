## August 2nd 2021

DMZ network 

    DMZ - Demilitarized zone

    The name comes from military world - where you usually have buffer zones / strip of Land between countries or during a WAR
    So DMZ is subnet that separates a local area network (LAN) from other untrusted networks -- usually, the public internet.
    And you need this zone because any internal IP should not be exposed in the first landing zone of a Request
    
    A demiliterized zone usually has  
        - firewall setup in it -> to avoid DOS and DDOS attack
        - Does DNATing [Destination Address Translation] from Public IP --> Private IP
        - It should also have some intrusion detection and intrusion prevention system

    