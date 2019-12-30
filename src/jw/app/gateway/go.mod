module app/gateway

go 1.13

replace common => /../../common

require (
	common v0.0.0
	github.com/davyxu/cellnet v4.1.0+incompatible
	github.com/davyxu/golog v0.1.0
)
