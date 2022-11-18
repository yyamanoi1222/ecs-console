package ecs

import (
  aecs "github.com/aws/aws-sdk-go/service/ecs"
  "github.com/aws/aws-sdk-go/aws"
  "testing"
)

type mockEcs struct {}
func (m mockEcs) RunTask(i *aecs.RunTaskInput) (*aecs.RunTaskOutput, error) {
  return &aecs.RunTaskOutput{
    Tasks: []*aecs.Task{
      &aecs.Task{
        TaskArn: aws.String("hoge"),
      },
    },
  }, nil
}
func (m mockEcs) StopTask(i *aecs.StopTaskInput) (*aecs.StopTaskOutput, error) {
  return &aecs.StopTaskOutput{}, nil
}
func (m mockEcs) WaitUntilTasksRunning(i *aecs.DescribeTasksInput) error {
  return nil
}
func (m mockEcs) DescribeTasks(i *aecs.DescribeTasksInput) (*aecs.DescribeTasksOutput, error) {
  return &aecs.DescribeTasksOutput{
    Tasks: []*aecs.Task{
      &aecs.Task{
        Containers: []*aecs.Container{
          &aecs.Container{
            Name: aws.String("hoge"),
            RuntimeId: aws.String("testContainerId"),
          },
        },
      },
    },
  }, nil
}

func TestRunEcsTask(t *testing.T) {
  type want struct {
    task *aecs.Task
    err error
  }
  ecsClient := mockEcs{}

  cases := []struct {
    name string
    config CreateTaskConfig
    want want
  }{
    {
      name: "valid",
      config: CreateTaskConfig{
        Client: ecsClient,
      },
      want: want{
        task: &aecs.Task{
          TaskArn: aws.String("hoge"),
        },
        err: nil,
      },
    },
  }

  for _, c := range cases {
    t.Run(c.name, func(t *testing.T) {
      task, err := RunEcsTask(c.config)
      if !(*c.want.task.TaskArn == *task.TaskArn && c.want.err == err) {
        t.Error("failed")
      }
    })
  }
}

func TestStopEcsTask(t *testing.T){
  ecsClient := mockEcs{}
  got := StopEcsTask(StopTaskConfig{
    Client: ecsClient,
    TaskArn: "hoge",
    ClusterName: "hoge",
  })
  var want error = nil

  if got != want {
    t.Error("failed")
  }
}

func TestGetContainerId(t *testing.T) {
  ecsClient := mockEcs{}
  gotContainerId, gotError := GetContainerId(ecsClient, "hoge", "hoge", "hoge")
  containerId := "testContainerId"
  var err error = nil

  if !(gotContainerId == containerId && gotError == err) {
    t.Error("failed")
  }
}
