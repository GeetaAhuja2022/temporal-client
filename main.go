package main

import (
	"context"
	"log"

	"github.com/GeetaAhuja2022/temporal-worker/common"
	"github.com/GeetaAhuja2022/temporal-worker/workflows"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create a Temporal client
	c, err := client.NewLazyClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Start a workflow execution
	workflowOptions := client.StartWorkflowOptions{
		ID:        "factory-workflow",
		TaskQueue: "my-task-queue",
	}

	input := make(map[string]interface{})
	userData := common.UserInput{
		To:      "geeta.ahuja@capitalone.com",
		Subject: "My Subject",
	}
	input["userData"] = userData
	paymentData := common.PaymentInput{
		Amount: "10000",
	}
	input["paymentData"] = paymentData

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.MyWorkflow, input)
	if err != nil {
		log.Fatalln("Unable to start workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
