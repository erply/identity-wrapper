# Identity Wrapper
Identity Wrapper is a GO component to provide Identity endpoints for other services.
For example to get, revoke and verify JWT. Also possible to login with credentials etc.

### Requirements
* go 1.12+

## Functions

### Import package

```Go
import (
	"git.nimi24.com/erply-api/identity-wrapper/Identity"
)
```

### Setup and init Identity API
```Go
identityAPI := Identity.SetupAPI("http://localhost/identity/api/", 1, 1, 1)
identityAPI.Init()
```

### LoginWithCredentials()
Login to Identity Launchpad with email and password to get JWT. 

```Go

login, err := identityAPI.LoginWithCredentials("john.toe@example.com", "Test12345")

jwt := login.Result.JWT
companyID := login.Result.DefaultCompanyId

// If default company isn't setted. Then you have to do GetUserConfirmedCompanies() 
// request.

```

### GetApplications()
Get list of all applications.
```Go

apps, err := identityAPI.GetApplications(jwt)
	
```

### GetAccountAccess()
Get application IDs to see where account has access. Also, returns company based 
AccountID to launch app.
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
Use JWT and AccountID to launch app and get launchCode.
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