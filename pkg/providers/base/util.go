package base

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

func GetImagesForFileBasedProvider(providerType string, e types.Environment) ([]types.Image, error) {
	var ret []types.Image
	imagesDir := viper.GetString(common.ViperImagesDir)
	listing, _ := os.ReadDir(filepath.Join(imagesDir, providerType))
	for _, dirEntry := range listing {
		ret = append(ret, models.NewImage(filepath.Join(providerType, dirEntry.Name()), dirEntry.Name(), time.Now(), e))
	}
	return ret, nil
}

func DestroyImageForFileBasedProvider(imageID string) error {
	imagesDir := viper.GetString(common.ViperImagesDir)

	imagePath := filepath.Join(imagesDir, imageID)
	return os.RemoveAll(imagePath)

}