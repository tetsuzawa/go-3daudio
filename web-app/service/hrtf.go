package service

import (
	"context"
	"github.com/tetsuzawa/go-3daudio/web-app/proto/hrtf"
)

type HRTFService struct {
}

func (s *HRTFService) GetRandomHRTF(ctx context.Context, req *hrtf.GetHRTFFromNameReq) (*hrtf.HRTFData, error) {
	dbName := req.GetName()

}
