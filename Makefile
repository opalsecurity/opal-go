SHELL := /bin/bash

OPENAPI_GEN=openapi-generator generate -i api/openapi.yaml -g go -o . -c config.json
OPENAPI_GEN_CI=openapi-generator-cli generate --enable-post-process-file -i api/openapi.yaml -g go -o . -c config.json
PULL_REMOTE_OPENAPI=curl https://app.opal.dev/openapi.yaml > api/openapi.yaml

gen-openapi:
	$(OPENAPI_GEN)
gen-openapi-remote:
	$(PULL_REMOTE_OPENAPI)
	$(OPENAPI_GEN)
gen-openapi-remote-for-ci:
	$(PULL_REMOTE_OPENAPI)
	$(OPENAPI_GEN_CI)
