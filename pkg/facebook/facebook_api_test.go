package facebook

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/THOUSAND-SKY/facebook-conversions-go/pkg/facebookgen"
)

func TestAPI_PostEvent(t *testing.T) {
	type fields struct {
		accessToken string
		pixelID     string
		client      *facebookgen.DefaultApiService
		timeout     time.Duration
	}
	type args struct {
		ctx context.Context
		p   facebookgen.EventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    facebookgen.ResponseSuccess
		wantErr bool
	}{
		{
			name: "works",
			fields: fields{
				accessToken: os.Getenv("FB_ACCESS_TOKEN"),
				pixelID:     os.Getenv("FB_PIXEL_ID"),
				client:      &facebookgen.DefaultApiService{},
				timeout:     30000 * time.Second,
			},
			args: args{
				ctx: context.Background(),
				p: facebookgen.EventRequest{
					TestEventCode: os.Getenv("FB_TEST_EVENT_CODE"),
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
				},
			},
			want: facebookgen.ResponseSuccess{
				EventsReceived: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := New(APIInput{
				AccessToken: tt.fields.accessToken,
				PixelID:     tt.fields.pixelID,
				Timeout:     tt.fields.timeout,
			})
			got, err := a.PostEvent(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.PostEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.EventsReceived != tt.want.EventsReceived {
				t.Errorf("API.PostEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should work as in docs",
			args: args{
				s: "John_Smith@gmail.com",
			},
			// https://developers.facebook.com/docs/marketing-api/conversions-api/parameters/customer-information-parameters#em
			want: "62a14e44f765419d10fea99367361a727c12365e2520f32218d505ed9aa0f62f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashString(tt.args.s); got != tt.want {
				t.Errorf("hashString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashUserDataFields(t *testing.T) {
	t.Run("should not mutate arguments", func(t *testing.T) {
		fac := func() facebookgen.EventRequest {
			return facebookgen.EventRequest{
				Data: []facebookgen.Event{
					{
						EventName: "123",
						UserData: &facebookgen.UserData{
							Em: "123",
						},
					},
				},
			}

		}
		thing := fac()
		if got := hashUserDataFields(thing); !reflect.DeepEqual(thing, fac()) || reflect.DeepEqual(got, thing) {
			t.Errorf("hashUserDataFields() = %v, want %v", thing, fac())
		}
	})
}
