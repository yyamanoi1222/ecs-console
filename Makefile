ecs-console: cmd/*/*.go internal/*/*.go
	go build  -o ecs-console cmd/ecs_console/main.go