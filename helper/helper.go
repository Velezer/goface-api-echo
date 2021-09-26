package helper

import (
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kagami/go-face"
)

const BaseDir = "."

var (
	ModelDir    = filepath.Join(BaseDir, "dat_models")
	EncodedDir  = filepath.Join(BaseDir, "encoded")
)

func RecognizeFile(rec *face.Recognizer, folder string, filename string) ([]face.Face, error) {
	knownFaces, err := rec.RecognizeFile(filepath.Join(folder, filename))
	os.Remove(filepath.Join(folder, filename))
	if err != nil {
		return nil, err
	}
	if len(knownFaces) > 1 {
		err = errors.New("detected more than one face, only support one face")
		return nil, err
	}
	if len(knownFaces) == 0 {
		err = errors.New("no face detected")
		return nil, err
	}
	return knownFaces, err
}

func SaveFile(dir string, filename string, content multipart.File) error {
	os.MkdirAll(dir, os.ModeDir)

	destination, err := os.Create(filepath.Join(dir, filename))
	defer destination.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(destination, content)
	if err != nil {
		return err
	}

	return nil
}

func OSReadDir(root string, extension string) (dirs []string, files []string) {
	f, _ := os.Open(root)
	defer f.Close()

	fileInfo, _ := f.Readdir(-1)
	for _, file := range fileInfo {
		if strings.Contains(file.Name(), extension) {
			files = append(files, file.Name())
		}
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return
}

func DumpToJson(dir string, filename string, content face.Descriptor) {
	os.Mkdir(dir, os.ModeDir)
	file, _ := os.Create(filepath.Join(dir, filename))
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(content)

}

func DecodeFromJson(dir string, filename string) (content face.Descriptor) {
	file, _ := os.Open(filepath.Join(dir, filename))
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&content)

	return content
}
