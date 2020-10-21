// Copyright (c) 2019-2020 Latona. All rights reserved.

package main

import (
	"bitbucket.org/latonaio/aion-core/pkg/go-client/msclient"
	"bitbucket.org/latonaio/aion-core/pkg/log"
	"context"
)

func main() {
	ctx := context.Background()
	log.SetFormat("sample-microservice-second")
	log.Printf("start sample-microservice(second)")
	c, err := msclient.NewKanbanClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// カンバンを１つ取得する
	// GetKanbanChでchannelを取得可能
	k, err := c.GetOneKanban("sample-microservice-second", c.GetProcessNumber())
	if err != nil {
		log.Fatal(err)
	}
	// メタデータをmap[string]interface{}で取得する
	metadata, err := k.GetMetadataByMap()
	if err != nil {
		log.Fatal(err)
	}
	// メタデータの内容を出力する
	// mapの中身はInterfaceなので、キャストする必要がある
	// 自分で用意したStructに嵌め込んでも良い
	log.Printf("output metadata list")
	for k, v := range metadata {
		if str, ok := v.(string); !ok {
			log.Printf("cant cast to string(key: %s)", k)
		} else {
			log.Printf("%s: %s", k, str)
		}
	}
	metadata["test"] = "changed value"

	// カンバンの内容をデータとして用意する
	// オプションの例を全て示す
	od, err := msclient.NewOutputData(
		msclient.SetMetadata(metadata),
		msclient.SetConnectionKey("sample-connection-key"),
		msclient.SetDeviceName("next-device"),
		msclient.SetFileList([]string{"test.txt"}),
		msclient.SetResult(false),
		msclient.SetDataPath("/var/lib/aion/"),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.OutputKanban(od); err != nil {
		log.Fatal(err)
	}
	log.Printf("finish sample-microservice(second)")
}
