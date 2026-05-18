import { createFileRoute } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Copy, Check, BookOpen, Github, Package, Terminal } from "lucide-react";
import { useState } from "react";

export const Route = createFileRoute("/")({
  component: Index,
});

const packages = [
  { name: "errnew", desc: "Constructor utilities for creating wrapped errors." },
  { name: "errwrappers", desc: "Core error wrapping, collection, and aggregation." },
  { name: "errtype", desc: "Type-safe error kind classification." },
  { name: "errconv", desc: "Convert between error representations and data models." },
  { name: "errverify", desc: "Streaming collection verifier with five match modes." },
  { name: "errdefer", desc: "Defer-pattern helpers for error cleanup." },
  { name: "errfunc", desc: "Functional utilities over error values." },
  { name: "errcmd / errcmdportable", desc: "Command execution with portable, edge-safe runners." },
  { name: "errdata/*", desc: "Typed result generics (erranygen) and legacy typed results." },
  { name: "trydo / eithererr", desc: "Result-monad and either-style error handling." },
  { name: "refs", desc: "Reference helpers for pointer-safe error workflows." },
  { name: "linuxservicecmd", desc: "Linux service command wrappers." },
];

const features = [
  {
    title: "Structured Wrapping",
    body: "Chain, collect, and aggregate errors with full context preservation. Supports nested wrappers and bulk operations.",
  },
  {
    title: "Type-Safe Generics",
    body: "erranygen.Result[T] provides a generic result type that works with any Go type, reducing boilerplate without breaking legacy callers.",
  },
  {
    title: "Streaming Verification",
    body: "Feed collections or channels into a verifier with Equal, Contains, Regex, and fold-aware match modes — no intermediate slices required.",
  },
  {
    title: "Edge-Runtime Safe",
    body: "errcmdportable separates the os/exec adapter into a subpackage so bundlers targeting Workers/edge never pull process-spawning code.",
  },
];

const codeExamples = [
  {
    title: "Create a typed error",
    code: `err := errnew.Message(errtype.NotFound, "resource not found")
if err.IsError() {
    fmt.Println(err.FullString())
}`,
  },
  {
    title: "Collect multiple errors",
    code: `coll := errwrappers.New(5)
coll.AddWrapperPtr(&err1)
coll.AddWrapperPtr(&err2)

// Serialize to JSON
jsonResult := coll.Json()
jsonResult.Unmarshal(&newColl)`,
  },
  {
    title: "Generic result (Go 1.18+)",
    code: `result := erranygen.NewResult[string]()
result.SetValue("hello")
if result.HasError() {
    fmt.Println(result.GetError().Error())
}`,
  },
];

function CopyButton({ text }: { text: string }) {
  const [copied, setCopied] = useState(false);
  return (
    <button
      onClick={() => {
        navigator.clipboard.writeText(text);
        setCopied(true);
        setTimeout(() => setCopied(false), 2000);
      }}
      className="inline-flex items-center gap-1 rounded-md border border-input bg-background px-2 py-1 text-xs text-foreground transition-colors hover:bg-accent"
    >
      {copied ? <Check className="h-3.5 w-3.5 text-green-500" /> : <Copy className="h-3.5 w-3.5" />}
      {copied ? "Copied" : "Copy"}
    </button>
  );
}

function CodeBlock({ title, code }: { title: string; code: string }) {
  return (
    <Card>
      <CardHeader className="pb-2 flex flex-row items-center justify-between">
        <CardTitle className="text-sm font-mono text-muted-foreground">{title}</CardTitle>
        <CopyButton text={code} />
      </CardHeader>
      <CardContent>
        <pre className="overflow-x-auto rounded-md bg-muted p-4 text-sm font-mono text-foreground">
          <code>{code}</code>
        </pre>
      </CardContent>
    </Card>
  );
}

function NavBar() {
  return (
    <nav className="sticky top-0 z-50 border-b bg-background/80 backdrop-blur">
      <div className="mx-auto max-w-5xl px-6 py-3 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <Terminal className="h-5 w-5 text-primary" />
          <span className="font-semibold text-foreground">errorwrapper-v3</span>
        </div>
        <div className="flex items-center gap-4">
          <a
            href="/docs"
            className="flex items-center gap-1.5 text-sm text-muted-foreground hover:text-foreground transition-colors"
          >
            <BookOpen className="h-4 w-4" />
            Docs
          </a>
          <a
            href="https://github.com/alimtvnetwork/errorwrapper-v3"
            target="_blank"
            rel="noopener noreferrer"
            className="flex items-center gap-1.5 text-sm text-muted-foreground hover:text-foreground transition-colors"
          >
            <Github className="h-4 w-4" />
            GitHub
          </a>
        </div>
      </div>
    </nav>
  );
}

