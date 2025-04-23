package types

import (
	"fmt"
	"strings"

	"k8s.io/klog/v2"
)

type JenkinsUpdateCenter struct {
	ConnectionCheckUrl  string                   `json:"connectionCheckUrl,omitempty"`
	Core                Core                     `json:"core,omitempty"`
	Deprecations        map[string]Deprecation   `json:"deprecations,omitempty"`
	GenerationTimestamp string                   `json:"generationTimestamp,omitempty"`
	ID                  string                   `json:"id,omitempty"`
	Plugins             map[string]PluginDetails `json:"plugins,omitempty"`
}

var (
	ErrPluginDeprecated    = fmt.Errorf("plugin deprecated")
	ErrPluginNotFound      = fmt.Errorf("plugin not found")
	ErrPluginHasInvalidGAV = fmt.Errorf("plugin has invalid gav")
)

func (c *JenkinsUpdateCenter) GetPlugin(pluginName string, withDependencies bool) ([]Plugin, error) {
	_, ok := c.Deprecations[pluginName]
	if ok {
		return nil, ErrPluginDeprecated
	}

	plugin, ok := c.Plugins[pluginName]
	if !ok {
		return nil, ErrPluginNotFound
	}

	gavParts := strings.Split(plugin.GAV, ":")
	if len(gavParts) != 3 {
		return nil, ErrPluginHasInvalidGAV
	}

	var plugins []Plugin

	p := Plugin{
		GroupID:    gavParts[0],
		ArtifactID: gavParts[1],
		Source: Source{
			Version: gavParts[2],
		},
	}

	plugins = append(plugins, p)

	if withDependencies {
		for _, dep := range plugin.Dependencies {
			depPlugin, err := c.GetPlugin(dep.Name, true)
			if err != nil {
				klog.ErrorS(err, "get plugin dependency failed", "dependent-plugin", dep.Name, "plugin", pluginName)
				continue
			}
			plugins = append(plugins, depPlugin...)
		}
	}

	return plugins, nil
}

type Core struct {
	BuildDate string `json:"buildDate,omitempty"`
	Name      string `json:"name,omitempty"`
	Sha1      string `json:"sha1,omitempty"`
	Sha256    string `json:"sha256,omitempty"`
	Size      int    `json:"size,omitempty"`
	URL       string `json:"url,omitempty"`
	Version   string `json:"version,omitempty"`
}

type Deprecation struct {
	URL string `json:"url,omitempty"`
}

type PluginDetails struct {
	BuildDate         string         `json:"buildDate,omitempty"`
	DefaultBranch     string         `json:"defaultBranch,omitempty,omitempty"`
	Dependencies      []Dependency   `json:"dependencies,omitempty"`
	Developers        []Developer    `json:"developers,omitempty"`
	Excerpt           string         `json:"excerpt,omitempty"`
	GAV               string         `json:"gav,omitempty"`
	IssueTrackers     []IssueTracker `json:"issueTrackers,omitempty"`
	Labels            []string       `json:"labels,omitempty"`
	Name              string         `json:"name,omitempty"`
	Popularity        int            `json:"popularity,omitempty"`
	PreviousTimestamp string         `json:"previousTimestamp,omitempty"`
	PreviousVersion   string         `json:"previousVersion,omitempty"`
	ReleaseTimestamp  string         `json:"releaseTimestamp,omitempty"`
	RequiredCore      string         `json:"requiredCore,omitempty"`
	SCM               string         `json:"scm,omitempty"`
	Sha1              string         `json:"sha1,omitempty"`
	Sha256            string         `json:"sha256,omitempty"`
	Size              int            `json:"size,omitempty"`
	Title             string         `json:"title,omitempty"`
	URL               string         `json:"url,omitempty"`
	Version           string         `json:"version,omitempty"`
	Wiki              string         `json:"wiki,omitempty"`
}

type Dependency struct {
	Name     string `json:"name,omitempty"`
	Optional bool   `json:"optional,omitempty"`
	Version  string `json:"version,omitempty"`
}

type Developer struct {
	DeveloperID string `json:"developerId,omitempty"`
	Name        string `json:"name,omitempty"`
}

type IssueTracker struct {
	ReportUrl string `json:"reportUrl,omitempty"`
	Type      string `json:"type,omitempty"`
	ViewUrl   string `json:"viewUrl,omitempty"`
}
