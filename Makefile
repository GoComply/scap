GO=GO111MODULE=on go
GOBUILD=$(GO) build

all: regenerate-models test

test:
	$(GO) test ./...

.PHONY: gocomply_xsd2go
gocomply_xsd2go:
ifeq ("$(wildcard $(GOPATH)/bin/gocomply_xsd2go)","")
	go get -u -v github.com/gocomply/xsd2go/cli/gocomply_xsd2go
endif

regenerate-models: gocomply_xsd2go .scap_schemas
	gocomply_xsd2go convert .scap_schemas/schemas/cpe/2.3/cpe-dictionary_2.3.xsd github.com/gocomply/scap pkg/scap/models
	gocomply_xsd2go convert .scap_schemas/schemas/xccdf/1.2/xccdf_1.2.xsd github.com/gocomply/scap pkg/scap/models
	gocomply_xsd2go convert .scap_schemas/schemas/oval/5.11.3/oval-results-schema.xsd github.com/gocomply/scap pkg/scap/models/oval

.scap_schemas:
	git clone --depth 1 https://github.com/openscap/openscap .scap_schemas
