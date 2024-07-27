package pkg

import (
	"io/ioutil"
	"os"
	"testing"
)

// Helper function to create a temporary file with specific content
func createTempFile(content string) (string, error) {
	file, err := ioutil.TempFile("", "testfile_*.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return "", err
	}

	return file.Name(), nil
}

func TestGetLine(t *testing.T) {
	tests := []struct {
		content      string
		lineNumber   int
		expectedLine string
		shouldError  bool
	}{
		{
			content:      "line 1\nline 2\nline 3\n",
			lineNumber:   2,
			expectedLine: "line 2",
			shouldError:  false,
		},
		{
			content:      "line 1\nline 2\nline 3\n",
			lineNumber:   4,
			expectedLine: "",
			shouldError:  false,
		},
		{
			content:      "",
			lineNumber:   1,
			expectedLine: "",
			shouldError:  false,
		},
	}

	for _, tt := range tests {
		tempFileName, err := createTempFile(tt.content)
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tempFileName)

		result := GetLine(tempFileName, tt.lineNumber)

		if result != tt.expectedLine {
			t.Errorf("For file content %q and line number %d, expected %q but got %q", tt.content, tt.lineNumber, tt.expectedLine, result)
		}
	}
}

// TestWriteToOutputFile tests the WriteToOutputFile function
func TestWriteToOutputFile(t *testing.T) {
    tests := []struct {
        filename     string
        content      string
        expectError  bool
        expectedContent string
    }{
        {
            filename:      "testfile1.txt",
            content:       "Hello, World!",
            expectError:   false,
            expectedContent: "Hello, World!",
        },
        {
            filename:      "testfile2.txt",
            content:       "Another test content.",
            expectError:   false,
            expectedContent: "Another test content.",
        },
        {
            filename:      "", // Invalid filename
            content:       "Content should not be written.",
            expectError:   true,
            expectedContent: "",
        },
    }

    for _, tt := range tests {
        err := WriteToOutputFile(tt.filename, tt.content)
        
        if (err != nil) != tt.expectError {
            t.Errorf("WriteToOutputFile() error = %v, expectError %v", err, tt.expectError)
            continue
        }

        if !tt.expectError {
            content, readErr := ioutil.ReadFile(tt.filename)
            if readErr != nil {
                t.Fatalf("Failed to read file %s: %v", tt.filename, readErr)
            }
            if string(content) != tt.expectedContent {
                t.Errorf("For file %s, expected content %q but got %q", tt.filename, tt.expectedContent, string(content))
            }
            os.Remove(tt.filename) // Clean up
        }
    }
}