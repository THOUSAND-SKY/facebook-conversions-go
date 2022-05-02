# Meta/Facebook Conversions API

A mostly-generated api client for the [Meta conversions API](https://developers.facebook.com/docs/marketing-api/conversions-api). It is generated from the [swagger definitions](https://github.com/facebookincubator/Facebook-Server-Side-API-Swagger). There is only a tiny wrapper added to make the api sane.

## Example

```go
import (
	"github.com/THOUSAND-SKY/facebook-conversions-go/pkg/facebookgen"
	"github.com/THOUSAND-SKY/facebook-conversions-go/pkg/facebook"
)

a := facebook.New(facebook.APIInput{
	AccessToken: tt.fields.accessToken,
	PixelID:     tt.fields.pixelID,
	Timeout:     tt.fields.timeout,
})
got, err := a.PostEvent(ctx, facebookgen.EventRequest{
	TestEventCode: "TEST123", // remove this before putting in production.
	Data: []facebookgen.Event{
		{
			EventName:      "Purchase",
			EventTime:      time.Now().Unix(),
			EventSourceUrl: "https://example.com",
			EventId:        fmt.Sprintf("EventID%v", time.Now().Unix()),
			UserData: &facebookgen.UserData{
				ClientIpAddress: "11.111.11.111",
				Country:         "en",
				ClientUserAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:91.0) Gecko/20100101 Firefox/91.0",
			},
			CustomData: &facebookgen.CustomData{
				Value:       1.00,
				Currency:    "USD",
				ContentType: "product",
				OrderId:     "123444567",
				ContentIds:  []string{"1234456"},
			},
		},
	},
})
```
