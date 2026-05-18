import { createFileRoute } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

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

function Index() {
  return (
    <div className="min-h-screen bg-background text-foreground">
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
              <a href="https://github.com/alimtvnetwork/errorwrapper-v3">View on GitHub</a>
            </Button>
            <Button variant="outline" size="lg" asChild>
              <a href="/docs">Documentation</a>
            </Button>
          </div>
        </div>
      </section>

      {/* Quick install */}
      <section className="mx-auto max-w-5xl px-6 py-12">
        <h2 className="text-2xl font-semibold tracking-tight">Quick start</h2>
        <div className="mt-4 rounded-lg border bg-card p-4 font-mono text-sm text-card-foreground overflow-x-auto">
          <code>go get github.com/alimtvnetwork/errorwrapper-v3</code>
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
        <h2 className="text-2xl font-semibold tracking-tight">Package index</h2>
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
