// Package ovirt contains ovirt-specific Terraform-variable logic.
package ovirt

import (
	"encoding/json"

	"github.com/openshift/installer/pkg/rhcos"
	"github.com/openshift/installer/pkg/tfvars/internal/cache"
)

type config struct {
	URL                    string `json:"ovirt_url"`
	Username               string `json:"ovirt_username"`
	Password               string `json:"ovirt_password"`
	Cafile                 string `json:"ovirt_cafile,omitempty"`
	ClusterID              string `json:"ovirt_cluster_id"`
	TemplateID             string `json:"ovirt_template_id"`
	BaseImageName          string `json:"openstack_base_image_name,omitempty"`
	BaseImageLocalFilePath string `json:"openstack_base_image_local_file_path,omitempty"`
}

// TFVars generates ovirt-specific Terraform variables.
func TFVars(
	engineURL string,
	engineUser string,
	enginePass string,
	engineCafile string,
	clusterID string,
	templateID string,
	baseImageName string,
	infraID string) ([]byte, error) {

	cfg := config{
		URL:                    engineURL,
		Username:               engineUser,
		Password:               enginePass,
		Cafile:                 engineCafile,
		ClusterID:              clusterID,
		TemplateID:             templateID,
		BaseImageName:          baseImageName,
	}

	imageName, isURL := rhcos.GenerateOpenStackImageName(baseImageName, infraID)
	cfg.BaseImageName = imageName
	if isURL {
		imageFilePath, err := cache.DownloadImageFile(imageName)
		if err != nil {
			return nil, err
		}
		cfg.BaseImageLocalFilePath = imageFilePath
	}

	return json.MarshalIndent(cfg, "", "  ")
}
