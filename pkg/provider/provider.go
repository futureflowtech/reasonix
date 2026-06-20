// Package provider exposes chat-capable model backend types from
// internal/provider so external embedders can implement Provider and
// construct/config flows without importing internal packages.
package provider

import iprovider "reasonix/internal/provider"

// Role is the role of a message.
type Role = iprovider.Role

const (
	RoleSystem    = iprovider.RoleSystem
	RoleUser      = iprovider.RoleUser
	RoleAssistant = iprovider.RoleAssistant
	RoleTool      = iprovider.RoleTool
)

// Message is a single conversation message.
type Message = iprovider.Message

// ToolCall is a tool invocation requested by the model.
type ToolCall = iprovider.ToolCall

// ToolSchema is a tool definition exposed to the model.
type ToolSchema = iprovider.ToolSchema

// Request is a single completion request.
type Request = iprovider.Request

// ChunkType identifies the kind of a streamed increment.
type ChunkType = iprovider.ChunkType

const (
	ChunkText          = iprovider.ChunkText
	ChunkReasoning     = iprovider.ChunkReasoning
	ChunkToolCallStart = iprovider.ChunkToolCallStart
	ChunkToolCall      = iprovider.ChunkToolCall
	ChunkUsage         = iprovider.ChunkUsage
	ChunkDone          = iprovider.ChunkDone
	ChunkError         = iprovider.ChunkError
)

// Usage reports token accounting for a completion.
type Usage = iprovider.Usage

// Pricing is a provider's per-1M-token rates.
type Pricing = iprovider.Pricing

// Chunk is a single streamed event.
type Chunk = iprovider.Chunk

// Provider is a chat-capable model backend.
type Provider = iprovider.Provider

// Config is a resolved provider instance configuration.
type Config = iprovider.Config

// AuthError reports that a provider rejected the API key.
type AuthError = iprovider.AuthError

// StreamInterruptedError marks a recoverable transport cut.
type StreamInterruptedError = iprovider.StreamInterruptedError

// IsStreamInterrupted reports whether err is a *StreamInterruptedError.
var IsStreamInterrupted = iprovider.IsStreamInterrupted

// Factory builds a Provider from a resolved Config.
type Factory = iprovider.Factory

// Register adds a factory under a kind (e.g. "openai"). Intended for init().
var Register = iprovider.Register

// New instantiates the provider of the given kind.
var New = iprovider.New

// Kinds returns the registered provider kinds.
var Kinds = iprovider.Kinds

// NormalizeMessages repairs a conversation history for the wire.
var NormalizeMessages = iprovider.NormalizeMessages

// NormalizeSessionMessages applies safe-to-persist repairs.
var NormalizeSessionMessages = iprovider.NormalizeSessionMessages

// SanitizeToolPairing is the provider-side alias for NormalizeMessages.
var SanitizeToolPairing = iprovider.SanitizeToolPairing

// ParseImageDataURL splits a data URL into media type and base64 payload.
var ParseImageDataURL = iprovider.ParseImageDataURL

// CanonicalizeSchema normalizes JSON Schema for prefix-cache stability.
var CanonicalizeSchema = iprovider.CanonicalizeSchema

// MaxRetries is the maximum number of retry attempts.
const MaxRetries = iprovider.MaxRetries

// RetryInfo describes a retry attempt.
type RetryInfo = iprovider.RetryInfo

// RetryNotify is a callback for retry notifications.
type RetryNotify = iprovider.RetryNotify

// SendOptions configures a single SendWithRetry attempt.
type SendOptions = iprovider.SendOptions

// APIError is an HTTP-level error from a provider.
type APIError = iprovider.APIError

// RetryableStatus reports whether an HTTP status is retryable.
var RetryableStatus = iprovider.RetryableStatus

// IsConnReset reports whether err is a connection-reset error.
var IsConnReset = iprovider.IsConnReset

// WithRetryNotify attaches a RetryNotify to a context.
var WithRetryNotify = iprovider.WithRetryNotify
