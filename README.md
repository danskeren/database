### Usage

```go
package main

import (
  "fmt"

  "github.com/danskeren/database/kv"
  "github.com/dgraph-io/badger"
  "github.com/dgraph-io/badger/options"
)

func main() {
  opts := badger.DefaultOptions
  opts.Dir = "./badger.db"
  opts.ValueDir = "./badger.db"
  opts.ValueLogLoadingMode = options.FileIO
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