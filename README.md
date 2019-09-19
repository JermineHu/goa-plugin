# goa-plugin
To generate golang file for get actions by goa plugin. (for goav1.4.1)
# How to use

```
# to get plugin
go get github.com/JermineHu/goa-plugin/gen-res

# to generate code by plugin
goagen gen -d /your/design/path/ --pkg-path=github.com/JermineHu/goa-plugin/gen-res -o app
```

example

```
[jermine@jermine-pc crius]$ cat Makefile 
#! /usr/bin/make
#
# Makefile for crius
#
# Targets:
# - clean     delete all generated files
# - generate  (re)generate all goagen-generated files.
# - build     compile executable
# - ae-build  build appengine
# - ae-dev    deploy to local (dev) appengine
# - ae-deploy deploy to appengine
#
# Meta targets:
# - all is the default target, it runs all the targets in the order above.
#
DEPEND= bitbucket.org/pkg/inflect \
        github.com/goadesign/goa \
        github.com/goadesign/goa/goagen \
        github.com/goadesign/goa/logging/logrus \
        github.com/sirupsen/logrus \
        gopkg.in/yaml.v2 \
        golang.org/x/tools/cmd/goimports

CURRENT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: depend clean generate build

depend:
        @go get $(DEPEND)

clean:
        @rm -rf app
        @rm -rf client
        @rm -rf tool
        @rm -rf public/swagger
        @rm -rf public/schema
        @rm -rf public/js
        @rm -f crius

generate:
        @goagen app -d  example.cn/crius/design/apis
        @goagen client -d  example.cn/crius/design/apis
        @goagen main -d  example.cn/crius/design/apis --force -o tmp
        @goagen schema -d example.cn/crius/design/apis -o public
        @goagen swagger -d example.cn/crius/design/apis -o public
        @goagen js -d example.cn/crius/design/apis -o public
        @goagen gen -d example.cn/crius/design/apis --pkg-path=github.com/JermineHu/goa-plugin/gen-res -o app

build:
        @go build -o crius

ae-build:
        @if [ ! -d $(HOME)/crius ]; then \
                mkdir $(HOME)/crius; \
                ln -s $(CURRENT_DIR)/appengine.go $(HOME)/crius/appengine.go; \
                ln -s $(CURRENT_DIR)/app.yaml     $(HOME)/crius/app.yaml; \
        fi

ae-deploy: ae-build
        cd $(HOME)/crius
        gcloud app deploy --project crius
[jermine@jermine-pc crius]$ 
[jermine@jermine-pc crius]$ make generate
app
app/contexts.go
app/controllers.go
app/security.go
app/hrefs.go
app/media_types.go
app/user_types.go
app/test
app/test/account_testing.go
app/test/alertgroup_testing.go
app/test/alertrule_testing.go
app/test/api_logs_testing.go
app/test/cluster_testing.go
app/test/cluster_type_testing.go
app/test/company_testing.go
app/test/company_request_notice_testing.go
app/test/health_testing.go
app/test/ingress_testing.go
app/test/namespace_testing.go
app/test/node_testing.go
app/test/notice_testing.go
app/test/persistent_volume_testing.go
app/test/persistent_volume_claim_testing.go
app/test/project_testing.go
app/test/projectalertgroup_testing.go
app/test/projectalertrule_testing.go
app/test/rbac_testing.go
app/test/service_testing.go
app/test/statistic_testing.go
app/test/storage_testing.go
app/test/subscribe_testing.go
tool/cli
tool/cli/commands.go
client
client/client.go
client/account.go
client/alertgroup.go
client/alertrule.go
client/api_logs.go
client/cluster.go
client/cluster_type.go
client/company.go
client/company_request_notice.go
client/health.go
client/ingress.go
client/namespace.go
client/node.go
client/notice.go
client/persistent_volume.go
client/persistent_volume_claim.go
client/project.go
client/projectalertgroup.go
client/projectalertrule.go
client/rbac.go
client/service.go
client/statistic.go
client/storage.go
client/subscribe.go
client/user_types.go
client/media_types.go
tmp/main.go
tmp/account.go
tmp/alertgroup.go
tmp/alertrule.go
tmp/api_logs.go
tmp/cluster.go
tmp/cluster_type.go
tmp/company.go
tmp/company_request_notice.go
tmp/health.go
tmp/ingress.go
tmp/namespace.go
tmp/node.go
tmp/notice.go
tmp/persistent_volume.go
tmp/persistent_volume_claim.go
tmp/project.go
tmp/projectalertgroup.go
tmp/projectalertrule.go
tmp/rbac.go
tmp/service.go
tmp/statistic.go
tmp/storage.go
tmp/subscribe.go
public/schema
public/schema/schema.json
public/swagger
public/swagger/swagger.json
public/swagger/swagger.yaml
public/js
public/js/client.js
public/js/axios.min.js
public/js/index.html
public/js/example.go
app/res_actions.go
[jermine@jermine-pc crius]$ 

```
