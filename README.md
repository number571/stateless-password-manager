# Stateless Password Manager

> A simple and reliable password manager that generates passwords based on a master-key

## Dependencies

1. Go library [github.com/number571/go-rfc1751](https://github.com/number571/go-rfc1751)
2. Go library [golang.org/x/crypto](https://golang.org/x/crypto)

## How it works

The password manager uses `RFC1751` (to generate master keys) and a `Scrypt` (to generate passwords using the master key). Due to the determinism of the scrypt, the result (password) is always the same. This feature allows you to get rid of the need for synchronization, and as a result, both servers and any files.

### Used scrypt parameters

```
N=1048576 (iterations count)
r=8 (block size)
p=1 (parallelism factor)
```

## Installation

```bash
go install github.com/number571/stateless-password-manager/cmd/spm
```

## Generate master-key

```bash
spm -m 128
ANN GORY LOP WANT ELAN AUTO TEAR BUNT LEN OHIO TURF OVA
```

## Generate password

> [!CAUTION]
> Do not use the specified master key!

```bash
spm -t "google.com"
Master-Key: ANN GORY LOP WANT ELAN AUTO TEAR BUNT LEN OHIO TURF OVA
Please wait a few seconds...
Password: df626b36ae2f1f8d74525f9b7c620a2f82272df99da8c576cd1955bad6128ac6
```
