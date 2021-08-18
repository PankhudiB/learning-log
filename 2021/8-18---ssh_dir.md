
    cd ~/.ssh/
    -> ur public and private key
    -> known_hosts

    SSH keeps a list with an identifier (known as a host key) for each server to which you connect. 
    This is the known_hosts file.

* authorized_key -> exists on the server. It says “Anybody with the public key matching *this* is allowed access as this user (owning the file) without a password”.
* known_hosts -> on ssh client side. 
  * Used to verify that the host key offered by the server has not changed.
    * If server host key changed : 
      * Indication of man-in-the-middle attack
      * or that the host has been reinstalled
      