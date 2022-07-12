package rest

// func APIResponse(c *gin.Context, status int, message []string, payload interface{},
// 	notifications []*v1.Notification, redirectURL string, errors []string) {
// 	res := client.Response{
// 		Notifications: notifications,
// 		Redirect:      redirectURL,
// 		Errors:        errors,
// 		Messages:      message,
// 	}

// 	jsonConfigs, err := json.Marshal(&res)
// 	if err != nil {
// 		logrus.Errorf("can not convert client response message to json, res: %s, err: %s \n",
// 			res.String(), err.Error())
// 	}

// 	//c.Data(status, gin.MIMEJSON, bytes)
// 	//c.String(status, string(bytes))
// 	c.JSON(status, gin.H{
// 		"metadata": string(jsonConfigs),
// 		"payload":  payload,
// 	})
// }