package types

type BundleConfig struct {
	Bundle           Bundle            `yaml:"bundle,omitempty"`
	BuildSettings    BuildSettings     `yaml:"buildSettings,omitempty"`
	War              Plugin            `yaml:"war,omitempty"`
	Plugins          []Plugin          `yaml:"plugins,omitempty"`
	SystemProperties map[string]string `yaml:"systemProperties,omitempty"`
	GroovyHooks      []GroovyHook      `yaml:"groovyHooks,omitempty"`
}

type Bundle struct {
	GroupID     string `yaml:"groupId,omitempty"`
	ArtifactID  string `yaml:"artifactId,omitempty"`
	Description string `yaml:"description,omitempty"`
	Vendor      string `yaml:"vendor,omitempty"`
}

type BuildSettings struct {
	Docker DockerSettings `yaml:"docker,omitempty"`
}

type DockerSettings struct {
	Base     string `yaml:"base,omitempty"`
	Tag      string `yaml:"tag,omitempty"`
	Build    bool   `yaml:"build,omitempty"`
	BuildX   bool   `yaml:"buildx,omitempty"`
	Platform string `yaml:"platform,omitempty"`
	Output   string `yaml:"output,omitempty"`
}

type Source struct {
	Version string `yaml:"version,omitempty"`
}

type Plugin struct {
	GroupID    string `yaml:"groupId,omitempty"`
	ArtifactID string `yaml:"artifactId,omitempty"`
	Source     Source `yaml:"source,omitempty"`
}

type GroovyHook struct {
	Type   string     `yaml:"type,omitempty"`
	ID     string     `yaml:"id,omitempty"`
	Source HookSource `yaml:"source,omitempty"`
}

type HookSource struct {
	Dir string `yaml:"dir,omitempty"`
}
