# Wrapper Platform-Agnostic Security Tokens

The following Go code creates a new instance of the "tokenMaker" using the "token.NewPasetoMaker" function and passing in the string argument "example".
```go
tokenMaker, err := token.NewPasetoMaker("example")
    if err != nil {
		//..
		return
	}
```
The next code uses the "tokenMaker" instance to create an "accessToken" and an "accessPayload" using the "tokenMaker.CreateToken" function, passing in "userID" and "AccessTokenDuration" as arguments.
```go
accessToken, accessPayload, err := server.tokenMaker.CreateToken("userID", "AccessTokenDuration")
	if err != nil {
		//..
		return
	}
```
Finally, the code uses the "tokenMaker" instance to verify a token by calling the "tokenMaker.VerifyToken" function and passing in "example" as the argument.
```go
payload, err := tokenMaker.VerifyToken("example")
	if err != nil {
		//..
		return
	}
```