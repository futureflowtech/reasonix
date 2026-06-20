// Package tool exposes the Tool abstraction and Registry from internal/tool so
// external embedders can implement custom tools and register them without
// importing internal packages.
package tool

import itool "reasonix/internal/tool"

// Tool is a capability the model can invoke.
type Tool = itool.Tool

// Previewer is an optional capability a writer Tool may implement.
type Previewer = itool.Previewer

// ProgressFunc is a callback for live tool output streaming.
type ProgressFunc = itool.ProgressFunc

// WithProgress attaches a ProgressFunc to a context.
var WithProgress = itool.WithProgress

// ProgressFrom retrieves a ProgressFunc from a context.
var ProgressFrom = itool.ProgressFrom

// Registry is a per-run set of tools.
type Registry = itool.Registry

// NewRegistry returns an empty registry.
var NewRegistry = itool.NewRegistry

// RegisterBuiltin registers a compile-time built-in tool. Intended for init().
var RegisterBuiltin = itool.RegisterBuiltin

// Builtins returns all registered built-in tools, sorted by name.
var Builtins = itool.Builtins

// LookupBuiltin returns a registered built-in by name.
var LookupBuiltin = itool.LookupBuiltin

// PreviewChange returns the change a writer tool would make for args.
var PreviewChange = itool.PreviewChange

// SplitMCPName splits an MCP-qualified tool name into server and tool parts.
var SplitMCPName = itool.SplitMCPName

// MCPNamePrefix is the prefix for MCP-qualified tool names.
const MCPNamePrefix = itool.MCPNamePrefix
