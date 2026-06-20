// Package boot exposes the one-shot Build function from internal/boot so
// external embedders can assemble a ready-to-drive Controller from configuration
// without importing internal packages. By default Build reads config from
// ./reasonix.toml / ~/.reasonix/config.toml on disk, same as the CLI; an
// embedder that wants to supply a config in memory instead (built from
// pkg/config.Default() and its own overrides) sets Options.Config, and
// Options.APIKeyOverride to hand through an already-resolved credential
// without writing it to disk — see both fields' doc comments.
package boot

import (
	"context"
	"os"

	iboot "reasonix/internal/boot"
	iconfig "reasonix/internal/config"
	iplugin "reasonix/internal/plugin"
	"reasonix/pkg/control"

	// Blank-import providers and builtins so registries are populated
	// for anyone that imports this package.
	_ "reasonix/internal/provider/anthropic"
	_ "reasonix/internal/provider/openai"
	_ "reasonix/internal/tool/builtin"
)

// Options carries the per-run knobs a frontend chooses; everything else is read
// from configuration.
type Options = iboot.Options

// PluginSpec declares an external MCP server for Options.ExtraPlugins —
// see internal/plugin.Spec for field docs. Exported here so embedders can
// construct entries without reaching into an internal package.
type PluginSpec = iplugin.Spec

// HTTPPlugin builds a PluginSpec for an MCP server reachable over Streamable
// HTTP — e.g. an in-process server the embedder itself exposes, for
// session-scoped custom tools that don't warrant a subprocess.
func HTTPPlugin(name, url string, headers map[string]string) PluginSpec {
	return PluginSpec{Name: name, Type: "http", URL: url, Headers: headers}
}

// ErrUnknownModel is returned by Build when the configured model can't be
// resolved to a provider.
var ErrUnknownModel = iboot.ErrUnknownModel

// Build loads config, resolves the model(s), and returns a Controller wrapping a
// single Agent, or a two-model Coordinator when agent.planner_model is set.
func Build(ctx context.Context, opts Options) (*control.Controller, error) {
	c, err := iboot.Build(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &control.Controller{Controller: c}, nil
}

// ModelInfo describes one configured model preset — the same registry
// Build resolves Options.Model against and that ErrUnknownModel's error
// message already summarizes by name. Exposed publicly so an embedder (e.g.
// a server wanting to answer "what models can I offer") doesn't have to
// parse that error string.
type ModelInfo struct {
	// Name is the preset name to pass as Options.Model (e.g. "deepseek-flash"
	// — NOT a raw provider model slug, which is Model/Models below).
	Name string
	// Model is the underlying provider-side model id/slug this preset
	// resolves to by default (e.g. "deepseek-v4-flash") — ProviderEntry's
	// explicit `default`, else the first of Models. A config using the
	// older single-`model` field (rather than `models`+`default`) surfaces
	// the same way here.
	Model string
	// Models lists every model slug this provider entry offers (a config
	// may bundle several models — e.g. one entry per API key/base_url —
	// behind one preset Name); Model above is Models' effective default.
	Models []string
	// Kind is the provider protocol family (e.g. "openai", "anthropic").
	Kind string
	// ContextWindow is the configured max context size in tokens, 0 if unset.
	ContextWindow int
	// HasCredentials reports whether the API key this preset needs is
	// actually resolvable from its configured env var right now — a preset
	// can be defined but not currently usable if the key isn't set.
	HasCredentials bool
}

// AvailableModels returns every model preset configured for workspaceRoot
// (resolution order: flag > ./reasonix.toml > ~/.reasonix/config.toml >
// built-in defaults — same as Build), regardless of whether credentials are
// currently resolvable for each (see ModelInfo.HasCredentials).
func AvailableModels(workspaceRoot string) ([]ModelInfo, error) {
	cfg, err := iconfig.LoadForRoot(workspaceRoot)
	if err != nil {
		return nil, err
	}
	out := make([]ModelInfo, 0, len(cfg.Providers))
	for _, p := range cfg.Providers {
		models := p.ModelList()
		model := p.DefaultModel()
		if model == "" {
			model = p.Model // back-compat single-`model` config, if ModelList/DefaultModel found nothing
		}
		out = append(out, ModelInfo{
			Name:           p.Name,
			Model:          model,
			Models:         models,
			Kind:           p.Kind,
			ContextWindow:  p.ContextWindow,
			HasCredentials: p.APIKeyEnv == "" || os.Getenv(p.APIKeyEnv) != "",
		})
	}
	return out, nil
}

// NewProvider builds a provider.Provider from a configured entry.
var NewProvider = iboot.NewProvider

// NewProviderWithProxy builds a provider.Provider with configured proxy settings.
var NewProviderWithProxy = iboot.NewProviderWithProxy

// PluginSpecs maps configured plugin entries to plugin.Spec.
var PluginSpecs = iboot.PluginSpecs

// PluginSpecsForRoot maps configured plugin entries to plugin.Spec with
// workspace-aware compatibility overrides.
var PluginSpecsForRoot = iboot.PluginSpecsForRoot

// MCPStartupNotice formats the warning when MCP servers failed to connect.
var MCPStartupNotice = iboot.MCPStartupNotice

// LSPSpecs returns the language → server map.
var LSPSpecs = iboot.LSPSpecs
