package service

import (
	"context"
	"fmt"

	"github.com/hungtran150/api-app/proto/v1/app_data_monitoring_bp"
)

// Create Alert notification
func (s *Service) CreateAlertNotification(ctx context.Context, req *app_data_monitoring_bp.SlackAlertRequest) (*app_data_monitoring_bp.SlackAlertResponse, error) {
	fmt.Println(req)
	return &app_data_monitoring_bp.SlackAlertResponse{
		Code: 200,
	}, nil
}

// Send Slack notification to external channel
func sendSlackAlert(message string) {

}