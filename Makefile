mock:
	mockgen -source vendor/github.com/uber-common/bark/interface.go -package mock_bark > mock_bark/mock_bark.go
