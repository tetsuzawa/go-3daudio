package service

import (
	"context"
	"github.com/tetsuzawa/go-3daudio/web-app/app/models"
	"github.com/tetsuzawa/go-3daudio/web-app/proto/hrtf"
)

type HRTFService struct {
}

func (s *HRTFService) GetHRTFFromName(ctx context.Context, req *hrtf.GetHRTFFromNameReq) (*hrtf.HRTFData, error) {
	name := req.GetName()
	var hrtfData = new(hrtf.HRTFData)
	dbHRTF, err := models.GetHRTFFromName(name)
	if err != nil {
		return nil, err
	}
	hrtfData.ID = dbHRTF.ID
	hrtfData.Name = dbHRTF.Name
	hrtfData.Path = dbHRTF.Path
	hrtfData.DatabaseName = dbHRTF.DatabaseName
	return hrtfData, nil
}
