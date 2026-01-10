# Shrike-Go

A basic library to program the FPGA of [Shrike Lite](https://github.com/vicharak-in/shrike) using TinyGo.

## Basic Usage

An example with the [sample LED blinky firmware](https://github.com/vicharak-in/shrike/blob/main/test/bitstreams/v1_4/led_blink.bin) is given in the [example](./example/) folder.

```go
import 	"github.com/yajtpg/shrike_go"

func main() {
    // ...
    shrike_go.Debug = true // Enables verbose logging
    err := shrike_go.Flash(<data in []byte>)
    if err != nil {
        panic(err)
    }
    // ...
}
```
