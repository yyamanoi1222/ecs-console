# ecs-console
ecs-console is a tool for launching one-time ECS tasks and running ecs-exec or portforwarding on ECS

# Motivation
When using a CLI tool such as Ralis console in a production environment, there is a need to start and stop a one-time ECS task each time.  
This tool hides the start-stop process, so developers can start shells and perform port forwarding without being aware of the task start-stop process.

# Install

```
$ go install github.com/yyamanoi1222/ecs-console/cmd/ecs-console@latest
```

# Usage

You can see the detailed options in the -h option

## Execute shell

```
$ ecs-console exec -h

NAME:
   ecs-console exec

USAGE:
   ecs-console exec [command options] [arguments...]

OPTIONS:
   --cluster value          ECS Cluster Name
   --task-def value         ECS Taskdefinition arn
   --command value          command passing to ecs-exec (default: "/bin/bash")
   --container value        container name for ecs-exec (default: "app")
   --subnets value          subnets name for task
   --security-groups value  sg for task
   --help, -h               show help (default: false)

```

## Execute portforward

```
$ ecs-console portforward -h

NAME:
   ecs-console portforward

USAGE:
   ecs-console portforward [command options] [arguments...]

OPTIONS:
   --cluster value          ECS Cluster Name
   --task-def value         ECS Taskdefinition arn
   --container value        container name for ecs-exec (default: "app")
   --subnets value          subnets name for task
   --security-groups value  sg for task
   --remote-port value      remote port
   --local-port value       local port
   --help, -h               show help (default: false)

```
