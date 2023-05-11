OAPICODEGEN=go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen

deepl:
	# As of 2023/5/10, we need to manually edit the official openapi.yaml of DeepL
	# Also we need to manually edit "text" to remove the array support.
	# We have a diff.patch to do this in the schema/directory
	# See https://github.com/DeepLcom/openapi/issues/1
	mkdir -p gen/deepl
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
		-generate=types,client -include-tags=TranslateText,MetaInformation \
		-o gen/deepl/gen-deepl.go -package deepl schema/deepl-openapi-fixed.yaml