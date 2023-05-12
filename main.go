package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver"
)

func main() {
	input := os.Args
	if len(input) < 2 {
		fmt.Println("Informe o nome da pasta a ser processada")
		fmt.Println("Exemplo: unzip ./processamento")
		fmt.Println("Onde: ./processamento é a pasta onde estão os arquivos ZIP")
		fmt.Println("Caso esteja no windows utilize a barra invertida")
		return
	}
	dir := input[1]
	dest := "./processado"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if filepath.Ext(path) == ".zip" {
				fmt.Printf("Extraindo arquivo ZIP: %s\n", path)
				err = extractZipFile(path, dest)
				if err != nil {
					return err
				}
			}
			if filepath.Ext(path) == ".rar" {
				fmt.Printf("Extraindo arquivo RAR: %s\n", path)
				err = extractRarFile(path, dest)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Erro ao percorrer diretórios: %s\n", err)
	}
}

func extractZipFile(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		f := file.Name
		f = strings.Replace(f, `\`, "/", -1)
		f = strings.Replace(f, ":", "_", -1)
		p := strings.Split(f, "/")
		filename := p[len(p)-1]
		var path string
		switch {
		case filepath.Ext(filename) == ".xml":
			path = filepath.Join(dest, "xml", filename)
		case filepath.Ext(filename) == ".txt":
			path = filepath.Join(dest, "txt", filename)
		default:
			path = filepath.Join(dest, "outros", filename)
		}

		if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		srcFile, err := file.Open()
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(path)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func extractRarFile(rarFile, dest string) error {
	rar := archiver.NewRar()
	err := rar.Walk(rarFile, func(f archiver.File) error {
		if f.IsDir() {
			return nil
		}
		fname := f.Name()
		fname = strings.Replace(fname, `\`, "/", -1)
		fname = strings.Replace(fname, ":", "_", -1)
		p := strings.Split(fname, "/")
		filename := p[len(p)-1]
		var path string
		switch {
		case filepath.Ext(filename) == ".xml":
			path = filepath.Join(dest, "xml", filename)
		case filepath.Ext(filename) == ".txt":
			path = filepath.Join(dest, "txt", filename)
		default:
			path = filepath.Join(dest, "outros", filename)
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		destFile, err := os.Create(path)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, f)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
