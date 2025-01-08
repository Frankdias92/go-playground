package errorhandler

import (
	"errors"
	"io"
	"os"
	"strings"
)

func Reader() {
	str := "hello, word\n"
	reader := strings.NewReader(str)
	writer := MyWriter{os.Stdout}

	// io.ReadAll(reader)
	// io.ReadFull(reader, nil)

	buffer := make([]byte, 2)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		// fmt.Println(n)
		// fmt.Println(buffer[:n])

		writer.Write(buffer[:n])
	}
}

type MyWriter struct {
	w io.Writer
}

func (mw MyWriter) Write(b []byte) (int, error) {
	for i, bb := range b {
		b[i] = bb + 10
	}
	return mw.w.Write(b)
}
