# cstlog

A simple log handler for [apex/log](https://github.com/apex/log) that formats timestamps in **Central Standard Time (CST/CDT)** for human-readable CLI output.

## Features

- Replaces the default `cli` handler from `apex/log`
- Prints log timestamps in `America/Chicago` time zone
- Automatically shows `CST` or `CDT` depending on daylight saving time
- Small, dependency-free (except `apex/log`)

## Installation

The tricky part is that `tzdata` necessary to find the `America/Chicago` zone is not in most minimal production containers. So you will need to add this line to your Dockerfile for Alpine builds.

```docker
RUN apk add --no-cache tzdata
```

## Usage
```go
package main

import (
    "os"
    "github.com/apex/log"
    "github.com/zacsketches/cstlog"
)

func init() {
    log.SetHandler(cstlog.New(os.Stdout))
}

func main() {
    log.Info("server started")
}
```

## Example Output

`2025-05-05 13:30:02 CDT INFO server started`

## License

MIT © zacsketches 2025
