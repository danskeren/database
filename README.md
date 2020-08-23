This package is for personal use. I recommend against using it as I will make breaking changes.

### KV Usage

```go
package main

import (
  "fmt"

  "github.com/danskeren/database/kv"
  "github.com/dgraph-io/badger/v2"
  "github.com/dgraph-io/badger/v2/options"
)

func main() {
  opts := badger.DefaultOptions("./badger.db")
  opts.ValueLogLoadingMode = options.FileIO
  opts.TableLoadingMode = options.FileIO
  badgerDB, err := kv.Open(opts)
  if err != nil {
    // handle err
  }
  
  if err = badgerDB.Set([]byte("mykey"), []byte("myvalue")); err != nil {
    // handle err
  }
  
  val, err := badgerDB.Get([]byte("mykey"))
  if err != nil {
    // handle err
  }
  fmt.Println("value:", string(val))
}
```