## How Oauth 2.0 improves app security ?

#### June 18th 2022

1. Was brought in 2 solve the problem of 3rd party authorisation ->

   -> eg. login via Google / Facebook

2. Should Oauth2.0 used even when you do not have 3rd party ? 

    -> Yes ! 
    -> the request to the api will look the same for all. How do you make sure the request came from your app at all ?

    -> Attacker might take a dump of usr/pwd and try hit your apis directly
        
    -> how do you whether it came from an actual user or not ?

    -> Then you will do Multifactor auth to login flow -> email/sms etc..

    -> now you have to go to each app dev team to add support for this

HOW DOES OAUTH SOLVE THIS ???

    -> what oauth does at high level is require that every application
        sends the user out to the OAUTH server to login there and then 
        redirect them back to the app, so that app can get tokens

    -> the key : USER actually leaves the app & go type their PWD at oauth server
        instead of ever giving their PWD to app.

    -> so app never sees user's pwd. 
        u dont have to worry about any unreliable 3rd party apps.

    -> also for MFA -> u just enable it on oauth server -> and get it enabled across
        
ANALOGY TO HOTEL MANAGEMENT :

-> your documents are verfied by receptionist (~ Oauth server)

-> they give you Access Card to access facilities at hotel -> particular room / pool / sauna / maybe not to play area

-> now the smart locks across the hotel -> need not know who you are -> it only needs to know what all can you access

-> the card store only this -> what all can be accessed + till what expiry

