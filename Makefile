GOCMD=go
GOVET=$(GOCMD) vet
GOIMPORTS=goimports
STATICCHECK=staticcheck

install: _install

_install:
	go install golang.org/x/tools/cmd/goimports@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/tools/cmd/godoc@latest

lint: _imports _vet _staticcheck

_imports:
	$(GOIMPORTS) -d main.go
	

_vet:
	$(GOVET) .
	

_staticcheck:
	$(STATICCHECK) main.go
	
