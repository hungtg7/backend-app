package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hungtran150/api-app/proto/v1/app_data_monitoring_bp"
	"google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

// Create Alert notification
func (s *Service) CreateButtonAlertNotification(ctx context.Context, req *app_data_monitoring_bp.SlackButtonRequest) (*app_data_monitoring_bp.SlackButtontResponse, error) {
	resp := &app_data_monitoring_bp.SlackButtontResponse{}
	fmt.Println(req)

	for _, action := range req.Payload.Actions {
		if action.Value == "no" {
			resp.Code = int32(codes.OK)
			resp.Message = "cancle sending"
			return resp, nil
		}
	}
	messageContent := req.Payload.Actions[0].Value

	sendSlackAlert(messageContent, resp)

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

	payload := strings.NewReader(fmt.Sprintf(`{"text":"%s"}`, message))

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