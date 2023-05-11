package main

import (
	"fmt"

	"github.com/sjdaws/powervault-reverse-engineering/capabilities"
	"github.com/sjdaws/powervault-reverse-engineering/licence"
)

// generate generates and saves the licence key
func generate(device string, featureEnableIdentifier string, output string) error {
	deviceLicence, err := licence.New(featureEnableIdentifier)
	if err != nil {
		return err
	}

	switch device {
	case "MD3200":
		fallthrough
	case "MD3220":
		deviceLicence.AddFeature(capabilities.HighPerformanceTier, 0)
		deviceLicence.AddFeature(capabilities.PhysicalDiskSlots192, 0)
		deviceLicence.AddFeature(capabilities.Snapshot, 256)
		deviceLicence.AddFeature(capabilities.SnapshotVirtualDisks, 0)
		deviceLicence.AddFeature(capabilities.SsdCache, 0)
		deviceLicence.AddFeature(capabilities.VirtualDiskCopy, 0)
	case "MD3460":
		deviceLicence.AddFeature(capabilities.DataAssurance, 0)
		deviceLicence.AddFeature(capabilities.PhysicalDiskSlots180, 0)
		deviceLicence.AddFeature(capabilities.Snapshot, 512)
		deviceLicence.AddFeature(capabilities.VirtualDiskCopy, 0)
	default:
		return fmt.Errorf("unsupported array: %s", device)
	}

	err = deviceLicence.Save(output)
	if err != nil {
		return err
	}

	return nil
}
