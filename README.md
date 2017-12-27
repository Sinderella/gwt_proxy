# GWT Proxy

modified from https://github.com/elazarl/goproxy/tree/master/examples/goproxy-customca

intended to use with burp suite as part of pentesting.

generate keys from burp

1. Proxy>>Options>>Export CA Cert
2. choose "certificate and private key in PKS#12 keystore"
3. enter password
4. `openssl pkcs12 -in keyStore.pfx -out keyStore.pem -nodes`
5. enter the password above
6. put them in `cert.go`