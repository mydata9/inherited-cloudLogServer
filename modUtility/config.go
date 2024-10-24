package modUtility

import (
	"fmt"
)

const CG_APPName = "cloudLogServer"

func config_Initialize() error {
	err := GetSingleGatlingConfig().Initialize(CG_APPName)
	if err != nil {
		fmt.Println("config initialize error:", err)
		return err
	}

	/*G_MongodbUrl = GetSingleGatlingConfig().Get(CG_Key_MongodbUrl)
	G_LocalLogPath = GetSingleGatlingConfig().Get(CG_Key_LocalLogPath)
	G_XApiKey = GetSingleGatlingConfig().Get(CG_Key_xApiKey)

	portStr := GetSingleGatlingConfig().Get(CG_Key_HttpPort)
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			G_HttpPort = port
		}
	}
	if G_HttpPort < 10 {
		G_HttpPort = GD_Default_HttpPort
	}*/

	return nil
}
