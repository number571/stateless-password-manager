ROOPATH=.
BINPATH=$(ROOPATH)/bin

.PHONY: default build clean 

default: build 
build: 
	for app in spm; \
	do \
		for arch in amd64 arm64; \
		do \
			for platform in linux windows darwin; \
			do \
				echo "build $${app}_$${arch}_$${platform}"; \
				if [[ $$platform == "windows" ]] \
				then \
					CGO_ENABLED=0 GOOS=$${platform} GOARCH=$${arch} go build -o $(BINPATH)/$${app}_$${arch}_$${platform}.exe ./cmd/$${app}; \
				else \
					CGO_ENABLED=0 GOOS=$${platform} GOARCH=$${arch} go build -o $(BINPATH)/$${app}_$${arch}_$${platform} ./cmd/$${app}; \
				fi; \
			done; \
		done; \
	done;

clean:
	rm -rf ./bin
