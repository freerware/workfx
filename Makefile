all: bins

clean:
	@echo cleaning...
	@GO111MODULE=on go clean -x
	@echo done!

bins:
	@echo building...
	@#v1 + v2
	@GO111MODULE=on go build github.com/freerware/workfx
	@#v4
	@cd ./v4 && GO111MODULE=on go build github.com/freerware/workfx/v4 && cd ..
	@echo done!

test: bins
	@echo testing...
	@#v1 + v2
	@GO111MODULE=on go test -v -race -covermode=atomic -coverprofile=workfx.coverprofile github.com/freerware/workfx
	@#v4
	@cd ./v4 && GO111MODULE=on go test -v -race -covermode=atomic -coverprofile=workfx.coverprofile github.com/freerware/workfx/v4 && cd ..
	@echo done!
