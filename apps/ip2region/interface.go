package ip2region

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Service todo
type Service interface {
	UpdateDBFile(context.Context, *UpdateDBFileRequest) error
	LookupIP(ctx context.Context, ip string) (*IPInfo, error)
}

// NewDefaultIPInfo todo
func NewDefaultIPInfo() *IPInfo {
	return &IPInfo{}
}

// IPInfo todo
type IPInfo struct {
	CityID   int64  `bson:"city_id" json:"city_id"`
	Country  string `bson:"country" json:"country"`
	Region   string `bson:"region" json:"region"`
	Province string `bson:"province" json:"province"`
	City     string `bson:"city" json:"city"`
	ISP      string `bson:"isp" json:"isp"`
}

func (ip IPInfo) String() string {
	return strconv.FormatInt(ip.CityID, 10) + "|" + ip.Country + "|" + ip.Region + "|" + ip.Province + "|" + ip.City + "|" + ip.ISP
}

// NewUploadFileRequestFromHTTP todo
func NewUploadFileRequestFromHTTP(r *http.Request) (*UpdateDBFileRequest, error) {
	req := &UpdateDBFileRequest{
		reader: r.Body,
	}
	return req, nil
}

// UpdateDBFileRequest 上传文件请求
type UpdateDBFileRequest struct {
	reader io.ReadCloser
}

// Validate 校验参数
func (req *UpdateDBFileRequest) Validate() error {
	if req.reader == nil {
		return fmt.Errorf("file reader is nil")
	}

	return nil
}

// ReadCloser todo
func (req *UpdateDBFileRequest) ReadCloser() io.ReadCloser {
	return req.reader
}
