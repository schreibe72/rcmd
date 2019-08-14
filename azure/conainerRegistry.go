package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/Azure/go-autorest/autorest/azure"
)

type Registry struct {
	LoginServer   string
	Name          string
	ResourceGroup string
	Login         string
	Password      string
}

func GetContainerRegistries(subscription string) (registries []Registry, err error) {
	cc := containerregistry.NewRegistriesClient(subscription)
	if cc.Authorizer, err = getAuthorizer(); err != nil {
		return registries, err
	}

	r, err := cc.ListComplete(context.Background())
	if err != nil {
		return registries, err
	}
	for r.NotDone() {
		item := r.Value()
		idParts, err := azure.ParseResourceID(*item.ID)
		if err != nil {
			return registries, err
		}
		name := *item.Name
		loginServer := *item.LoginServer
		resourceGroup := idParts.ResourceGroup

		cred, err := cc.ListCredentials(context.Background(), resourceGroup, name)
		if err != nil {
			return registries, err
		}
		login := *cred.Username
		password := ""
		for _, i := range *cred.Passwords {
			if i.Name == "password" {
				password = *i.Value
			}
		}

		registries = append(registries, Registry{
			Name:          name,
			LoginServer:   loginServer,
			ResourceGroup: idParts.ResourceGroup,
			Login:         login,
			Password:      password,
		})
		if err := r.Next(); err != nil {
			return registries, err

		}
	}
	return registries, nil
}
