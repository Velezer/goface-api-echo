package helper

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kagami/go-face"
)

const DataDir = "."

var (
	ModelDir   = filepath.Join(DataDir, "models")
	ImagesDir  = filepath.Join(DataDir, "images")
	EncodedDir = filepath.Join(DataDir, "encoded")
)

func GetSamplesLabels(rec *face.Recognizer) (samples []face.Descriptor, labels []string) {
	dirs, _ := OSReadDir(EncodedDir, ".jpg")
	for _, dir := range dirs {
		encFolder := filepath.Join(EncodedDir, dir)
		_, files := OSReadDir(encFolder, ".jpg")
		for _, file := range files {
			descriptor := DecodeFromJson(encFolder, file)
			
			samples = append(samples, descriptor)
			labels = append(labels, dir)
		}
	}

	return
}

func SaveFile(dir string, filename string, content multipart.File) {
	os.Mkdir(dir, os.ModeDir)
	destination, _ := os.Create(filepath.Join(dir, filename))
	defer destination.Close()
	io.Copy(destination, content)
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

func DumpToJson(dir string, filename string, object face.Descriptor) {
	os.Mkdir(dir, os.ModeDir)
	file, _ := os.Create(filepath.Join(dir, filename))
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(object)

}

func DecodeFromJson(dir string, filename string) (descriptor face.Descriptor) {
	file, _ := os.Open(filepath.Join(dir, filename))
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&descriptor)

	return descriptor
}
