package service

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hungtran150/api-app/proto/v1/app_data_monitoring_bp"
	"google.golang.org/grpc/codes"
)

// Create Alert notification
func (s *Service) CreateButtonAlertNotification(ctx context.Context, req *app_data_monitoring_bp.SlackButtonRequest) (*app_data_monitoring_bp.SlackButtontResponse, error) {
	resp := &app_data_monitoring_bp.SlackButtontResponse{}

	for _, action := range req.Actions {
		if action.Name == "No" {
			resp.Code = int32(codes.OK)
			resp.Message = "cancel sending"
			return resp, nil
		}
	}
	content := req.Actions[0].Value
	

	sendSlackAlert(content, resp)

	return resp, nil
}

// Send Slack notification to external channel
func sendSlackAlert(message string, resp *app_data_monitoring_bp.SlackButtontResponse) {
	url, exist := os.LookupEnv("SLACK_WEB_HOOK")
	if !exist {
		log.Fatal("Please set SLACK_WEB_HOOK env")
		return
	}
	method := "POST"
	httpErrorFlag := true

	payload := strings.NewReader(message)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return
	}

	req.Header.Add("Content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	defer func () {
		if httpErrorFlag {
			resp.Code = int32(res.StatusCode)
			resp.Message = err.Error()
		}
		
	}()

	resp.Code = int32(res.StatusCode)
	resp.Message = string(body)
	res.Body.Close()
	httpErrorFlag = false
}