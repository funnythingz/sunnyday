# Sunnyday

Sunnyday is a Truncate implemented in Go.

## Installation

```sh
% go get github.com/funnythingz/sunnyday
```

## Usage

```go
package main

import(
  "fmt"
  "github.com/funnythingz/sunnyday"
)

func main() {
    fmt.Println(sunnyday.Truncate("うんこっこたのしいお！", 5))
}
```

```sh
うんこっこ...
```

### options

custom omission

```go
fmt.Println(sunnyday.Truncate("うんこっこたのしいお！", 5, "...more"))
```

```sh
うんこっこ...more
```

&copy; [funnythingz](http://hiropo.co.uk/)
