// Package agent exposes the core Agent, Runner, Session, and plugin interfaces
// from internal/agent so external embedders can construct and drive an agent
// without importing internal packages.
package agent

import iagent "reasonix/internal/agent"

// Runner carries out one task turn.
type Runner = iagent.Runner

// Renderer redraws the assistant's final-answer text as styled output.
type Renderer = iagent.Renderer

// Asker puts structured multiple-choice questions to the user.
type Asker = iagent.Asker

// Gate decides, per tool call, whether it may run.
type Gate = iagent.Gate

// ToolHooks fires user-configured shell hooks around each tool call.
type ToolHooks = iagent.ToolHooks

// KeepPolicy is a bitmask selecting which messages to keep across compaction.
type KeepPolicy = iagent.KeepPolicy

const (
	KeepErrors     = iagent.KeepErrors
	KeepUserMarked = iagent.KeepUserMarked
)

// Agent drives a single task: a Provider, a tool Registry, and a Session
// wired into the main loop.
type Agent = iagent.Agent

// Session holds the conversation history for one task.
type Session = iagent.Session

// Options carries per-agent configuration.
type Options = iagent.Options

// New constructs an Agent.
var New = iagent.New

// NewSession initializes a session with an optional system prompt.
var NewSession = iagent.NewSession

// Coordinator runs two models in separate sessions.
type Coordinator = iagent.Coordinator

// NewCoordinator wires a planner provider to an executor.
var NewCoordinator = iagent.NewCoordinator

// DefaultPlannerPrompt steers the planner toward concise plans, not execution.
const DefaultPlannerPrompt = iagent.DefaultPlannerPrompt

// PlannerPromptWithContext appends cache-stable standing context.
var PlannerPromptWithContext = iagent.PlannerPromptWithContext

// HandoffTask extracts the original user task from a handoff message.
var HandoffTask = iagent.HandoffTask

// MidTurnSteerPrefix is the prefix prepended to mid-turn steer messages.
const MidTurnSteerPrefix = iagent.MidTurnSteerPrefix

// CallContext returns the executing call's ID, sink, and asker.
var CallContext = iagent.CallContext

// PlanModeFromContext reports whether the tool call is executing under the
// agent's read-only planning gate.
var PlanModeFromContext = iagent.PlanModeFromContext

// WithParentSession stamps the active parent session ID onto a turn context.
var WithParentSession = iagent.WithParentSession

// ParentSession returns the active parent session ID.
var ParentSession = iagent.ParentSession

// WithUserImages carries the data URLs of images the user attached to this turn.
var WithUserImages = iagent.WithUserImages

// BranchMeta is metadata about a conversation branch.
type BranchMeta = iagent.BranchMeta

// BranchInfo describes a saved branch.
type BranchInfo = iagent.BranchInfo

// BranchMetaCountsVersion is the current schema version for BranchMeta.
const BranchMetaCountsVersion = iagent.BranchMetaCountsVersion
