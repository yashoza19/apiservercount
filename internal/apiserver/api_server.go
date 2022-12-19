package apiserver

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/openshift/api/apiserver/v1"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type Client interface {
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.APIRequestCount, error)
}

type operatorClient struct {
	Client runtimeClient.WithWatch
}

func addSchemes(scheme *runtime.Scheme) error {
	if err := v1.AddToScheme(scheme); err != nil {
		return err
	}

	if err := configv1.Install(scheme); err != nil {
		return err
	}

	return nil
}

func NewClient(kubeconfig *rest.Config) (Client, error) {
	scheme := runtime.NewScheme()

	if err := addSchemes(scheme); err != nil {
		return nil, fmt.Errorf("could not add schemes to client: %v", err)
	}

	client, err := runtimeClient.NewWithWatch(kubeconfig, runtimeClient.Options{Scheme: scheme})
	if err != nil {
		return nil, fmt.Errorf("could not get subscription client: %v", err)
	}

	var operatorClient Client = &operatorClient{
		Client: client,
	}
	return operatorClient, nil
}
