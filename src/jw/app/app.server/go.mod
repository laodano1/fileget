module app/app.server

go 1.13

replace common => /../../common

require (
	github.com/davyxu/cellnet v4.1.0+incompatible
	github.com/davyxu/golog v0.1.0
	github.com/davyxu/goobjfmt v0.1.0 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/gorilla/websocket v1.4.1 // indirect
)
