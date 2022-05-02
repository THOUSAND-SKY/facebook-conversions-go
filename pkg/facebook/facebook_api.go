package facebook

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/THOUSAND-SKY/facebook-conversions-go/pkg/facebookgen"
)

const defaultTimeoutSecs = 30

type API struct {
	accessToken string
	pixelID     string
	client      *facebookgen.DefaultApiService
	timeout     time.Duration
}

type APIInput struct {
	AccessToken string
	PixelID     string

	// Timeout for the http request. Defaults to 30s
	Timeout time.Duration
}

func New(input APIInput) *API {
	t := input.Timeout
	if t == 0 {
		t = time.Second * defaultTimeoutSecs
	}
	return &API{
		client:      facebookgen.NewAPIClient(facebookgen.NewConfiguration()).DefaultApi,
		accessToken: input.AccessToken,
		pixelID:     input.PixelID,
		timeout:     t,
	}
}

// RawSend does not hash the fields but just posts instead.
//
// May want to use `PostEvent` instead.
func (a *API) RawSend(ctx context.Context, p facebookgen.EventRequest) (facebookgen.ResponseSuccess, *http.Response, error) {
	ctx = context.WithValue(ctx, facebookgen.ContextAPIKey, facebookgen.APIKey{
		Key: a.accessToken,
	})
	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	return a.client.PixelIdEventsPost(ctx, p, a.pixelID)
}

// PostEvent hashes the fields that require hashing (UserData) and sends the event.
//
// This call will hash fields, but the caller is expected to make sure they're in correct format.
//
// https://developers.facebook.com/docs/marketing-api/conversions-api/parameters/customer-information-parameters
func (a *API) PostEvent(ctx context.Context, p facebookgen.EventRequest) (facebookgen.ResponseSuccess, error) {
	success, _, err := a.RawSend(ctx, hashUserDataFields(p))
	return success, err
}

// hashUserDataFields hashes event keys.
//
// https://developers.facebook.com/docs/marketing-api/conversions-api/parameters
func hashUserDataFields(e facebookgen.EventRequest) facebookgen.EventRequest {
	newData := make([]facebookgen.Event, 0)
	for _, e2 := range e.Data {
		if e2.UserData != nil {
			u := e2.UserData
			e2.UserData = &facebookgen.UserData{
				Em:              hashString(u.Em),
				Ph:              hash(u.Ph),
				Fn:              hashString(u.Fn),
				Ln:              hashString(u.Ln),
				Ge:              hash(u.Ge),
				Db:              hash(u.Db),
				Ct:              hash(u.Ct),
				St:              hash(u.St),
				Zp:              hash(u.Zp),
				Country:         hashString(u.Country),
				ExternalId:      hash(u.ExternalId),
				ClientIpAddress: u.ClientIpAddress,
				ClientUserAgent: u.ClientUserAgent,
				Fbc:             u.Fbc,
				Fbp:             u.Fbp,
				SubscriptionId:  u.SubscriptionId,
				FbLoginId:       u.FbLoginId,
			}
			newData = append(newData, e2)
		}
	}
	e.Data = newData
	return e
}

func hashString(s string) string {
	return hash(strings.ToLower(s))
}

func hash(thingToHash string) string {
	if thingToHash == "" {
		return ""
	}
	h := sha256.Sum256([]byte(thingToHash))
	return fmt.Sprintf("%x", h)
}
