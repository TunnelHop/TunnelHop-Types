JSON_FILES = $(wildcard ./schemas/*.json)

generate: $(JSON_FILES)
	rm -rf common/*
	@for file in $^ ; do \
		echo "Processing" $${file} ; \
		filename=$$(basename -- "$$file" .json); \
        lowercase=$$(echo "$$filename" | tr '[:upper:]' '[:lower:]'); \
        go-jsonschema -p common "$$file" --resolve-extension json > "common/$$lowercase".go; \
		echo "\t- generated $$lowercase.go from" $${file} ; \
        capitalized=$$(echo "$${filename:0:1}" | tr '[:lower:]' '[:upper:]')$${filename:1}; \
        json2ts ./schemas/$$filename.json "./common/$$capitalized".ts; \
		echo "\t- generated $$capitalized.ts from" $${file} ; \
	done