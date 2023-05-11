package licence

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"os"

	"github.com/sjdaws/powervault-reverse-engineering/capabilities"
)

// Licence represents a licence for a device
type Licence struct {
	featureEnableIdentifier []byte
	features                []*feature
	fileVersion             int
	preamble                []byte
	salt                    []byte
	signature               string
}

// New returns a new instance of a licence
func New(featureEnableIdentifier string) (*Licence, error) {
	decodedfeatureEnableIdentifier, err := hex.DecodeString(featureEnableIdentifier)
	if err != nil {
		return nil, err
	}

	return &Licence{
		featureEnableIdentifier: decodedfeatureEnableIdentifier,
		features:                []*feature{},
		fileVersion:             2,
		// Single byte uint8 endian representation of PREAMBLE from devmgr.v1125api15.sam.scriptengine.FeatureKeyDecoder
		// private static final byte[] PREAMBLE = new byte[] { -84, -19, 0, 5, 116 };
		preamble: []byte{0xAC, 0xED, 0x00, 0x05, 0x74},
		// Single byte uint8 endian representation of salt retrieved by https://github.com/ReverseAllTheThings in
		// https://github.com/ReverseAllTheThings/PowervaultKeygen/blob/master/DellPowerVaultFeatureKeygen/Feature.cs#L31
		salt:      []byte{0x82, 0x64, 0x09, 0x17, 0xB3, 0x5D, 0x7A, 0x9F},
		signature: "FEATURE KEYS",
	}, nil
}

// AddFeature adds a feature to a licence
func (l *Licence) AddFeature(capability capabilities.Capability, limit int) error {
	feature, err := l.createFeature(capability, limit)
	if err != nil {
		return err
	}

	l.features = append(l.features, feature)

	return nil
}

// Save saves the licence to a file
func (l *Licence) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer closeFile(file)

	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	_, err = writer.Write(l.preamble)
	if err != nil {
		return err
	}

	_, err = writer.Write(bigEndian16(len(l.signature)))
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(l.signature))
	if err != nil {
		return err
	}

	_, err = writer.Write(bigEndian(l.fileVersion))
	if err != nil {
		return err
	}

	_, err = writer.Write(bigEndian(len(l.features)))
	if err != nil {
		return err
	}

	for _, feature := range l.features {
		err = feature.save(l.featureEnableIdentifier, writer)
		if err != nil {
			return err
		}
	}

	// Flush to buffer
	err = writer.Flush()
	if err != nil {
		return err
	}

	// Save to file
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// closeFile closes the file ignoring any errors
func closeFile(file *os.File) {
	_ = file.Close()
}
