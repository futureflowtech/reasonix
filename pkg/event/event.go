// Package event exposes the typed event stream types from internal/event so
// external embedders can implement event.Sink and receive agent events without
// importing internal packages.
package event

import ievent "reasonix/internal/event"

// Kind tags an Event.
type Kind = ievent.Kind

const (
	TurnStarted       = ievent.TurnStarted
	Reasoning         = ievent.Reasoning
	Text              = ievent.Text
	Message           = ievent.Message
	ToolDispatch      = ievent.ToolDispatch
	ToolResult        = ievent.ToolResult
	Usage             = ievent.Usage
	Notice            = ievent.Notice
	Phase             = ievent.Phase
	ApprovalRequest   = ievent.ApprovalRequest
	AskRequest        = ievent.AskRequest
	TurnDone          = ievent.TurnDone
	CompactionStarted = ievent.CompactionStarted
	CompactionDone    = ievent.CompactionDone
	ToolProgress      = ievent.ToolProgress
	MCPSurfaceReady   = ievent.MCPSurfaceReady
	Retrying          = ievent.Retrying
	Steer             = ievent.Steer
)

// Level classifies a Notice.
type Level = ievent.Level

const (
	LevelInfo = ievent.LevelInfo
	LevelWarn = ievent.LevelWarn
)

// Sink consumes a turn's events.
type Sink = ievent.Sink

// FuncSink adapts a plain function to a Sink.
type FuncSink = ievent.FuncSink

// Discard is a Sink that drops every event.
var Discard = ievent.Discard

// ReadinessAuditSink is an optional sink capability.
type ReadinessAuditSink = ievent.ReadinessAuditSink

// RecordReadinessAudit forwards a readiness audit receipt to sinks that opt in.
var RecordReadinessAudit = ievent.RecordReadinessAudit

// Sync wraps a Sink for safe concurrent emission.
var Sync = ievent.Sync

// Event is one increment in a turn's event stream.
type Event = ievent.Event

// Tool describes a tool call for ToolDispatch / ToolResult events.
type Tool = ievent.Tool

// FileDiff is a previewed change.
type FileDiff = ievent.FileDiff

// Approval identifies a pending tool-call approval.
type Approval = ievent.Approval

// AskOption is one choice for an AskQuestion.
type AskOption = ievent.AskOption

// AskQuestion is one structured question.
type AskQuestion = ievent.AskQuestion

// Ask carries an AskRequest.
type Ask = ievent.Ask

// AskAnswer is the user's reply to one AskQuestion.
type AskAnswer = ievent.AskAnswer

// Compaction carries a context-compaction pass.
type Compaction = ievent.Compaction

// CacheDiagnostics describes cache-prefix change attribution.
type CacheDiagnostics = ievent.CacheDiagnostics

// Profile carries the subagent model/effort resolved for this call.
type Profile = ievent.Profile
