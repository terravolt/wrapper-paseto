# wrapper-paseto
Wrapper Platform-Agnostic Security Tokens

```go
tokenMaker, err := token.NewPasetoMaker("example")
if err != nil {
return nil, fmt.Errorf("cannot create token maker: %w", err)
}
```

```go
payload, err := token.tokenMaker.VerifyToken("example")
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}
```