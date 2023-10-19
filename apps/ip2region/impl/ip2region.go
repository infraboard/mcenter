package impl

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/ioc/apps/oss"

	"github.com/infraboard/mcenter/apps/ip2region"
)

func (s *service) UpdateDBFile(ctx context.Context, req *ip2region.UpdateDBFileRequest) error {
	if err := req.Validate(); err != nil {
		return exception.NewBadRequest("validate update db file requrest error, %s", err)
	}

	reader := req.ReadCloser()
	defer reader.Close()

	uploadReq := oss.NewUploadFileRequest(s.bucketName, s.dbFileName, reader)
	return s.oss.UploadFile(ctx, uploadReq)
}

func (s *service) LookupIP(ctx context.Context, ip string) (*ip2region.IPInfo, error) {
	dbReader, err := s.getDBReader(ctx)
	if err != nil {
		return nil, err
	}

	return dbReader.MemorySearch(ip)
}

func (s *service) getDBReader(ctx context.Context) (*ip2region.IPReader, error) {
	s.Lock()
	defer s.Unlock()

	if s.dbReader != nil {
		return s.dbReader, nil
	}

	// 优先从本地文件加载DB文件
	if err := s.loadDBFileFromLocal(); err != nil {
		s.log.Info().Msgf("load ip2region db file from local error, %s, retry other load method ", err)
	} else {
		return s.dbReader, nil
	}

	if err := s.loadDBFileFromBucket(ctx); err != nil {
		s.log.Info().Msgf("load ip2region db file from bucket error, %s", err)
	} else {
		return s.dbReader, nil
	}

	return nil, fmt.Errorf("load ip2region db file error")
}

func (s *service) loadDBFileFromLocal() error {
	file, err := os.Open(s.dbFileName)
	if err != nil {
		return fmt.Errorf("open file error, %s", err)
	}

	reader, err := ip2region.New(file)
	if err != nil {
		return err
	}
	s.dbReader = reader
	return nil
}

func (s *service) loadDBFileFromBucket(ctx context.Context) error {
	buf := bytes.NewBuffer([]byte{})
	downloadReq := oss.NewDownloadFileRequest(s.bucketName, s.dbFileName, buf)
	if err := s.oss.Download(ctx, downloadReq); err != nil {
		return err
	}

	reader, err := ip2region.New(buf)
	if err != nil {
		return err
	}
	s.dbReader = reader

	return nil
}
