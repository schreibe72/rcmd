package azure

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var authorizer autorest.Authorizer

func getAuthorizer() (autorest.Authorizer, error) {
	var err error
	if authorizer == nil {
		authorizer, err = auth.NewAuthorizerFromEnvironment()
		if err != nil {
			return nil, err
		}
	}
	return authorizer, nil
}
