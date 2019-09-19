# Identity Wrapper
Identity Wrapper is a GO component which provides ERPLY Identity Service endpoints for other GO projects.

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
	"github.com/erply/identity-wrapper/Identity"
)
```

### How to setup?

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


<details><summary>Init Identity API</summary>

### Init Identity API

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

</details>

<details><summary>LoginWithCredentials()</summary>

### LoginWithCredentials()
Login to Identity Launchpad with email and password to get JWT. Launchpad JWT is in
limited permissions. Check permissions from here https://jwt.io

```Go

login, err := identityAPI.LoginWithCredentials("john.toe@example.com", "ExamplePass12")

jwt := login.Result.JWT
companyID := login.Result.DefaultCompanyId

// For other companies IDs you will get if you use func GetUserConfirmedCompanies()

```

</details>

<details><summary>GetApplications()</summary>

### GetApplications()
Get list of all applications.
```Go

apps, err := identityAPI.GetApplications(jwt)
	
```

</details>

<details><summary>GetAccountAccess()</summary>

### GetAccountAccess()
Get list of applications IDs where user account has access. Also, returns company based 
AccountID to launch apps.
```Go

access, err := identityAPI.GetAccountAccess(jwt, companyID)
accountID := access.Result.AccountID
	
```

</details>

<details><summary>GetUserConfirmedCompanies()</summary>

### GetUserConfirmedCompanies()
Get list of companies where user has access.
```Go

companies, err := identityAPI.GetUserConfirmedCompanies(jwt)
	
```

</details>

<details><summary>LaunchApp()</summary>

### LaunchApp()
Use JWT and and your selected AccountID to launch app and get launchCode.
```Go

launch, err := identityAPI.LaunchApp(jwt, accountID)
launchCode := launch.Result.LaunchCode
	
```

</details>

<details><summary>GetJWT()</summary>

### GetJWT()
Get JWT by launchCode. LunchCode is a hash which expires after 30 sec.
Returns JWT with all permissions you have. JWT lives in 24 hours.
```Go

getJWT, err := identityAPI.GetJWT(launchCode)
appJWT := getJWT.Result.JWT
	
```

</details>

<details><summary>GetSessionID()</summary>

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

</details>

<details><summary>RevokeJWT()</summary>

### RevokeJWT()
It revokes persistence token which makes token unusable. Persistence token is 
inside of JWT.
```Go

revoke, err := identityAPI.RevokeJWT(jwt)
	
```

</details>

<details><summary>VerifyJWT()</summary>

### VerifyJWT()
Verify persistence token by sending JWT. Returns (boolean) TRUE if it's valid and 
FALSE if expired or not exist.
```Go

verify, err := identityAPI.VerifyJWT(jwt)
	
```

</details>

## Use public key to verify JWT
##### Production Env
`https://apps.erply.com/jwt/pubkey.pem`

##### Sandbox Env 
`https://apps-sb.erply.com/jwt/pubkey.pem`
    
### Author
* Siim Roostalu - siim.roostalu@erply.com

### Changelog

<details><summary>1.1.1</summary>

##### 1.1.1
* Changed module to github.com

</details>

<details><summary>1.1.0</summary>

##### 1.1.0
* __[IS-19]__ New endpoint getSessionID.

</details> 

<details><summary>1.0.2</summary>

##### 1.0.2
* security update: better config for net/http

</details>

<details><summary>1.0.1</summary>

##### 1.0.1
* more pointers, optional parameters for SetupAPI func

</details>

<details><summary>1.0.0</summary>

##### 1.0.0
* __[IS-19]__ Updated documentation and ready to use.

</details>

<details><summary>0.1.0</summary>

##### 0.1.0
* __[IS-19]__ Initial implementation.

</details>