module log-pilot

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/containerd/containerd v1.5.7
	github.com/containerd/typeurl v1.0.2
	github.com/docker/docker v20.10.8+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/elastic/go-ucfg v0.8.3
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/moby/term v0.0.0-20210610120745-9d4ed1856297 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/opencontainers/selinux v1.8.3 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210520170846-37e1c6afe023
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.39.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	github.com/kubernetes-sigs/cri-tools v1.22.0

)

replace github.com/sirupsen/logrus v1.8.1 => github.com/Sirupsen/logrus v1.8.1

go 1.13
