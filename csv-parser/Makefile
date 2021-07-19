build:
	export GIT_COMMIT=$(git rev-list -1 HEAD) && go build -ldflags "-X main.GitCommit=${GIT_COMMIT}"
test:
	go test -v . && go test -v parse/*
vet:
	 go vet $(go list)
