//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

//go:generate wire

package builder

import (
	"database/sql"

	"github.com/google/wire"

	wireTools "github.com/smartlog/api/tools/wire"
)

var ApiSet = wire.NewSet(wireTools.NewClickHouse, NewApiService)

type ApiService struct {
	ClickHouse *sql.DB
}

func NewApiService(clickhouse *sql.DB) *ApiService {
	return &ApiService{
		ClickHouse: clickhouse,
	}
}

func Initialize() (*ApiService, func(), error) {
	panic(wire.Build(ApiSet))
}
