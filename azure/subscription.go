package azure

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-06-01/subscriptions"
)

type azureSubscriptions []subscriptions.Subscription

var as azureSubscriptions

func contains(a []string, needle string) bool {
	for _, i := range a {
		if i == needle {
			return true
		}
	}
	return false
}

func GetSubscriptions(names ...string) (azureSubscriptions, error) {
	var err error
	if len(as) != 0 {
		return as, nil
	}
	sc := subscriptions.NewClient()
	if sc.Authorizer, err = getAuthorizer(); err != nil {
		return azureSubscriptions{}, err
	}
	s, err := sc.List(context.Background())
	if err != nil {
		return azureSubscriptions{}, err
	}

	for s.NotDone() {

		for _, i := range s.Values() {
			if len(names) == 0 || contains(names, *i.DisplayName) {
				as = append(as, i)
			}
		}
		if err := s.Next(); err != nil {
			log.Println(err)
			break
		}
	}

	return as, nil
}

func (as azureSubscriptions) GetID(name string) (id string, err error) {
	err = nil
	for _, s := range as {
		if *s.DisplayName == name {
			id = *s.SubscriptionID
			return id, err
		}
	}
	err = fmt.Errorf("SubscriptionNotFound")
	return id, err
}

func (as azureSubscriptions) GetIDs() (ids []string) {
	for _, s := range as {
		ids = append(ids, *s.SubscriptionID)
	}
	return ids
}
