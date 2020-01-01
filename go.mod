module github.com/vulcan-frame/vulcan-gate

go 1.15

require (
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.0.0
	github.com/golang/protobuf v1.1.3
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.1
	github.com/prometheus/client_golang v1.8.0
	github.com/prometheus/common v0.15.0
	github.com/redis/go-redis/v6 v6.0.0
	github.com/stretchr/testify v1.6.1
	go.etcd.io/etcd/api/v3 v3.1.0-alpha.0
	go.etcd.io/etcd/client/pkg/v3 v3.5.0-alpha.0
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20201216054612-986b41b23924
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/sys v0.0.0-20201214210602-f9fddec55a7e
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d
	google.golang.org/grpc v1.5.0
)

replace (
	github.com/vulcan-frame/vulcan-pkg-tool => ./pkg/tool/vulcan-pkg-tool
	github.com/vulcan-frame/vulcan-pkg-app => ./pkg/app/vulcan-pkg-app
)
