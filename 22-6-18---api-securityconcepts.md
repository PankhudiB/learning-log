## Api security concepts

#### June 18th 2022

###Roles

1. User <---------------------> Resource Owner
2. The device (mbl/browser) <-----> User Agent
3. The application <--------> Oauth Client
4. the APIs <----------------> Resource Server
5. Auth Server

User -> uses Device -> to run Application --> Auth Server --> to acsess Resource Server

-----------------------

###Application / Client Types

1. Confidential 
2. Public

Diff between the 2 -> whether they can be deployed with some sort of credentials

Anytime the application is running code on a device that user controls --> u can't include the credentials
-> it could be any smart tv, mobile, 

-> server side java app for publishing blog posts -> the application is considered a "confidential client" since it can be deployed with credentials

->  

-----------------------

Each application/oauth client has a ClientID

-------------

Front Channel : Using the browser’s address bar to move data between two other pieces of software is using the front channel.

Back Channel : Any HTTP client that makes a request to an HTTP server is using the back channel, even if that client is JavaScript code in a browser

----------------

When a user is using Travis CI to test code from GitHub, GitHub is playing which of the OAuth roles?
-> Auth Server & Resource Server

-------------