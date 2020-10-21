// Copyright (c) 2019-2020 Latona. All rights reserved.

package main

import (
	"bitbucket.org/latonaio/aion-core/pkg/go-client/msclient"
	"bitbucket.org/latonaio/aion-core/pkg/log"
	"context"
)

func main() {
	ctx := context.Background()
	log.SetFormat("sample-microservice-first")
	log.Printf("start sample-microservice(first)")
	c, err := msclient.NewKanbanClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = c.SetKanban("sample-microservice-first", c.GetProcessNumber()); err != nil {
		log.Fatal(err)
	}
	metadata := map[string]interface{}{"test": "metadata sample"}
	od, err := msclient.NewOutputData(msclient.SetMetadata(metadata))
	if err != nil {
		log.Fatal(err)
	}
	if err := c.OutputKanban(od); err != nil {
		log.Fatal(err)
	}
	log.Printf("finish sample-microservice(first)")
}
