# Bitcoin

## Library for manipulating bitcoin addresses

### Usage

#### Generate a btc private/public key pair

```go
import btcpk "github.com/steakknife/bitcoin/private_key"

// Generate a random pub/priv key pair
pk := btcpk.MustGenerate()
fmt.Println("BTC Private key address (wallet format): %s", pk)
fmt.Println("BTC                  Public key address: %s", pk.PublicAddress())
```

#### Parse a btc public key address
```go
import btcpub "github.com/steakknife/bitcoin/public_key"

pubKey := btcpub.MustDecode("19Bq1gipWrLxFGqVH41Un2suWnGzWxNjbZ")
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
