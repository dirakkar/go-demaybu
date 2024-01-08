module github.com/dirakkar/go-demaybu/dapp/crypto

go 1.21.1

replace github.com/dirakkar/go-demaybu/dapp => ../../dapp

require (
	github.com/dirakkar/go-demaybu/dapp v0.0.0-00010101000000-000000000000
	github.com/zeebo/blake3 v0.2.3
)

require github.com/klauspost/cpuid/v2 v2.0.12 // indirect
