import { createFileRoute, Link } from "@tanstack/react-router";
import { ArrowLeft, BookOpen } from "lucide-react";

export const Route = createFileRoute("/docs/llm-guideline")({
  component: LlmGuidelinePage,
});

function LlmGuidelinePage() {
  return (
    <div className="min-h-screen bg-background text-foreground">
      <div className="border-b bg-muted/40">
        <div className="mx-auto max-w-4xl px-6 py-8">
          <Link
            to="/docs"
            className="inline-flex items-center gap-1.5 text-sm text-muted-foreground hover:text-foreground transition-colors"
          >
            <ArrowLeft className="h-4 w-4" />
            Back to docs
          </Link>
          <h1 className="mt-4 text-3xl font-bold tracking-tight flex items-center gap-2">
            <BookOpen className="h-7 w-7 text-primary" />
            LLM Guideline
          </h1>
          <p className="mt-2 text-muted-foreground">
            Guidelines for AI agents working with the errorwrapper-v3 codebase.
          </p>
        </div>
      </div>

      <article className="mx-auto max-w-4xl px-6 py-12 prose prose-neutral dark:prose-invert">
        <h2>Framework Overview</h2>
        <p>
          errorwrapper-v3 is a Go library for structured error handling. It provides error creation,
          wrapping, type classification, conversion, verification, functional utilities, command
          execution, and typed results.
        </p>

        <h2>Package Layout</h2>
        <pre className="rounded-md bg-muted p-4 overflow-x-auto">
          <code className="text-sm font-mono">
{`root/
  errnew/           Error constructors
  errwrappers/      Collection, chaining, JSON serialization
  errtype/          Error kind enum and classification
  errconv/          Conversion between error representations
  errverify/        Testing utilities and streaming verifier
  errdefer/         Defer-pattern helpers
  errfunc/          Functional utilities over errors
  errdata/          Typed result packages
    erranygen/      Generic Result[T] (preferred for new code)
    errstr/         String result (frozen, legacy)
    errbool/        Bool result (frozen, legacy)
    errint/         Int result (frozen, legacy)
    errfloat/       Float32 result (frozen, legacy)
    errfloat64/     Float64 result (frozen, legacy)
    errbyte/        Byte result (frozen, legacy)
    errany/         Any result (frozen, legacy)
    errjson/        JSON helpers (frozen, legacy)
    errcasted/      Casting helpers (frozen, legacy)
  errcmd/           Command execution wrappers
  errcmdportable/   Edge-safe command execution
    osadapter/      os/exec adapter (separate subpackage)
    errcmdbridge/   errcmd -> portable converter
  linuxservicecmd/  Linux service command wrappers
  trydo/            Try/do pattern
  eithererr/        Either monad for errors
  refs/             Reference helpers for pointer-safe workflows
  internal/         Internal packages — NOT for external use
  tests/            Test suites
    integratedtests/  Integration test packages`}
          </code>
        </pre>

        <h2>Naming Conventions</h2>
        <ul>
          <li>
            <strong>Files</strong>: PascalCase for exported, camelCase for internal
          </li>
          <li>
            <strong>Packages</strong>: lowercase, no underscores
          </li>
          <li>
            <strong>Tests</strong>: <code>*_test.go</code> in same directory or{" "}
            <code>tests/integratedtests/&lt;pkg&gt;tests/</code>
          </li>
          <li>
            <strong>Interfaces</strong>: Named with <code>-er</code> suffix where possible (
            <code>ErrorWrapper</code>)
          </li>
        </ul>

        <h2>Testing Rules</h2>
        <table>
          <thead>
            <tr>
              <th>Rule</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Public packages</td>
              <td>MUST have tests under tests/integratedtests/&lt;pkg&gt;tests/</td>
            </tr>
            <tr>
              <td>Internal packages</td>
              <td>
                <strong>NEVER</strong> write tests for internal/* packages
              </td>
            </tr>
            <tr>
              <td>Framework</td>
              <td>Prefer goconvey (Convey/So) for BDD-style tests</td>
            </tr>
            <tr>
              <td>Coverage</td>
              <td>Run via run.ps1 -tc or run.sh -tc</td>
            </tr>
          </tbody>
        </table>
        <p>
          <strong>CRITICAL</strong>: Do NOT create test files for any package under{" "}
          <code>internal/</code>.
        </p>

        <h2>Anti-Patterns</h2>
        <ol>
          <li>
            <strong>Don&apos;t import internal/* from outside the module</strong> — these are
            intentionally unexported
          </li>
          <li>
            <strong>Don&apos;t modify frozen errdata/* packages</strong> — use erranygen for new
            generic code
          </li>
          <li>
            <strong>Don&apos;t use os/exec directly in edge-targeted code</strong> — use
            errcmdportable
          </li>
          <li>
            <strong>Don&apos;t break JSON shape of errwrappers.Collection</strong> — downstream
            consumers depend on it
          </li>
        </ol>

        <h2>Edge Runtime Compatibility</h2>
        <p>When modifying command execution code:</p>
        <ul>
          <li>
            <code>errcmdportable/</code> is safe for edge bundlers
          </li>
          <li>
            <code>errcmdportable/osadapter/</code> contains real os/exec — only import if running
            on real OS
          </li>
          <li>
            <code>errcmdportable/errcmdbridge/</code> converts <code>*errcmd.Result</code> to
            portable <code>Result</code>
          </li>
          <li>
            Default <code>Detect()</code> returns <code>NoProcessRunner</code> on unknown/edge
            targets
          </li>
        </ul>
      </article>

      <footer className="border-t bg-muted/40">
        <div className="mx-auto max-w-4xl px-6 py-8 text-sm text-muted-foreground text-center">
          errorwrapper-v3 — MIT License
        </div>
      </footer>
    </div>
  );
}
