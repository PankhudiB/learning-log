## Authorization code flow with web application

#### June 19th 2022

Code Verifier (Secret) ==> random string -> 43-128 chars long

-----
Code Challenge (Public Hash) ==> base64url(SHA256(Code_Verifier))

----

auth-endpoint (from docs || from ouath server's metadata url)

==>  https://authorization-server.com/auth?

with some query params :

```https://authorization-server.com/auth?```

```response_type=code&```    -------> to tell the oauth server u r doing auth code flow

```client_id=CLIENT_ID&``` -------->  to tell the oauth server which app is making the request

```redirect_uri=REDIRECT_URI``` ----> it has to match the 1 of the redirect_urls you added during registration of app

```scope=photos``` --> scope

```state=XXXX``` ---> initially used for CSRF , but PKCE provides it.. so u can use it to store appln specific state
                eg. which page to redirect to after they login
                (Only safe to use with PKCE)

```code_challenge=XXXX``` --> hash of the code_verifier -> base64url(SHA256(Code_Verifier))

```code_challenge_method=S256``` ---> hash method you used SHA256

The app sends this to server for user to login.
Your app wont see the user again until they come back to the redirect URL with auth code.

Now user goes and logs in  -> approves the request

----------------

Now the auth server generates the 1-time auth code and sends them back to you app.

Or sends them back error-code

the auth server will give :

```https://example-app.com/redirect?```

```code=AUTH_CODE&```

```state=XXXX```

OR 

```https://example-app.com/redirect?```

```error=ACCESS_DENIED&```

```state=XXXX```

we should check the state value -> with -> the 1 used in the request. Thats for CSRF protection

----------------

Now you are ready to exchange the auth code with ```access token```

Now u'll make a backchannel HTTPS request from application server to oauth server's ```token endpoint```

```POST https://authorization-server.com/token?̌```
```grant_type=authorization_code&``` -> this tell server that u r doing auth code grant
```code=AUTH_CODE``` -> 1 got from prev request
```redirect_uri=REDIRECT_URI``` -> used earlier
```code_verifier=VERIFIER_STRING``` ->if u r using PKCE , the plain text string that your app made up in the beginning
```client_id=CLIENT_ID```
```client_secret=CLIENT_SECRET```
-------------
The AUTH server will respond back with :
{
    "token_type" : "Bearer",
    "access_token" : "BLAHBLAH",
    "expires_in": 3600,
    "scope" : photos,
    "refresh_token": "BLAAAAH"
}

Now your app can use this access_token to make API calls

If you got a refresh token then on expiration of current access_token you can get new one without going through the whole flow.

you just make 

```POST https://authorization-server.com/token?̌```
```grant_type=refresh_token&``` -> changed grant type
```refresh_token=BLAAAAH&``` -> the refresh_token u got initially
```client_id=CLIENT_ID```
```client_secret=CLIENT_SECRET``` --> so that oauth server knows it is this app making the request

<-- you will again get new access token and new refresh token

-------------------





-------

