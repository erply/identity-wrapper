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
Identity __Host__ is `https://id-api-sb.erply.com/`
* __apiWorkers__ - Default value `1`
* __maxIdleConnections__ - Default value `1`
* __maxConnections__ - Default value `1`

```Go

identityAPI := Identity.SetupAPI("http://localhost/identity/api/", 1, 1, 1)

identityAPI.Init()

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
    
### Author
* Siim Roostalu - siim.roostalu@erply.com

### Changelog

##### 0.1.0
* __[Feature]__ Initial implementation