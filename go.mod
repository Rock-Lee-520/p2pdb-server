module github.com/kkguan/p2pdb-server

go 1.17

require (
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/dolthub/sqllogictest/go v0.0.0-20201107003712-816f3ae12d81
	github.com/dolthub/vitess v0.0.0-20211215165926-1490f8c93e81
	github.com/go-kit/kit v0.12.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gocraft/dbr/v2 v2.7.3
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/oliveagle/jsonpath v0.0.0-20180606110733-2e52cf6e6852 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/sanity-io/litter v1.5.1
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.8.1
	github.com/src-d/go-oniguruma v1.1.0 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/src-d/go-errors.v1 v1.0.0
)

require (
	github.com/caarlos0/env/v6 v6.9.1
	github.com/kkguan/p2pdb-store v0.0.6
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/favframework/debug v0.0.0-20150708094948-5c7e73aafb21 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.11 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf // indirect
	golang.org/x/sys v0.0.0-20210917161153-d61c044b1678 // indirect
	google.golang.org/genproto v0.0.0-20210917145530-b395a37504d4 // indirect
	google.golang.org/grpc v1.40.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/kkguan/p2pdb-server => ../p2pdb-server

replace github.com/kkguan/p2pdb-store => ../p2pdb-store

replace github.com/kkguan/p2pdb-store/sqlite; => ../p2pdb-store/sqlite
