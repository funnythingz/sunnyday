# Sunnyday

Sunnyday is a Truncate implemented in Go.

## Installation

```
% go get github.com/funnythingz/sunnyday
```

## Usage

```
package main

import(
  "fmt"
  "github.com/funnythingz/sunnyday"
)

func main() {
    fmt.Println(sunnyday.Truncate("うんこっこたのしいお！", 5))
}
```

```
うんこっこ...
```

&copy; funnythingz
