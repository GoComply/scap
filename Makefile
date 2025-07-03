GO=GO111MODULE=on go
GOBUILD=$(GO) build

all: regenerate-models test

test:
	$(GO) test ./...

.PHONY: gocomply_xsd2go vendor
gocomply_xsd2go:
ifeq ("$(wildcard $(GOPATH)/bin/gocomply_xsd2go)","")
	go install github.com/gocomply/xsd2go/cli/gocomply_xsd2go@latest
endif

regenerate-models: gocomply_xsd2go .scap_schemas
	find  pkg/scap/models -name models.go | xargs rm --
	gocomply_xsd2go convert .scap_schemas/schemas/cpe/2.3/cpe-dictionary_2.3.xsd github.com/gocomply/scap pkg/scap/models
	gocomply_xsd2go convert \
		--xmlns-override=http://cpe.mitre.org/language/2.0=cpe_language \
		.scap_schemas/schemas/xccdf/1.2/xccdf_1.2.xsd github.com/gocomply/scap pkg/scap/models
	gocomply_xsd2go convert .scap_schemas/schemas/oval/5.11.3/oval-results-schema.xsd github.com/gocomply/scap pkg/scap/models
	gocomply_xsd2go convert \
		--xmlns-override=http://cpe.mitre.org/language/2.0=cpe_language \
		.scap_schemas/schemas/sds/1.3/scap-source-data-stream_1.3.xsd github.com/gocomply/scap pkg/scap/models
	gofmt -s -w ./pkg/scap/models


.scap_schemas:
	git clone --depth 1 https://github.com/openscap/openscap .scap_schemas

vendor:
	$(GO) mod tidy
	$(GO) mod vendor
	$(GO) mod verify