function Index() {
  return (
    <div className="min-h-screen bg-background text-foreground">
      <NavBar />

      {/* Hero */}
      <section className="border-b bg-muted/40">
        <div className="mx-auto max-w-5xl px-6 py-20 md:py-28">
          <div className="flex flex-wrap gap-2 mb-6">
            <Badge variant="secondary">Go 1.25+</Badge>
            <Badge variant="outline">v3</Badge>
            <Badge variant="outline">MIT</Badge>
          </div>
          <h1 className="text-4xl md:text-6xl font-bold tracking-tight text-foreground">
            errorwrapper-v3
          </h1>
          <p className="mt-4 text-lg md:text-xl text-muted-foreground max-w-2xl">
            A comprehensive Go library for structured error handling, wrapping,
            verification, and portable command execution.
          </p>
          <div className="mt-8 flex flex-wrap gap-3">
            <Button asChild size="lg">
              <a href="https://github.com/alimtvnetwork/errorwrapper-v3" target="_blank" rel="noopener noreferrer">
                <Github className="mr-2 h-4 w-4" />
                View on GitHub
              </a>
            </Button>
            <Button variant="outline" size="lg" asChild>
              <a href="/docs">
                <BookOpen className="mr-2 h-4 w-4" />
                Read the docs
              </a>
            </Button>
          </div>
        </div>
      </section>

      {/* Quick install */}
      <section className="mx-auto max-w-5xl px-6 py-12">
        <h2 className="text-2xl font-semibold tracking-tight">Quick start</h2>
        <div className="mt-4 flex items-center justify-between rounded-lg border bg-card p-4">
          <code className="font-mono text-sm text-card-foreground overflow-x-auto">
            go get github.com/alimtvnetwork/errorwrapper-v3
          </code>
          <CopyButton text="go get github.com/alimtvnetwork/errorwrapper-v3" />
        </div>
      </section>

      {/* Code examples */}
      <section className="mx-auto max-w-5xl px-6 pb-12">
        <h2 className="text-2xl font-semibold tracking-tight">Usage examples</h2>
        <div className="mt-6 grid gap-4">
          {codeExamples.map((ex) => (
            <CodeBlock key={ex.title} title={ex.title} code={ex.code} />
          ))}
        </div>
      </section>

      {/* Features */}
      <section className="mx-auto max-w-5xl px-6 pb-12">
        <h2 className="text-2xl font-semibold tracking-tight">Features</h2>
        <div className="mt-6 grid gap-4 sm:grid-cols-2">
          {features.map((f) => (
            <Card key={f.title}>
              <CardHeader>
                <CardTitle>{f.title}</CardTitle>
              </CardHeader>
              <CardContent>
                <CardDescription className="text-muted-foreground">
                  {f.body}
                </CardDescription>
              </CardContent>
            </Card>
          ))}
        </div>
      </section>

      {/* Package index */}
      <section className="mx-auto max-w-5xl px-6 pb-16">
        <h2 className="text-2xl font-semibold tracking-tight flex items-center gap-2">
          <Package className="h-5 w-5" />
          Package index
        </h2>
        <div className="mt-6 grid gap-3 sm:grid-cols-2 lg:grid-cols-3">
          {packages.map((pkg) => (
            <Card key={pkg.name} className="hover:shadow-md transition-shadow">
              <CardHeader className="pb-3">
                <CardTitle className="text-base font-mono">{pkg.name}</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-muted-foreground">{pkg.desc}</p>
              </CardContent>
            </Card>
          ))}
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t bg-muted/40">
        <div className="mx-auto max-w-5xl px-6 py-8 text-sm text-muted-foreground flex flex-col sm:flex-row items-center justify-between gap-3">
          <span>errorwrapper-v3 — MIT License</span>
          <span>Built with TanStack Start + Tailwind CSS</span>
        </div>
      </footer>
    </div>
  );
}
