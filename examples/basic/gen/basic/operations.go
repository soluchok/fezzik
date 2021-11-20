package basic

import (
	"context"

	"github.com/inigolabs/fezzik"
)

var OneAllTypesOperation string = `
query OneAllTypes {
   one {
      oneInt
      oneIntMust
      oneIntList
      oneIntMustList
      oneIntMustListMust
   }
}
`

type OneAllTypesResponse struct {
	One *struct {
		OneInt             *int
		OneIntMust         int
		OneIntList         *[]*int
		OneIntMustList     *[]int
		OneIntMustListMust []int
	}
}

func (c *client) OneAllTypes(ctx context.Context) (*OneAllTypesResponse, *fezzik.GQLErrors, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "OneAllTypes",
		Query:     OneAllTypesOperation,
		Variables: map[string]interface{}{},
	}

	var gqldata *OneAllTypesResponse
	var gqlerrs *fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, gqldata, gqlerrs)
	if err != nil {
		return nil, nil, err
	}
	return gqldata, gqlerrs, nil
}

var OneWithSubSelectionsOperation string = `
query OneWithSubSelections {
   one {
      two {
         twoInt
         twoStr
         three {
            threeInt
         }
      }
   }
}
`

type OneWithSubSelectionsResponse struct {
	One *struct {
		Two *struct {
			TwoInt *int
			TwoStr *string
			Three  *struct {
				ThreeInt *int
			}
		}
	}
}

func (c *client) OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, *fezzik.GQLErrors, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "OneWithSubSelections",
		Query:     OneWithSubSelectionsOperation,
		Variables: map[string]interface{}{},
	}

	var gqldata *OneWithSubSelectionsResponse
	var gqlerrs *fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, gqldata, gqlerrs)
	if err != nil {
		return nil, nil, err
	}
	return gqldata, gqlerrs, nil
}

var QueryWithInputsOperation string = `
query QueryWithInputs($input_one : String, $input_two : String) {
   one(input: $input_one) {
      oneInt
   }
   two(input: $input_two) {
      twoInt
   }
}
`

type QueryWithInputsInputArgs struct {
	InputOne *string `json:"input_one"`
	InputTwo *string `json:"input_two"`
}

type QueryWithInputsResponse struct {
	One *struct {
		OneInt *int
	}
	Two *struct {
		TwoInt *int
	}
}

func (c *client) QueryWithInputs(ctx context.Context, input *QueryWithInputsInputArgs) (
	*QueryWithInputsResponse, *fezzik.GQLErrors, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "QueryWithInputs",
		Query:     QueryWithInputsOperation,
		Variables: map[string]interface{}{
			"input_one": input.InputOne,
			"input_two": input.InputTwo,
		},
	}

	var gqldata *QueryWithInputsResponse
	var gqlerrs *fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, gqldata, gqlerrs)
	if err != nil {
		return nil, nil, err
	}
	return gqldata, gqlerrs, nil
}

var OneAddOperation string = `
mutation OneAdd($input : OneInput) {
   oneAdd(input: $input)
}
`

type OneAddInputArgs struct {
	Input *OneInput `json:"input"`
}

type OneAddResponse struct {
	OneAdd *string
}

func (c *client) OneAdd(ctx context.Context, input *OneAddInputArgs) (
	*OneAddResponse, *fezzik.GQLErrors, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "OneAdd",
		Query:     OneAddOperation,
		Variables: map[string]interface{}{
			"input": input.Input,
		},
	}

	var gqldata *OneAddResponse
	var gqlerrs *fezzik.GQLErrors
	err := c.gql.Query(ctx, gqlreq, gqldata, gqlerrs)
	if err != nil {
		return nil, nil, err
	}
	return gqldata, gqlerrs, nil
}
