*Simple template for generate API Solidity Smart-Contract*

**Features:**

- gRPC
- HTTP/HTTPS
- Auto-generate doc for API
- Auth with Bearer (Temporary Access Token)

<br>
    
**Build**

1) Install Go 1.9.2 `golang.org`
2) Install Dep `github.com/golang/dep`
3) run `dep ensure` in root project directory
4) `go build main.go`

<br>

Note: Always include ```Authorization: Bearer <your token>``` in Header.

PS: For ```/v1/auth/*``` routes, you can use random token for get a **real** token by one of the ways: <br>
    1) Json KeyStore + Password. <br>
    2) Private Key.
    
