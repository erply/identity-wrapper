# Identity Wrapper
Identity Wrapper is a GO component to provide ERPLY Identity Service endpoints for other GO projects.

For example: 
* login with credentials
* launch apps
* get JWT
* revoke JWT 
* verify JWT.
* etc. 

### Requirements
* go 1.12+

## Functions

### Import package

```Go
import (
	"gitlab.com/erply/identity-wrapper/Identity"
)
```

### Setup and init Identity API

##### Production Env
* __host__ is `https://id-api.erply.com/`
* __apiWorkers__ - Default value `1`
* __maxIdleConnections__ - Default value `1`
* __maxConnections__ - Default value `1`

##### Sandbox Env  
* __host__ is `https://id-api-sb.erply.com/`
* __apiWorkers__ - Default value `1`
* __maxIdleConnections__ - Default value `1`
* __maxConnections__ - Default value `1`

```Go

identityAPI := Identity.SetupAPI(host, apiWorkers, maxIdleConnections, maxConnections)

identityAPI.Init()
```

*Note: `apiWorkers, maxIdleConnections, maxConnections` are variadic input parameters. 
SetupAPI function can be called like:
```Go
//Option 1

identityAPI := Identity.SetupAPI(host)	

//Option 2

identityAPI := Identity.SetupAPI(host, apiWorkers)	

//Option 3

identityAPI := Identity.SetupAPI(host, apiWorkers, maxIdleConnections)

```

### LoginWithCredentials()
Login to Identity Launchpad with email and password to get JWT. Launchpad JWT is in
limited permissions. Check permissions from here https://jwt.io

```Go

login, err := identityAPI.LoginWithCredentials("john.toe@example.com", "ExamplePass12")

jwt := login.Result.JWT
companyID := login.Result.DefaultCompanyId

// For other companies IDs you will get if you use func GetUserConfirmedCompanies()

```

### GetApplications()
Get list of all applications.
```Go

apps, err := identityAPI.GetApplications(jwt)
	
```

### GetAccountAccess()
Get list of applications IDs where user account has access. Also, returns company based 
AccountID to launch apps.
```Go

access, err := identityAPI.GetAccountAccess(jwt, companyID)
accountID := access.Result.AccountID
	
```

### GetUserConfirmedCompanies()
Get list of companies where user has access.
```Go

companies, err := identityAPI.GetUserConfirmedCompanies(jwt)
	
```

### LaunchApp()
Use JWT and and your selected AccountID to launch app and get launchCode.
```Go

launch, err := identityAPI.LaunchApp(jwt, accountID)
launchCode := launch.Result.LaunchCode
	
```

### GetJWT()
Get JWT by launchCode. LunchCode is a hash which expires after 30 sec.
Returns JWT with all permissions you have. JWT lives in 24 hours.
```Go

getJWT, err := identityAPI.GetJWT(launchCode)
appJWT := getJWT.Result.JWT
	
```

### GetSessionID()
* GetSessionID by JWT. It returns short session token for Builder applications and 
Service Engine endpoints.
* Use this token in headers.
    * `JSESSIONID` - if using Builder apps.
    * `API_KEY` - if using Service Engine endpoints.

```Go
getSession, _ := identityAPI.GetSessionID(appJWT)
sessionID := getSession.Result.Session
```


### RevokeJWT()
It revokes persistence token which makes token unusable. Persistence token is 
inside of JWT.
```Go

revoke, err := identityAPI.RevokeJWT(jwt)
	
```

### VerifyJWT()
Verify persistence token by sending JWT. Returns (boolean) TRUE if it's valid and 
FALSE if expired or not exist.
```Go

verify, err := identityAPI.VerifyJWT(jwt)
	
```

## Use public key to verify JWT
##### Production Env
`https://apps.erply.com/jwt/pubkey.pem`

##### Sandbox Env 
`https://apps-sb.erply.com/jwt/pubkey.pem`
    
### Author
* Siim Roostalu - siim.roostalu@erply.com

### Changelog

##### 1.1.0
* __[IS-19]__ New endpoint getSessionID. 

##### 1.0.2
* security update: better config for net/http

##### 1.0.1
* more pointers, optional parameters for SetupAPI func

##### 1.0.0
* __[IS-19]__ Updated documentation and ready to use.

##### 0.1.0
* __[IS-19]__ Initial implementation.