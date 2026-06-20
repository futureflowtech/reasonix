// Package control exposes the transport-agnostic Controller and its driving
// interfaces from internal/control so external embedders can drive sessions
// without importing internal packages.
package control

import ictrl "reasonix/internal/control"

// Controller wraps the internal Controller. All exported methods are promoted
// so embedders call them directly (ctrl.Run(...), ctrl.Submit(...), etc.).
type Controller struct {
	*ictrl.Controller
}

// --- Sub-port interfaces (all type aliases — identical to internal) ---

// Lifecycle covers a session's identity and lifecycle.
type Lifecycle = ictrl.Lifecycle

// TurnControl covers driving a model turn and observing its run state.
type TurnControl = ictrl.TurnControl

// Approvals covers tool-approval and ask prompts.
type Approvals = ictrl.Approvals

// Goals covers the active-goal FSM and plan mode.
type Goals = ictrl.Goals

// SessionHistory covers checkpoint/rewind, branch/fork, compact/summarize.
type SessionHistory = ictrl.SessionHistory

// MemoryControl covers session/project memory reads and mutations.
type MemoryControl = ictrl.MemoryControl

// Capabilities covers MCP servers, skills, commands, hooks.
type Capabilities = ictrl.Capabilities

// Status covers run/usage/billing telemetry.
type Status = ictrl.Status

// SessionPersistence covers snapshotting a session and tear-down.
type SessionPersistence = ictrl.SessionPersistence

// Input covers composing a turn's text and resolving @-references.
type Input = ictrl.Input

// Settings covers runtime session settings.
type Settings = ictrl.Settings

// SessionAPI is the full driving port — the composition of every sub-port.
type SessionAPI = ictrl.SessionAPI

// --- Struct & const aliases ---

// Options carries the already-built pieces setup assembles.
type Options = ictrl.Options

// RuntimeStatus is the frontend-facing snapshot of foreground turn state.
type RuntimeStatus = ictrl.RuntimeStatus

// RememberResult describes what happened when an approval rule was persisted.
type RememberResult = ictrl.RememberResult

// RewindScope selects what a Rewind restores.
type RewindScope = ictrl.RewindScope

const (
	RewindCode         = ictrl.RewindCode
	RewindConversation = ictrl.RewindConversation
	RewindBoth         = ictrl.RewindBoth
)

// ToolApproval mode constants.
const (
	ToolApprovalAsk  = ictrl.ToolApprovalAsk
	ToolApprovalAuto = ictrl.ToolApprovalAuto
	ToolApprovalYolo = ictrl.ToolApprovalYolo
)

// Goal status constants.
const (
	GoalStatusRunning  = ictrl.GoalStatusRunning
	GoalStatusComplete = ictrl.GoalStatusComplete
	GoalStatusBlocked  = ictrl.GoalStatusBlocked
	GoalStatusStopped  = ictrl.GoalStatusStopped
)

// GoalResearchMode selects the research posture for a goal.
type GoalResearchMode = ictrl.GoalResearchMode

const (
	GoalResearchAuto GoalResearchMode = iota
	GoalResearchOn
	GoalResearchOff
)

// GoalCommandAction describes a parsed goal command.
type GoalCommandAction = ictrl.GoalCommandAction

const (
	GoalCommandStatus = ictrl.GoalCommandStatus
	GoalCommandSet    = ictrl.GoalCommandSet
	GoalCommandClear  = ictrl.GoalCommandClear
)

// GoalCommand is a parsed /goal command.
type GoalCommand = ictrl.GoalCommand

// ParseGoalCommand parses a /goal command from user input.
var ParseGoalCommand = ictrl.ParseGoalCommand

// PlanModeMarker is the text prefix that gates execution in plan mode.
const PlanModeMarker = ictrl.PlanModeMarker

// StripComposePrefixes removes controller-injected prefixes from composed text.
var StripComposePrefixes = ictrl.StripComposePrefixes

// IsSyntheticUserMessage reports whether content is a synthetic user message.
var IsSyntheticUserMessage = ictrl.IsSyntheticUserMessage

// SessionDestroyHandle is an opaque handle for async session teardown.
type SessionDestroyHandle = ictrl.SessionDestroyHandle

// ToolResultData holds the data returned by ToolResult().
type ToolResultData = ictrl.ToolResultData

// SlashItem describes a slash-command completion item.
type SlashItem = ictrl.SlashItem

// ArgData holds session data for slash-command argument completions.
type ArgData = ictrl.ArgData

// SlashArgItems returns slash-command argument completions.
var SlashArgItems = ictrl.SlashArgItems

// ErrTurnRunning reports that a second foreground turn was started.
var ErrTurnRunning = ictrl.ErrTurnRunning

// ReconcileCleanupPending retries physical cleanup for logically removed sessions.
var ReconcileCleanupPending = ictrl.ReconcileCleanupPending

// New builds a Controller. Normally called by boot.Build; embedders can call it
// directly for manual assembly.
func New(opts Options) *Controller {
	return &Controller{ictrl.New(opts)}
}
