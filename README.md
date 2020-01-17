# jwt-issuer

A bare minimum service that issues JWTs with a JWKS endpoint.

This is meant for simple testing/demo, not production, so everything is as simple
and self-contained as possible with as little fluff as possible.

Generates a fresh ECDSA key pair on every execution.  

## Endpoints

### /generate

Generates a new token.  A private claim of "UserID" is added.

```bash
# Generates a new token
curl -s localhost:8080/generate

# Generates a new token and stores the token in TOKEN
export TOKEN = $(curl -s localhost:8080/generate)

# Generates a new token with the username myfakeid
curl -s -H "X-User-ID: myfakeid" localhost:8080/generate
```

### /check

Checks a token given in the Authorization header as a bearer token.  Returns 200
if valid and writes the user's name as the response, returns 401 otherwise.

### /jwks

Returns the JSON jwks object for the generated public key.

