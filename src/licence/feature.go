package licence

import (
	"bufio"
	"crypto/sha1"

	"github.com/sjdaws/powervault-reverse-engineering/capabilities"
)

// feature represents a feature in a licence
type feature struct {
	capability               capabilities.Capability
	digest                   []byte
	duration                 int
	keyVersion               int
	limit                    int
	supportedFeatureBundleId int
	uniqueId                 []byte
}

// CreateFeature creates a new feature
func (l *Licence) createFeature(capability capabilities.Capability, limit int) *feature {
	f := &feature{
		capability: capability,
		limit:      limit,
		uniqueId:   make([]byte, 16),
	}

	hash := sha1.New()
	hash.Write(l.salt)
	hash.Write(l.featureEnableIdentifier)

	if limit > 0 {
		f.keyVersion = 2

		hash.Write(f.uniqueId)
		hash.Write(bigEndian(f.supportedFeatureBundleId))
		hash.Write(bigEndian(int(f.capability)))
		hash.Write(bigEndian(f.limit))
		hash.Write(bigEndian(f.duration))
	} else {
		f.keyVersion = 1

		hash.Write(bigEndian(int(f.capability)))
	}

	digest := sha1.New()
	digest.Write(l.featureEnableIdentifier)
	digest.Write(hash.Sum(nil))

	f.digest = digest.Sum(nil)

	return f
}

// save saves the feature to a writer
func (f *feature) save(featureEnableIdentifier []byte, writer *bufio.Writer) error {
	_, err := writer.Write(bigEndian(f.keyVersion))
	if err != nil {
		return err
	}

	_, err = writer.Write(f.uniqueId)
	if err != nil {
		return err
	}

	_, err = writer.Write(bigEndian(f.supportedFeatureBundleId))
	if err != nil {
		return err
	}

	_, err = writer.Write(bigEndian(int(f.capability)))
	if err != nil {
		return err
	}

	_, err = writer.Write(bigEndian(f.limit))
	if err != nil {
		return err
	}

	_, err = writer.Write(bigEndian(f.duration))
	if err != nil {
		return err
	}

	_, err = writer.Write(featureEnableIdentifier)
	if err != nil {
		return err
	}

	_, err = writer.Write(f.digest)
	if err != nil {
		return err
	}

	return nil
}
