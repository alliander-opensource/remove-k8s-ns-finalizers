ALL:
	GO111MODULE=on go build -o remove-k8s-ns-finalizers .

fmt:
	GO111MODULE=on go fmt ./...
