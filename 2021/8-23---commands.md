## August 23rd 2021

I came across shred command today in one of the script used to spin up a VM. It was used to delete the password file.

The interesting thing about this is it is not just called delete -> it is shred -> just as you shred paper into strips with those fancy machines -> you just cant get them back ! Ouch !

        Shred -> unix command

used to securely delete files and devices so that it is extremely difficult to recover them, even with specialized hardware and technology;


------------------------------
Today while building an OS base image for one of the VM, we logged in into a jump host (because the jump host had the access to nutanix which had RHEL hardened base image) 
Building an image would of course take a lot of time. More than the ssh session from my local to jump host  
So 
To run image building process in background via terminal , found a tool : Nohup

        Nohup -> https://www.geeksforgeeks.org/nohup-command-in-linux-with-examples/

------------------------------

        ps -ax

For identifying the programs that are running on the system and the resources they are using

------------------------------