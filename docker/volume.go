package docker

import (
	"context"
	"os/exec"

	"github.com/docker/docker/api/types"
	volumeTypes "github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/philips-software/gautocloud-connectors/hsdp"
)

type Volume struct {
	ClusterID     string
	Client        *client.Client
	EnableFluentd bool
}

type VolumeSpec struct {
	StoragePath   string `json:"storagePath"`
	ContainerPath string `json:"containerPath"`
	Output        bool   `json:"output"`
}

func (v *Volume) VolumeCreate(ctx context.Context, name string, labels map[string]string) (types.Volume, error) {
	var vol types.Volume
	var err error
	vol, err = v.Client.VolumeCreate(ctx, volumeTypes.VolumeCreateBody{
		Driver:     "local",
		Name:       name,
		DriverOpts: map[string]string{},
		Labels:     labels,
	})
	if err != nil {
		return vol, err
	}
	return vol, nil
}

func (v *Volume) VolumeRemove(ctx context.Context, volumeID string) error {
	err := v.Client.VolumeRemove(ctx, volumeID, true)
	if err != nil {
		return err
	}
	return nil
}

func (v *Volume) mirror(ctx context.Context, s3creds hsdp.S3Credentials, volumeName string, toS3 bool, spec VolumeSpec) error {
	minioVolumeName := "minio-" + volumeName
	// Create Minio volumes
	vol, err := v.VolumeCreate(ctx, minioVolumeName, map[string]string{"created_by": "containerh-agent"})
	if err != nil {
		return err
	}
	defer func() { _ = v.VolumeRemove(ctx, vol.Name) }()

	// Configure Minio
	cmd := exec.Command("docker",
		"run", "--rm", "-v", minioVolumeName+":/root", "minio/mc", "alias", "set", "s3", "https://"+s3creds.Endpoint, s3creds.APIKey, s3creds.SecretKey)
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Mirror
	source := "s3/" + s3creds.Bucket + spec.StoragePath
	dest := spec.ContainerPath
	if toS3 {
		dest = source
		source = spec.ContainerPath
	}
	args := []string{
		"run",
		"--rm",
		"-v", minioVolumeName + ":/root",
		"-v", volumeName + ":" + spec.ContainerPath,
	}
	if v.EnableFluentd {
		args = append(args, "--log-driver=fluentd")
	}
	args = append(args, "minio/mc", "mirror", "--overwrite", source, dest)
	cmd = exec.Command("docker", args...)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (v *Volume) MirrorFromS3(ctx context.Context, s3creds hsdp.S3Credentials, volumeName string, spec VolumeSpec) error {
	return v.mirror(ctx, s3creds, volumeName, false, spec)
}

func (v *Volume) MirrorToS3(ctx context.Context, s3creds hsdp.S3Credentials, volumeName string, spec VolumeSpec) error {
	return v.mirror(ctx, s3creds, volumeName, true, spec)
}
