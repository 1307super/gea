package zip

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
)

type ZipUtils struct {
	err       error
	buffer    *bytes.Buffer
	zipWriter *zip.Writer
}

func New(buf *bytes.Buffer) *ZipUtils {
	return &ZipUtils{
		buffer:    buf,
		zipWriter: zip.NewWriter(buf),
	}
}

// PackToBuffer 打包到buufer
func (z *ZipUtils) PackToBuffer(filename string, p []byte) error {
	zipEntry, err := z.zipWriter.Create(filename)
	if err != nil {
		return err
	}
	_, err = zipEntry.Write(p)
	if err != nil {
		return err
	}
	z.Flush()
	return nil
}

// PackToFile 打包并生成文件
func (z *ZipUtils) PackToFile(filename string, p []byte) error {
	zipEntry, err := z.zipWriter.Create(filename)
	if err != nil {
		return err
	}
	_, err = zipEntry.Write(p)
	if err != nil {
		return err
	}
	z.Flush()
	z.Close()
	ioutil.WriteFile(filename, z.buffer.Bytes(), 0644)
	return nil
}

func (z *ZipUtils) Flush() error {
	return z.zipWriter.Flush()
}
func (z *ZipUtils) Close() error {
	return z.zipWriter.Close()
}
