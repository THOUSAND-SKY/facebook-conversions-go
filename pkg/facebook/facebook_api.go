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
		timeout:     input.Timeout,
	}
}

// RawSend does not hash the fields but just posts instead.
//
// May want to use `PostEvent` instead.
func (a *API) RawSend(ctx context.Context, p facebookgen.EventRequest) (facebookgen.ResponseSuccess, *http.Response, error) {
	ctx = context.WithValue(ctx, facebookgen.ContextAPIKey, a.accessToken)
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
	hashUserDataFields(&p)
	success, _, err := a.RawSend(ctx, p)
	return success, err
}

// hashUserDataFields hashes event keys in-place.
//
// https://developers.facebook.com/docs/marketing-api/conversions-api/parameters
func hashUserDataFields(e *facebookgen.EventRequest) {
	if e == nil {
		return
	}
	for _, e2 := range e.Data {
		if e2.UserData != nil {
			u := e2.UserData
			u2 := &facebookgen.UserData{}
			u2.Em = hashString(u.Em)
			u2.Ph = hash(u.Ph)
			u2.Fn = hashString(u.Fn)
			u2.Ln = hashString(u.Ln)
			u2.Ge = hash(u.Ge)
			u2.Db = hash(u.Db)
			u2.Ct = hash(u.Ct)
			u2.St = hash(u.St)
			u2.Zp = hash(u.Zp)
			u2.Country = hashString(u.Country)
			u2.ExternalId = hash(u.ExternalId)
			e2.UserData = u2
		}
	}
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
