package parser

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig("../resources/sampleconfig.yml")
	if err != nil {
		t.Errorf("Error : %v", err)
	}

	fmt.Println(cfg)
}

func TestSaveConfig(t *testing.T) {
	cfg, err := LoadConfig("../resources/sampleconfig.yml")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	cfg.Name = "Modified Container"

	err = SaveConfig("../resources/sampleconfig1.yml", cfg)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}
