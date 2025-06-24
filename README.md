# stateless-password-manager

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
spm --new-mk 128
ANN GORY LOP WANT ELAN AUTO TEAR BUNT LEN OHIO TURF OVA
```

## Generate password

> [!DANGER]
> Do not use the specified master key!

```bash
spm --target "google.com"
Master-Key: GORY LOP WANT ELAN AUTO TEAR BUNT LEN OHIO TURF OVA
Please wait a few seconds...
Password: ad5d87cb4b4ad0753b695f0375f07f0010d0a8061991cc39a92b13ed34eec4af
```
