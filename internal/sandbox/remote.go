package sandbox

import "context"

// RemoteExecutor runs a shell command in an out-of-process sandbox — e.g. a
// Kubernetes Pod running under a Kata Containers RuntimeClass — instead of
// the local OS-sandbox machinery the rest of this package wraps. When
// ConfineBash is given one, the bash tool routes every command through it
// and skips local exec entirely (no OS sandbox wrapping, no background-job
// support, no host-terminal echo — none of that applies once the command
// isn't running on this process's own host at all).
//
// Deliberately shaped identically to flow's own wasm.ShellExecutor and
// flow/sandbox.Handle.Exec (cmd/args/dir in, stdout/stderr/exitCode/err
// out) so a flow/sandbox.Handle satisfies this interface structurally, with
// no glue code, from the embedder's side (agent/native.Backend.SetSandbox).
type RemoteExecutor interface {
	Exec(ctx context.Context, cmd string, args []string, dir string) (stdout, stderr string, exitCode int, err error)
}
