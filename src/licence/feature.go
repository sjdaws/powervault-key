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
func (l *Licence) createFeature(capability capabilities.Capability, limit int) (*feature, error) {
	f := &feature{
		capability: capability,
		limit:      limit,
		uniqueId:   make([]byte, 16),
	}

	hash := sha1.New()

	_, err := hash.Write(l.salt)
	if err != nil {
		return nil, err
	}

	_, err = hash.Write(l.featureEnableIdentifier)
	if err != nil {
		return nil, err
	}

	if limit > 0 {
		f.keyVersion = 2

		_, err = hash.Write(f.uniqueId)
		if err != nil {
			return nil, err
		}

		_, err = hash.Write(bigEndian(f.supportedFeatureBundleId))
		if err != nil {
			return nil, err
		}

		_, err = hash.Write(bigEndian(int(f.capability)))
		if err != nil {
			return nil, err
		}

		_, err = hash.Write(bigEndian(f.limit))
		if err != nil {
			return nil, err
		}

		_, err = hash.Write(bigEndian(f.duration))
		if err != nil {
			return nil, err
		}
	} else {
		f.keyVersion = 1

		_, err = hash.Write(bigEndian(int(f.capability)))
		if err != nil {
			return nil, err
		}
	}

	digest := sha1.New()

	_, err = digest.Write(l.featureEnableIdentifier)
	if err != nil {
		return nil, err
	}

	_, err = digest.Write(hash.Sum(nil))
	if err != nil {
		return nil, err
	}

	f.digest = digest.Sum(nil)

	return f, nil
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
