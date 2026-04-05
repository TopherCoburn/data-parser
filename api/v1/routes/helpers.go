package data_parser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

// FileExists checks if a file exists at the given path.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ReadFileContent reads the content of a file into a string.
func ReadFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filePath, err)
	}
	return string(content), nil
}

// WriteToFile writes data to a file.
func WriteToFile(filePath string, data string) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %w", filePath, err)
	}
	return nil
}

// ParseJSON parses a JSON string into an interface{}.
func ParseJSON(jsonString string) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}
	return data, nil
}

// ConvertToFloat64 attempts to convert a value to a float64.
func ConvertToFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		// Try to parse the string as a float64.
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot convert string '%s' to float64: %w", v, err)
		}
		return f, nil
	default:
		// If it's another type, attempt to convert it using reflection.
		val := reflect.ValueOf(value)
		if val.CanConvert(reflect.TypeOf(float64(0))) {
			return val.Convert(reflect.TypeOf(float64(0))).Float(), nil
		}
		return 0, fmt.Errorf("cannot convert type %T to float64", value)
	}
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

// SanitizeFilename removes invalid characters from a filename.
func SanitizeFilename(filename string) string {
	filename = strings.ReplaceAll(filename, " ", "_") // Replace spaces with underscores
	filename = strings.ReplaceAll(filename, "/", "-") // Replace slashes with dashes
	filename = strings.ReplaceAll(filename, "\\", "-") // Replace backslashes with dashes
	filename = strings.ReplaceAll(filename, ":", "-") // Replace colons with dashes
	filename = strings.ReplaceAll(filename, "*", "")  // Remove asterisks
	filename = strings.ReplaceAll(filename, "?", "")  // Remove question marks
	filename = strings.ReplaceAll(filename, "\"", "") // Remove double quotes
	filename = strings.ReplaceAll(filename, "<", "")  // Remove less than signs
	filename = strings.ReplaceAll(filename, ">", "")  // Remove greater than signs
	filename = strings.ReplaceAll(filename, "|", "")  // Remove pipes
	return filename
}