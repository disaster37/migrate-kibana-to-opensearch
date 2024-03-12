package migrate

import (
	"bufio"
	"bytes"

	"emperror.dev/errors"
)

func SplitNewLineByteBuffer(datas *bytes.Buffer) (dataSplited [][]byte, err error) {
	scanner := bufio.NewScanner(datas)

	dataSplited = make([][]byte, 0)

	for scanner.Scan() {
		dataSplited = append(dataSplited, scanner.Bytes())
	}

	if err = scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "Error when split exported objects")
	}

	// Remove last line that content what is export
	if len(dataSplited) > 0 {
		dataSplited = dataSplited[:len(dataSplited)-1]
	}

	return dataSplited, nil
}
