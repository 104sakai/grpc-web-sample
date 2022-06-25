PKG=github.com/104sakai/grpc-web-sample
FEATURE=webapi/server/feature
SERVICE=webapi/server/service
PLATFORM=--platform=linux/x86_64
UID := $(shell id -u)

.PHONY: protoc

protoc: proto/*.proto
	cd proto; \
	for n in $(subst .proto, , $(notdir $^)); do \
		/usr/local/bin/entrypoint.sh -f $$n.proto -o gen/pb-go/$$n -l go --with-gateway --with-openapi-json-names; \
		/usr/local/bin/entrypoint.sh -f $$n.proto -o gen/pb-web/$$n -l web; \
		mv gen/pb-go/$$n/$$n.swagger.json gen/pb-web/$$n; \
		mkdir -p ../$(FEATURE)/$$n; \
		mkdir -p ../$(SERVICE)/$$n; \
		touch ../$(SERVICE)/$$n/service.go; \
		mv gen/pb-go/$$n/$(PKG)/$(FEATURE)/$$n/* ../$(FEATURE)/$$n/; \
		mkdir -p ../client/src/assets/protobuf/web/$$n; \
		mv gen/pb-web/$$n/* ../client/src/assets/protobuf/web/$$n/; \
	done
	mkdir -p client/src/assets/protobuf/web/google/api/
	echo {} > client/src/assets/protobuf/web/google/api/annotations_pb.js
	rm -r proto/gen