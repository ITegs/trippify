sdk:
	@openapi-generator generate --config openapi-gen.json --input-spec backend/apiserver/openapi-spec.yaml --generator-name typescript-axios --output frontend/lib/trippify-client