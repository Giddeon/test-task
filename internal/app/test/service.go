package test

import (
	"test/api/test"
	"test/internal/repositories"
	swagger "test/pkg/clients/garantex"
)

type Implementation struct {
	test.UnimplementedTestServer

	rq       repositories.RateQuery
	depthApi swagger.DepthApiClient
}

func NewTest(
	tnq repositories.RateQuery,
	depthApi swagger.DepthApiClient,
) *Implementation {
	return &Implementation{
		rq:       tnq,
		depthApi: depthApi,
	}
}
