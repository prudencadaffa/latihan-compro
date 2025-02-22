package storage

import (
	"io"
	"latihan-compro/config"

	"github.com/labstack/gommon/log"
	storage_go "github.com/supabase-community/storage-go"
)

type SupabaseInterface interface {
	// path must be filename ex: /img/photo.jpg
	// file must be io.Reader
	UploadFile(path string, file io.Reader) (string, error)
}

type supabaseStruct struct {
	cfg *config.Config
}

// UploadFile implements SupabaseInterface.
func (s *supabaseStruct) UploadFile(path string, file io.Reader) (string, error) {
	client := storage_go.NewClient(s.cfg.Supabase.StorageUrl, s.cfg.Supabase.StorageKey, map[string]string{"Content-Type": "image/png"})

	_, err := client.UploadFile(s.cfg.Supabase.StorageBucket, path, file)
	if err != nil {
		log.Errorf("Error uploading file: %v", err)
		return "", err
	}

	result := client.GetPublicUrl(s.cfg.Supabase.StorageBucket, path)
	return result.SignedURL, nil
}

func NewSupabase(cfg *config.Config) SupabaseInterface {
	return &supabaseStruct{
		cfg: cfg,
	}
}
