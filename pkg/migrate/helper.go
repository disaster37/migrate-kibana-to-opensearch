package migrate

import (
	"bufio"
	"bytes"

	"emperror.dev/errors"
	log "github.com/sirupsen/logrus"
)

func SplitNewLineByteBuffer(datas *bytes.Buffer) (dataSplited [][]byte, err error) {
	scanner := bufio.NewScanner(datas)

	dataSplited = make([][]byte, 0)

	for scanner.Scan() {
		dataSplited = append(dataSplited, []byte(scanner.Text()))
	}

	if err = scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "Error when split exported objects")
	}

	log.Infof("Found %d exported objects", len(dataSplited))

	return dataSplited, nil
}
