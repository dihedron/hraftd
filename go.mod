module github.com/dihedron/brokerd

go 1.16

require (
	github.com/armon/go-metrics v0.3.6 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/hashicorp/go-hclog v0.15.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.0 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/raft v1.2.0
	github.com/hashicorp/raft-boltdb v0.0.0-20191021154308-4207f1bf0617
	github.com/jessevdk/go-flags v1.4.0
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/stretchr/testify v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	goa.design/goa/v3 v3.2.6
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/mod v0.4.1 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	honnef.co/go/tools v0.1.1 // indirect
)

replace github.com/hashicorp/raft-boltdb => github.com/dihedron/raft-boltdb v0.0.0-20210115232206-5b95c94bbbce
