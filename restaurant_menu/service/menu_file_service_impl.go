package service

import (
	"fmt"

	"github.com/proyectos01-a/restaurantMenu/dto/req"
	"github.com/proyectos01-a/restaurantMenu/dto/res"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/models"
	"github.com/sirupsen/logrus"
)

type MenuFileServiceImpl struct {
	menuFileRepo data.MenuFileRepository
	s3Repo       data.S3FileRepository
}

// GetMenuFilesURLByRestaurantID implements MenuFileService.
func (m *MenuFileServiceImpl) GetMenuFilesURLByRestaurantID(restaurantID uint) ([]string, error) {

	// Get details of menu files
	menuFiles, err := m.menuFileRepo.GetMenuFileByRestaurantID(restaurantID)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to get menu files")
		return nil, fmt.Errorf("failed to get menu files: %v", err)
	}

	// Map response
	var menuFileURLs []string
	for _, menuFile := range menuFiles {
		fileURL, err := m.s3Repo.GetFileURL(menuFile.FileName, 120)
		if err != nil {
			logrus.WithError(err).Error("[MenuFileServiceImpl] failed to get file URL")
			return nil, fmt.Errorf("failed to get file URL: %v", err)
		}
		menuFileURLs = append(menuFileURLs, fileURL)
	}

	return menuFileURLs, nil
}

// CreateMenuFile implements MenuFileService.
func (m *MenuFileServiceImpl) CreateMenuFile(fileReq *req.CreateMenuFileReq, fileBytes []byte) (*res.MenuFileResponse, error) {

	if fileReq == nil {
		logrus.Error("[MenuFileServiceImpl] fileReq is nil")
		return nil, fmt.Errorf("invalid request")
	}

	// Upload file to S3
	s3Path, err := m.s3Repo.UploadFile(fileReq.FileName, fileBytes, fileReq.FileType)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to upload file to S3")
		return nil, fmt.Errorf("failed to upload file to S3: %v", err)
	}

	// Create model for menu file
	menuFile := &models.MenuFile{
		FileName:     fileReq.FileName,
		FilePath:     s3Path,
		FileType:     fileReq.FileType,
		FileSize:     int64(len(fileBytes)),
		RestaurantID: fileReq.RestaurantID,
	}

	// Save menu file in database
	menuFile, err = m.menuFileRepo.SaveMenuFile(menuFile)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to save menu file")
		return nil, fmt.Errorf("failed to save menu file: %v", err)
	}

	// Map response
	return &res.MenuFileResponse{
		ID:           menuFile.ID,
		FileName:     menuFile.FileName,
		FilePath:     menuFile.FilePath,
		FileType:     menuFile.FileType,
		FileSize:     menuFile.FileSize,
		RestaurantID: menuFile.RestaurantID,
	}, nil
}

// DeleteMenuFile implements MenuFileService.
func (m *MenuFileServiceImpl) DeleteMenuFile(fileID uint) error {
	// Get details of menu file
	menuFile, err := m.menuFileRepo.GetMenuFileByID(fileID)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to get menu file")
		return fmt.Errorf("failed to get menu file: %v", err)
	}

	// Delete file from S3
	err = m.s3Repo.DeleteFile(menuFile.FilePath)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to delete file from S3")
		return fmt.Errorf("failed to delete file from S3: %v", err)
	}

	// Delete menu file from database
	err = m.menuFileRepo.DeleteMenuFile(fileID)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to delete menu file")
		return fmt.Errorf("failed to delete menu file: %v", err)
	}

	return nil
}

// GetMenuFileByID implements MenuFileService.
func (m *MenuFileServiceImpl) GetMenuFileByID(fileID uint) (*res.MenuFileResponse, error) {

	// Get details of menu file
	menuFile, err := m.menuFileRepo.GetMenuFileByID(fileID)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to get menu file")
		return nil, fmt.Errorf("failed to get menu file: %v", err)
	}

	// Map response
	return &res.MenuFileResponse{
		ID:           menuFile.ID,
		FileName:     menuFile.FileName,
		FilePath:     menuFile.FilePath,
		FileType:     menuFile.FileType,
		FileSize:     menuFile.FileSize,
		RestaurantID: menuFile.RestaurantID,
	}, nil
}

// GetMenuFileByRestaurantID implements MenuFileService.
func (m *MenuFileServiceImpl) GetMenuFileByRestaurantID(restaurantID uint) ([]res.MenuFileResponse, error) {

	// Get details of menu files
	menuFiles, err := m.menuFileRepo.GetMenuFileByRestaurantID(restaurantID)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to get menu files")
		return nil, fmt.Errorf("failed to get menu files: %v", err)
	}

	// Map response
	var menuFileResponses []res.MenuFileResponse
	for _, menuFile := range menuFiles {
		menuFileResponses = append(menuFileResponses, res.MenuFileResponse{
			ID:           menuFile.ID,
			FileName:     menuFile.FileName,
			FilePath:     menuFile.FilePath,
			FileType:     menuFile.FileType,
			FileSize:     menuFile.FileSize,
			RestaurantID: menuFile.RestaurantID,
		})
	}

	return menuFileResponses, nil
}

// UpdateMenuFile implements MenuFileService.
func (m *MenuFileServiceImpl) UpdateMenuFile(fileID uint, fileReq *req.CreateMenuFileReq, fileBytes []byte) (*res.MenuFileResponse, error) {

	if fileReq == nil {
		logrus.Error("[MenuFileServiceImpl] fileReq is nil")
		return nil, fmt.Errorf("invalid request")
	}

	// Get details of menu file
	menuFile, err := m.menuFileRepo.GetMenuFileByID(fileID)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to get menu file")
		return nil, fmt.Errorf("failed to get menu file: %v", err)
	}

	// Delete the old file from S3
	err = m.s3Repo.DeleteFile(menuFile.FilePath)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to delete file from S3")
		return nil, fmt.Errorf("failed to delete file from S3: %v", err)
	}

	// Upload new file to S3
	s3Path, err := m.s3Repo.UploadFile(fileReq.FileName, fileBytes, fileReq.FileType)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to upload file to S3")
		return nil, fmt.Errorf("failed to upload file to S3: %v", err)
	}

	// Update model for menu file
	menuFile.FileName = fileReq.FileName
	menuFile.FilePath = s3Path
	menuFile.FileType = fileReq.FileType
	menuFile.FileSize = int64(len(fileBytes))
	menuFile.RestaurantID = fileReq.RestaurantID

	// Update menu file in database
	menuFile, err = m.menuFileRepo.UpdateMenuFile(menuFile)
	if err != nil {
		logrus.WithError(err).Error("[MenuFileServiceImpl] failed to update menu file")
		return nil, fmt.Errorf("failed to update menu file: %v", err)
	}

	// Map response
	return &res.MenuFileResponse{
		ID:           menuFile.ID,
		FileName:     menuFile.FileName,
		FilePath:     menuFile.FilePath,
		FileType:     menuFile.FileType,
		FileSize:     menuFile.FileSize,
		RestaurantID: menuFile.RestaurantID,
	}, nil
}

func NewMenuFileService(menuFileRepo data.MenuFileRepository, s3Repo data.S3FileRepository) MenuFileService {
	return &MenuFileServiceImpl{
		menuFileRepo: menuFileRepo,
		s3Repo:       s3Repo,
	}
}
