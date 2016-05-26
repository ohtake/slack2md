package main

import (
	"bufio"
	"io"
)

type Translator interface {
	FileNameIndex() string
	FileNameChannel(channelName string) string
	FileNameHistory(channelName string, chunkNumber int) string

	ToHeading(level int, text string) []string
	ToChannelList(channels []Channel) []string
	ToChunkList(chunks []ChunkInfo) []string
	ToMessageList(chunk []MessageResolved) []string
	ToUserTable(users []User) []string
}

type TranslatingWriter struct {
	translator Translator
	writer     *bufio.Writer
}

func NewTranslatingWriter(translator Translator, writer io.Writer) *TranslatingWriter {
	w := new(TranslatingWriter)
	w.translator = translator
	w.writer = bufio.NewWriter(writer)
	return w
}

func (w *TranslatingWriter) writeLines(lines []string) error {
	for _, l := range lines {
		_, err := w.writer.WriteString(l)
		if err != nil {
			return err
		}
		_, err = w.writer.WriteString("\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *TranslatingWriter) WriteHeading(level int, text string) error {
	return w.writeLines(w.translator.ToHeading(level, text))
}
func (w *TranslatingWriter) WriteChannelList(channels []Channel) error {
	return w.writeLines(w.translator.ToChannelList(channels))
}
func (w *TranslatingWriter) WriteChunkList(chunks []ChunkInfo) error {
	return w.writeLines(w.translator.ToChunkList(chunks))
}
func (w *TranslatingWriter) WriteMessageList(chunk []MessageResolved) error {
	return w.writeLines(w.translator.ToMessageList(chunk))
}
func (w *TranslatingWriter) WriteUserTable(users []User) error {
	return w.writeLines(w.translator.ToUserTable(users))
}

func (w *TranslatingWriter) Flush() error {
	return w.writer.Flush()
}
