# ecs-console
ecs-console is a tool for launching one-time tasks and running shells and portforwarding on ECS

# Motivation
When using a CLI tool such as Ralis console in a production environment, there is a need to start and stop a one-time ECS task each time.  
This tool hides the start-stop process, so developers can start shells and perform port forwarding without being aware of the task start-stop process.

# Install

```
$ go install github.com/yyamanoi1222/ecs-console@latest
```

# Usage

You can see the detailed options in the -h option

## Execute shell

```
$ ecs-console exec --cluster <cluster-name> --task-def <task-def> --command <execute command (default /bin/bash)> --container <container name for login> --subnets <subnet names for task placement> --security-groups <sg ids for task placement>
```

## Execute portforward

```
$ ecs-console portforward --cluster <cluster-name> --task-def <task-def> --container <container name for login> --subnets <subnet names for task placement> --security-groups <sg ids for task placement> --local-port <local port number> --remote-port <remote port number>
```
