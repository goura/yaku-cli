OAPICODEGEN=$(shell go env GOPATH)/bin/oapi-codegen

deepl:
	# As of 2023/5/10, we need to manually edit the official openapi.yaml of DeepL
	# Also we need to manually edit "text" to remove the array support.
	# We have a diff.patch to do this in the schema/directory
	# See https://github.com/DeepLcom/openapi/issues/1
	mkdir -p gen/deepl
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
		-generate=types,client -include-tags=TranslateText,MetaInformation \
		-o gen/deepl/gen-deepl.go -package deepl schema/deepl/deepl-openapi-fixed.yaml

openai:
	# As of 2023/5/13, OpenAI's openapi.yaml is not tagged.
	# To limit the range of genaration, we manually tag the relavant operation.
	# We have a diff.patch to do this in the schema/directory.
	# -response-type-suffix=_ is needed to avoid name conflict with the generated types.
	mkdir -p gen/openaichatcomp
	$(OAPICODEGEN) \
		-generate=types,client -include-tags=ChatCompletion \
		-response-type-suffix=_ \
		-o gen/openaichatcomp/gen-openai.go -package openaichatcomp schema/openai/openai-openapi-patched.yaml
