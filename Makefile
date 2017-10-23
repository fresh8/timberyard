mock:
	go get github.com/uber-common/bark
	mockgen github.com/uber-common/bark Logger > mock_bark/mock_bark.go
