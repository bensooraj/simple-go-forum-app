Generate private key (.key):
```sh
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
```

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
```sh
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

A little more detailed explanation here: [How to trust self-signed localhost certificates on Linux Chrome and Firefox
](https://stackoverflow.com/a/50788371)

List of commands to compile and install cURL with HTTP/2 support in Mac OS X using Homebrew:
```sh
# install cURL with nghttp2 support
➜  brew install curl --with-nghttp2

# link the formula to replace the system cURL
➜  brew link curl --force

# now reload the shell

# test an HTTP/2 request passing the --http2 flag
➜  curl -I --http2 https://www.cloudflare.com/
```

http2 CURL Command:
```sh
curl -X GET -I --http2 --insecure https://localhost:8080/hello/Ben
```