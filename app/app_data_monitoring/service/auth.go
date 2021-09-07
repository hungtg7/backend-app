package service

import "context"

func (s *Service) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	// Skip Authen
	// TODO add slacktoken to context
	return ctx, nil
}