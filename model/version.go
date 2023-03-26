package model

var (
	LOGO = `
               _       _         _
__      ____ _| |_ ___| |__   __| | ___   __ _
\ \ /\ / / _  | __/ __| '_ \ / _  |/ _ \ / _  |
 \ V  V / (_| | || (__| | | | (_| | (_) | (_| |
  \_/\_/ \__,_|\__\___|_| |_|\__,_|\___/ \__, |
                                         |___/
`
	EMOJI    = map[string]string{"sccwatchdog": "🐶"}
	Historys = []History{
		{Version: "0.1.0",
			Description: "a scc deployment image and lastupdate time monitor",
		},
		{Version: "0.2.0",
			Description: "changed the way to get image and lastUpdateTime",
		}, {Version: "0.2.1",
			Description: "watch image changed and deployment restarted separated",
		}, {Version: "0.2.2",
			Description: "watch 'deployments' and 'all deployments' separated",
		}, {Version: "0.2.3",
			Description: "offer jsonpath to get infomations",
		},
		{Version: "0.3.0",
			Description: "deployment restarted show the message",
		},
		{Version: "0.4.0",
			Description: "add restart deployment and set image function",
		}, {Version: "0.4.1",
			Description: "watch restarted show the image",
		}, {Version: "0.4.2",
			Description: "change the dispaly way of watch image version",
		}, {Version: "0.4.3",
			Description: "fix No resources found in specified namespace problem",
		},
		{Version: "0.5.0",
			Description: "watch -e : not watch expect deployments",
		}, {Version: "0.5.1",
			Description: "fix Context problem",
		},
		{Version: "0.6.0",
			Description: "interactive client",
		}, {
			Version:     "0.6.1",
			Description: "fix splitImageInfo error",
		},
	}
)
