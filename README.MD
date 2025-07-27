# go-shelly-lite
go-shelly is an unofficial native go client for the [Shelly Gen2 API](https://shelly-api-docs.shelly.cloud/gen2/). It's forked from jcodybaker's excellent [go-shelly](https://github.com/jcodybaker/go-shelly) library. However with the dependency to [mos](github.com/mongoose-os/mos/) completely removed. As many of its transitive dependencies were causing issues for me.

## Maturity
This library is mostly for my own personal amusement and no guarantees or promises about stability, functionality, documentation or test coverage are made.

## Usage
In most cases, using the library is as simple building an Resty client and calling .DoResty().  See the [go doc](https://pkg.go.dev/github.com/DonRobo/go-shelly-lite), the [Resty doc](https://resty.dev/) and [Shelly Gen2 API documentation](https://shelly-api-docs.shelly.cloud/gen2/).
```
import (
	"context"
	"fmt"
    "log"

	shelly "github.com/DonRobo/go-shelly-lite"
)

const (
    shellyAddress = "http://192.168.1.20"
)

func main() {
	client := resty.New()
	client.SetBaseURL(shellyAddress)
	defer client.Close()

	req := &shelly.ShellyGetStatusRequest{}
	statusResp, _, err := req.DoResty(client)

    if err != nil {
        log.Fatalf("querying device status: %v", err)
    }
    for i, sw := range statusResp.Switches {
        fmt.Printf("Switch %d status: %v\n", i, *sw.Output)
    }
}
```

## Legal

### License
This fork contains modifications by [DonRobo](https://github.com/DonRobo), based on the original work by J Cody Baker.  
Original copyright [J Cody Baker](cody@codybaker.com)  
Licensed under the [MIT license](LICENSE.md).