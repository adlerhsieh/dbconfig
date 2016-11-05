# DBConfig

Read database config file in a Rails project.

### Installation

```
go get github.com/adlerhsieh/dbconfig
```

### Example

```go
package main

import (
  "fmt"

  "github.com/adlerhsieh/dbconfig"
)

func main() {
  // Specify filepath & enviroment in parameters
  config := dbconfig.ReadFile("/path/to/database.yml", "development")

  fmt.Println(config["database"])
}
```

### Development

Run test before push:

```
go test dbconfig_test.go
```

or with Godo server:

```
godo -w
```
