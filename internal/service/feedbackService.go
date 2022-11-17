package service

import (
	"feedback/internal/database"
	"feedback/internal/types"
	"feedback/pkg/logging"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FeedbackService struct {
	db            database.Feedback
	imagesDirPath string
	log           *logging.Logger
}

var ErrInvalidData = errors.New("invalid data")

func NewFeedbackService(db database.Feedback, imagesDirPath string) *FeedbackService {
	return &FeedbackService{db: db, imagesDirPath: imagesDirPath}
}

func (f *FeedbackService) GetAllCountries() ([]types.Countries, error) {
	return f.db.GetAllCountries()
}

func (f *FeedbackService) GetCountryCities(countryId int) ([]types.Cities, error) {
	return f.db.GetCountryCities(countryId)
}

func (f *FeedbackService) CreateFeedback(feedback *types.Feedbacks) (int, error) {
	return f.db.CreateFeedback(feedback)
}

func (f *FeedbackService) SaveImage(reader *multipart.Form, feedback *types.Feedbacks) (*types.Feedbacks, error) {
	log := logging.GetLogger()
	files := reader.File["image"]
	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if err := f.ValidateImage(files[i].Size); err != nil {
			log.Println(err)
			return nil, err
		}
		path := filepath.Join(f.imagesDirPath, files[i].Filename)
		feedback.Photo = append(feedback.Photo, path)
		dst, err := os.Create(path)
		defer dst.Close()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if _, err := io.Copy(dst, file); err != nil {
			return nil, err
		}
	}
	return feedback, nil
}
func (f *FeedbackService) ValidateImage(size int64) error {
	if size > 2_000_000_000 {
		f.log.Info(ErrInvalidData)
		return ErrInvalidData
	}
	return nil
}
