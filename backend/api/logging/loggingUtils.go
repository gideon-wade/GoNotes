package logging

import "github.com/gin-gonic/gin"

func RequestReceived(ctx *gin.Context) LogEvent {
	return NewLogEvent(
		InfoLevel, 
		"Received request: " + ctx.Request.Method + " " + ctx.FullPath())
}

func RequestCompleted(ctx *gin.Context) LogEvent {
	return NewLogEvent(
		InfoLevel, 
		"Completed request: " + ctx.Request.Method + " " + ctx.FullPath())
}

