// Package diff exposes file-change types from internal/diff so external
// embedders that implement custom tool.Previewer can reference the Change type
// without importing internal packages.
package diff

import idiff "reasonix/internal/diff"

// Kind classifies a change: add, remove, or modify.
type Kind = idiff.Kind

// Change describes what a writer tool would do to a file.
type Change = idiff.Change
