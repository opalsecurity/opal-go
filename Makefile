SHELL := /bin/bash

OPENAPI_GEN=openapi-generator generate --additional-properties=disallowAdditionalPropertiesIfNotPresent=false -i api/openapi.yaml -g go -o . -c config.json
OPENAPI_GEN_CI=openapi-generator-cli generate --additional-properties=disallowAdditionalPropertiesIfNotPresent=false --enable-post-process-file -i api/openapi.yaml -g go -o . -c config.json
PULL_REMOTE_OPENAPI=curl https://app.opal.dev/openapi.yaml > api/openapi.yaml
CLEAN=bash clean.sh
CLEAN_CI=CI=true bash clean.sh

gen-openapi:
	$(OPENAPI_GEN)
	$(CLEAN)
gen-openapi-remote:
	$(PULL_REMOTE_OPENAPI)
	$(OPENAPI_GEN)
	$(CLEAN)
gen-openapi-remote-for-ci:
	$(PULL_REMOTE_OPENAPI)
	$(OPENAPI_GEN_CI)
	$(CLEAN_CI)
clean:
	$(CLEAN)
clean-ci:
	$(CLEAN_CI)
