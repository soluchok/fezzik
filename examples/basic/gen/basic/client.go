package basic

import (
	"context"
	"net/http"

	"github.com/inigolabs/fezzik/client"
)

type Client interface {
	// OneAllTypes from examples/basic/operations/operations.graphql:2
	OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error)

	// OneWithSubSelections from examples/basic/operations/operations.graphql:13
	OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error)

	// QueryWithInputs from examples/basic/operations/operations.graphql:27
	QueryWithInputs(ctx context.Context,
		inputOne *string,
		inputTwo *string,
	) (*QueryWithInputsResponse, error)

	// OneAdd from examples/basic/operations/operations.graphql:38
	OneAdd(ctx context.Context,
		input *OneInput,
	) (*OneAddResponse, error)
}

func NewClient(url string, httpclient *http.Client) Client {
	return &gqlclient{
		gql: client.NewGQLClient(url, httpclient),
	}
}

type gqlclient struct {
	gql *client.GQLClient
}
