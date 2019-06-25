package main

import (
	"fmt"

	"github.com/bar-counter/checktag"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// TODO init test here
	cfg := &checktag.Cfg{}
	err := checktag.Register(r, cfg)
	if err != nil {
		fmt.Printf("test Register err %v\n", err)
		return
	}
	nowCfg := checktag.NowCfg()
	err = r.Run(nowCfg.RunAddr)
	if err != nil {
		fmt.Printf("run err %v\n", err)
		return
	}
}
