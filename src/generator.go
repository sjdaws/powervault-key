package main

import (
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
		err = deviceLicence.AddFeature(capabilities.HighPerformanceTier, 0)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.PhysicalDiskSlots192, 0)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.Snapshot, 256)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.SnapshotVirtualDisks, 0)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.SsdCache, 0)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.VirtualDiskCopy, 0)
		if err != nil {
			return err
		}
	case "MD3400":
		err = deviceLicence.AddFeature(capabilities.HighPerformanceTier, 0)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.PhysicalDiskSlots192, 0)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.Snapshot, 256)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.SsdCache, 0)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.VirtualDiskCopy, 0)
		if err != nil {
			return err
		}
	case "MD3460":
		err = deviceLicence.AddFeature(capabilities.DataAssurance, 0)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.PhysicalDiskSlots180, 0)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.Snapshot, 512)
		if err != nil {
			return err
		}

		err = deviceLicence.AddFeature(capabilities.VirtualDiskCopy, 0)
		if err != nil {
			return err
		}
	default:
		err = deviceLicence.AddFeature(capabilities.None, 0)
		if err != nil {
			return err
		}
	}

	err = deviceLicence.Save(output)
	if err != nil {
		return err
	}

	return nil
}
