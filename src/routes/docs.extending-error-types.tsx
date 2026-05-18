import { createFileRoute, Link } from "@tanstack/react-router";
import { ArrowLeft, BookOpen } from "lucide-react";

export const Route = createFileRoute("/docs/extending-error-types")({
  component: ExtendingErrorTypesPage,
});

function ExtendingErrorTypesPage() {
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
            Extending Error Types
          </h1>
          <p className="mt-2 text-muted-foreground">
            Three approaches for extending errorwrapper-v3 with custom error types and behaviors.
          </p>
        </div>
      </div>

      <article className="mx-auto max-w-4xl px-6 py-12 prose prose-neutral dark:prose-invert">
        <h2>Approach 1: Interface-Based Extension</h2>
        <p>
          Implement the <code>ErrorWrapper</code> interface directly for full control over error
          representation.
        </p>
        <pre className="rounded-md bg-muted p-4 overflow-x-auto">
          <code className="text-sm font-mono">
{`type CustomError struct {
    Code    int
    Message string
    Context map[string]interface{}
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *CustomError) Unwrap() error {
    // Return underlying error if wrapped
    return nil
}`}
          </code>
        </pre>
        <p>
          Use with <code>errwrappers.New()</code> to participate in collections and JSON
          serialization.
        </p>

        <h2>Approach 2: Registry-Based Extension</h2>
        <p>
          Register custom error kinds with the global type registry so they integrate with{" "}
          <code>errtype</code> classification.
        </p>
        <pre className="rounded-md bg-muted p-4 overflow-x-auto">
          <code className="text-sm font-mono">
{`import "github.com/alimtvnetwork/errorwrapper-v3/errtype"

var MyCustomKind = errtype.RegisterKind(
    errtype.KindConfig{
        Name:        "CustomDomainError",
        Code:        900,
        Description: "Error in custom business domain",
    },
)`}
          </code>
        </pre>
        <p>
          Registered kinds can then be used with <code>errnew</code> constructors:
        </p>
        <pre className="rounded-md bg-muted p-4 overflow-x-auto">
          <code className="text-sm font-mono">
{`err := errnew.Message(MyCustomKind, "something went wrong")`}
          </code>
        </pre>

        <h2>Approach 3: Context-Based Extension</h2>
        <p>
          Attach structured context to errors using <code>errnew.Refs</code> and{" "}
          <code>HandleErrorWithRefs</code>:
        </p>
        <pre className="rounded-md bg-muted p-4 overflow-x-auto">
          <code className="text-sm font-mono">
{`err := errnew.Refs(
    errtype.Validation,
    "userID",    userID,
    "field",     "email",
    ",",         "invalid email format",
)`}
          </code>
        </pre>
        <p>
          This preserves references for downstream logging, tracing, and debugging without
          breaking the error string.
        </p>

        <h2>Migration Notes</h2>
        <ul>
          <li>
            Legacy <code>errdata/*</code> packages (errstr, errbool, errint, etc.) are{" "}
            <strong>frozen</strong>.
          </li>
          <li>
            New code should use <code>errdata/erranygen.Result[T]</code> for generic typed
            results.
          </li>
          <li>See <code>errdata/erranygen</code> for examples of the new generic API.</li>
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
