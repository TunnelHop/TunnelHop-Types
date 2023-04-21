JSON_FILES = $(wildcard *.json)

generate: $(JSON_FILES)
	@for file in $^ ; do \
		echo "Processing" $${file} ; \
		filename=$$(basename -- "$$file" .json); \
        lowercase=$$(echo "$$filename" | tr '[:upper:]' '[:lower:]'); \
        go-jsonschema -p types "$$file" > "types/$$lowercase".go; \
		echo "\t- generated $$lowercase.go from" $${file} ; \
        capitalized=$$(echo "$${filename:0:1}" | tr '[:lower:]' '[:upper:]')$${filename:1}; \
        json2ts server.json "$$filename".ts; \
		echo "\t- generated $$capitalized.ts from" $${file} ; \
        mv "$$filename".ts "types/$$capitalized".ts; \
	done