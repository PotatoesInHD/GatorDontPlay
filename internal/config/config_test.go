package config

import "testing"
import "path/filepath"
import "os"


func TestConfig(t *testing.T) {
	t.Run("Get file path", func(t *testing.T) {
		path, err := getFilePath()
		if err != nil {
			t.Fatalf("getting actual file path: %v", err)
		}
		expected, err := os.UserHomeDir()
		if err != nil {
			t.Fatalf("getting expected file path: %v", err)
		}
		expected = filepath.Join(expected, configFileName)
		if path != expected {
			t.Errorf("Expected: %v Got: %v", expected, path)
		}
	})

	t.Run("read and set user", func(t *testing.T) {
		//fake home is used for test environment and will change back after test
		fakeHome := t.TempDir()
		t.Setenv("HOME", fakeHome)
		//writes a testfile into temp directory "fake home"
		testPath, err := os.UserHomeDir()
		if err != nil {
			t.Fatalf("getting test file path: %v", err)
		}
		testPath = filepath.Join(testPath, configFileName)
		testBytes := []byte(`{"db_url": "postgres://example"}`)
		if err := os.WriteFile(testPath, testBytes, 0600); err != nil {
			t.Fatalf("writing test file: %v", err)
		}
		//its actually using funcs on the test file in fakehome temp dir
		cfg, err := Read()
		if err != nil {
			t.Fatalf("reading file: %v", err)
		}

		if err = cfg.SetUser("bobs"); err != nil {
			t.Fatalf("setting user: %v", err)
		}

		cfg, err = Read()
		if err != nil {
			t.Fatalf("reading file: %v", err)
		}

		expected := Config{DbUrl: "postgres://example", CurrentUserName: "bobs"}
		if cfg != expected {
			t.Errorf("Expected: %v Got: %v", expected, cfg)
		}
	})
}
