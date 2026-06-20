// Package config exposes Reasonix's runtime Config type and its constructor
// from internal/config so external embedders can build a config value in
// memory (see pkg/boot's Options.Config) without importing internal
// packages or going through disk-based Load/LoadForRoot.
package config

import iconfig "reasonix/internal/config"

// Config is Reasonix's runtime configuration — see internal/config's own
// doc comment for the full field-by-field TOML schema.
type Config = iconfig.Config

// Default returns the compiled-in default configuration — the same starting
// point config.LoadForRoot itself layers TOML files on top of. An embedder
// building a Config in memory (rather than from a reasonix.toml on disk)
// should start here too, so unset fields fall back to sane defaults instead
// of Go zero values.
var Default = iconfig.Default

// ProviderEntry describes one configured model preset (name, backend kind,
// base URL, credential env var, pricing, ...). Exposed so an embedder can
// declare a custom [[providers]] entry programmatically.
type ProviderEntry = iconfig.ProviderEntry
