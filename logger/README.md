# Context logger

Use ContextLogging interface

LoggerRequestHandler is middleware for routing
```
router.Use(LoggerRequestHandler())
```

Can also manually log the start and end of request using 
```
log.StartContext(requestId)
log.EndContext(requestId)
```

Context logging 
```
log.PrintFContext(ctx,"log something with requestId taken from context")
log.PrintFRequest(requestId,"log something with requestId passed as parameter")
```