package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber"
)

type Hook struct {
	dashboardId string `json:"dashboardId"`
	evalMatches string `json:"evalMatches"`
	ImageUrl    string `json:"imageUrl"`
	Message     string `json:"message"`
	orgId       string `json:"orgId"`
	panelId     string `json:"panelId"`
	ruleId      string `json:"ruleId"`
	ruleName    string `json:"ruleName"`
	RuleUrl     string `json:"ruleUrl"`
	state       string `json:"state"`
	tags        string `json:"tags"`
	Title       string `json:"title"`
}

func GwWorker() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		h := new(Hook)
		if err := c.BodyParser(h); err != nil {
			fmt.Println(err)
			c.Send("Error on JSON format")
			return
		}
		fmt.Println(h.Title)

		// Send to WeChat Work

		// {
		// 	"msgtype": "news",
		// 	"news": {
		// 	  "articles": [
		// 		{
		// 		  "title": "%s",
		// 		  "description": "%s",
		// 		  "url": "%s",
		// 		  "picurl": "%s"
		// 		}
		// 	  ]
		// 	}
		//   }
		url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + c.Params("key")

		msgStr := fmt.Sprintf(`
		{
			"msgtype": "news",
			"news": {
			  "articles": [
				{
				  "title": "%s",
				  "description": "%s",
				  "url": "%s",
				  "picurl": "%s"
				}
			  ]
			}
		  }
		`, h.Title, h.Message, h.RuleUrl, h.ImageUrl)
		fmt.Println(msgStr)
		jsonStr := []byte(msgStr)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.Send("Error sending to WeChat Work API")
			return
		}
		defer resp.Body.Close()
		c.Send(resp)
	}
}
