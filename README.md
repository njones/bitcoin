# Bitcoin

## Library for manipulating bitcoin addresses

### Usage

#### Generate a btc private/public key pair

```go
import btcpk "github.com/steakknife/bitcoin/private_key"

pk := btcpk.MustGenerate()

fmt.Printf("BTC Private key address (wallet format): %s", pk)
fmt.Printf("BTC                  Public key address: %s", pk.PublicKey())
```

#### Parse a btc public key address
```go
import btcpub "github.com/steakknife/bitcoin/public_key"

pubKey := btcpub.MustDecode("19Bq1gipWrLxFGqVH41Un2suWnGzWxNjbZ")

fmt.Printf("BTC public key: %s\n", pubKey)
fmt.Printf("  network: %s (%d, 0x%02X)\n", pubKey.Network, pubKey.Network, pubKey.Network)
fmt.Printf("  address: %v\n", pubKey.Address)
```

#### Parse a btc private key address in wallet format
```go
import btcpk "github.com/steakknife/bitcoin/private_key"

privKey := btcpk.MustDecode("5Ju6hf57BPdusMDUg4C6gPKiauXSahHVnTGmTNHoJeGUwJHeqSY")
```


## Testing

### Run test cases

`make`

### Run test cases and benchmarks

`make GOTESTOPT='-bench=.'`


## License

MIT

## Copyright

Copyright (C) 2014,2015 Barry Allard
