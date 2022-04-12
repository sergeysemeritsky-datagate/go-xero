package integration

import (
	"context"
	"github.com/glebteterin/go-xero"
	"golang.org/x/oauth2"
	"os"
)

//TODO: fail if no variables
//TODO: add failIfNotStatusCode(t, resp, 200) and check status or every response

var (
	client *xero.Client
)

func init() {
	token := os.Getenv("XERO_AUTH_TOKEN")
	tenantId := os.Getenv("XERO_TENANT_ID")

	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))
	client = xero.NewClient(tc)
	client.TenantId = tenantId
}
