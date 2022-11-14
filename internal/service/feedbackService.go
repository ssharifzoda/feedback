package service

import (
	"feedback/internal/database"
	"feedback/internal/types"
	"feedback/pkg/logging"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
)

type FeedbackService struct {
	db            database.Feedback
	imagesDirPath string
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

func (f *FeedbackService) CreateFeedback(feedback *types.FeedBacks) error {
	return f.db.CreateFeedback(feedback)
}

func (f *FeedbackService) SaveImage(file io.Reader, fileName string, feedback *types.FeedBacks) (*types.FeedBacks, error) {
	lg := logging.GetLogger()
	path := filepath.Join(f.imagesDirPath, fileName)
	imageFile, err := os.Create(path)
	if err != nil {
		lg.Println(err)
		return nil, err
	}
	defer imageFile.Close()

	_, err = io.Copy(imageFile, file)
	if err != nil {
		lg.Println(err)
		return nil, err
	}
	feedback.Photo = path
	return feedback, nil
}
func (f *FeedbackService) ValidateImage(size int64) error {
	if size > 2_000_000_000 {
		return ErrInvalidData
	}
	return nil
}
