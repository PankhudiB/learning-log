## August 26th 2021

I came across a tool called Sops that can be used to handle workflow around managing secrets .

Makes the process soo smooth for a developer in a way :

SOPS -> workflow around (any encryption tool)
    
    - open file -> modify -> close file
      [decryption and modification and encryption of modified is taken care]

I personally liked the fact that it can do conflict resolution for me which is not possible with the existing tool -> Vault kdbx files + keepassXC I am using.

Where if there are changes in master then I have to make sure to be on latest commit and then check-in. Ugghh ! 

But again each tool has its own advantage -> Like the vault + keepassxc gells really well as browser extension where it can just pick appropriate secrets.

But yeah coming back to sops
This section gives perfect elaboration on what are the operational requirements for using such tool :

    https://github.com/mozilla/sops#operational-requirements

While encrypting -> u need Fingerprint (their term) 

I ended up using gpg, so in gpg
Fingerprint is analogous to -> public key

You can even rotate the encryption

Amazinggggg part ->
To export our secrets as environment variables to our process . eg. bash ->
    
    sops exec-env sample.yaml ‘bash’
which will essentially not even leave trace in the history of the terminal session you are running sops on

Github repo : https://github.com/mozilla/sops

    Use-case where we used it -> Secret.tfvars for terraform



-------------------

