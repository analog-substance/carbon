package base

import (
	"github.com/analog-substance/carbon/pkg/common"
	"github.com/analog-substance/carbon/pkg/models"
	"github.com/analog-substance/carbon/pkg/types"
	"os"
	"path/filepath"
	"time"
)

func GetImagesForFileBasedProvider(providerType string, e types.Environment) ([]types.Image, error) {
	var ret []types.Image
	listing, _ := os.ReadDir(filepath.Join(common.ImagesDir(), providerType))
	for _, dirEntry := range listing {
		ret = append(ret, models.NewImage(filepath.Join(providerType, dirEntry.Name()), dirEntry.Name(), time.Now(), e))
	}
	return ret, nil
}

func DestroyImageForFileBasedProvider(imageID string) error {
	imagePath := filepath.Join(common.ImagesDir(), imageID)
	return os.RemoveAll(imagePath)
}
