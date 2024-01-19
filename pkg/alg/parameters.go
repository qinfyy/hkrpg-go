package alg

import (
	"os"
)

func GetAppId() string {
	var appId string
	for id, sa := range os.Args {
		if sa == "-i" {
			if len(os.Args) >= id+2 {
				appId = os.Args[id+1]
			}
			break
		}
	}
	if appId == "" {
		panic("AppId error")
	}
	return appId
}
