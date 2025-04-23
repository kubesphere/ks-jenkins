package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/kubesphere/ks-jenkins/pkg/types"
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
)

func downloadUpdateCenterJSON(url, filePath string) error {
	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP response errors
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %+v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Strip the JavaScript function wrapper
	raw := string(body)
	prefix := "updateCenter.post("
	suffix := ");"

	if !strings.HasPrefix(raw, prefix) || !strings.HasSuffix(raw, suffix) {
		return fmt.Errorf("unexpected format")
	}

	jsonStr := raw[len(prefix) : len(raw)-len(suffix)]

	return os.WriteFile(filePath, []byte(jsonStr), 0755)
}

func parseFormulaFile(filePath string) (*types.BundleConfig, error) {
	bundleConfig := &types.BundleConfig{}
	fileFormula, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.NewDecoder(fileFormula).Decode(bundleConfig)
	if err != nil {
		return nil, err
	}
	return bundleConfig, nil
}

func parseJenkinsUpdateCenterFile(filePath string) (*types.JenkinsUpdateCenter, error) {
	jenkinsUpdateCenter := &types.JenkinsUpdateCenter{}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, jenkinsUpdateCenter)
	if err != nil {
		return nil, err
	}
	return jenkinsUpdateCenter, nil
}

func generateFormula(updateCenterDownload bool, updateCenterURL, updateCenterJson, formulaTemplateYaml, formulaOverrideYaml, formulaYaml string) error {
	var err error

	if updateCenterDownload {
		err = downloadUpdateCenterJSON(updateCenterURL, updateCenterJson)
		if err != nil {
			return err
		}
	}

	updateCenter, err := parseJenkinsUpdateCenterFile(updateCenterJson)
	if err != nil {
		return err
	}

	bundleConfig, err := parseFormulaFile(formulaTemplateYaml)
	if err != nil {
		return err
	}

	bundleConfigOverride, err := parseFormulaFile(formulaOverrideYaml)
	if err != nil {
		return err
	}

	pluginsMap := make(map[string]types.Plugin)
	for _, p := range bundleConfig.Plugins {
		ps, err := updateCenter.GetPlugin(p.ArtifactID, true)
		if err != nil {
			klog.ErrorS(err, "get plugin failed", "plugin", p.ArtifactID)
			//plugins = append(plugins, p)
			continue
		}

		for _, pl := range ps {
			pluginsMap[pl.ArtifactID] = pl
		}
	}

	for _, p := range bundleConfigOverride.Plugins {
		pluginsMap[p.ArtifactID] = p
	}

	var plugins []types.Plugin
	for _, v := range pluginsMap {
		plugins = append(plugins, v)
	}

	bundleConfig.Plugins = plugins

	err = os.Remove(formulaYaml)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	newFormulaFile, err := os.OpenFile(formulaYaml, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	yamlEncoder := yaml.NewEncoder(newFormulaFile)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(bundleConfig)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var updateCenterURL string
	var updateCenterJson string
	var updateCenterDownload bool
	var formulaTemplateYaml string
	var formulaOverrideYaml string
	var formulaYaml string
	flag.BoolVar(&updateCenterDownload, "update-center-download", false, "whether download the latest update-center.json, you needs provide one existing file by flag 'update-center-json' if false")
	flag.StringVar(&updateCenterURL, "update-center-url", "https://updates.jenkins.io/current/update-center.json", "the url to download update-center.json")
	flag.StringVar(&updateCenterJson, "update-center-json", "update-center.json", "the path of update-center.json to download or use")
	flag.StringVar(&formulaTemplateYaml, "formula-template", "formula-template.yaml", "the path of existing formula-template.yaml")
	flag.StringVar(&formulaOverrideYaml, "formula-override", "formula-override.yaml", "the path of existing formula-override.yaml")
	flag.StringVar(&formulaYaml, "formula", "formula.yaml", "the path of formula.yaml to save")
	err := generateFormula(updateCenterDownload, updateCenterURL, updateCenterJson, formulaTemplateYaml, formulaOverrideYaml, formulaYaml)
	if err != nil {
		klog.Fatalln(err)
	}
}
