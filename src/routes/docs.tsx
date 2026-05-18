import { createFileRoute, Link } from "@tanstack/react-router";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { BookOpen, FileText, ArrowLeft, ExternalLink } from "lucide-react";

export const Route = createFileRoute("/docs")({
  component: DocsIndex,
});

const docPages = [
  {
    title: "Extending Error Types",
    path: "/docs/extending-error-types",
    desc: "How to extend errorwrapper-v3 with custom error types and behaviors. Covers interface-based, registry-based, and context-based approaches.",
  },
  {
    title: "LLM Guideline",
    path: "/docs/llm-guideline",
    desc: "Guidelines for AI agents working with the errorwrapper-v3 codebase. Package layout, naming conventions, testing rules, and anti-patterns.",
  },
];

function DocsIndex() {
  return (
    <div className="min-h-screen bg-background text-foreground">
      <div className="border-b bg-muted/40">
        <div className="mx-auto max-w-5xl px-6 py-8">
          <Link
            to="/"
            className="inline-flex items-center gap-1.5 text-sm text-muted-foreground hover:text-foreground transition-colors"
          >
            <ArrowLeft className="h-4 w-4" />
            Back to home
          </Link>
          <h1 className="mt-4 text-3xl font-bold tracking-tight flex items-center gap-2">
            <BookOpen className="h-7 w-7 text-primary" />
            Documentation
          </h1>
          <p className="mt-2 text-muted-foreground">
            Guides and reference for the errorwrapper-v3 library.
          </p>
        </div>
      </div>

      <div className="mx-auto max-w-5xl px-6 py-12">
        <div className="grid gap-4 sm:grid-cols-2">
          {docPages.map((page) => (
            <Link key={page.path} to={page.path}>
              <Card className="hover:shadow-md transition-shadow h-full cursor-pointer">
                <CardHeader className="pb-3">
                  <CardTitle className="flex items-center gap-2 text-lg">
                    <FileText className="h-5 w-5 text-primary" />
                    {page.title}
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <CardDescription className="text-muted-foreground">
                    {page.desc}
                  </CardDescription>
                </CardContent>
              </Card>
            </Link>
          ))}
        </div>

        <div className="mt-12 rounded-lg border bg-card p-6">
          <h2 className="text-xl font-semibold tracking-tight flex items-center gap-2">
            <ExternalLink className="h-5 w-5 text-primary" />
            Additional resources
          </h2>
          <ul className="mt-4 space-y-2 text-sm text-muted-foreground">
            <li>
              <a
                href="https://github.com/alimtvnetwork/errorwrapper-v3"
                target="_blank"
                rel="noopener noreferrer"
                className="text-foreground hover:underline"
              >
                GitHub Repository
              </a>
              {" — Source code and issues"}
            </li>
            <li>
              <a
                href="https://github.com/alimtvnetwork/errorwrapper-v3/blob/main/README.md"
                target="_blank"
                rel="noopener noreferrer"
                className="text-foreground hover:underline"
              >
                README
              </a>
              {" — Quick start and examples"}
            </li>
          </ul>
        </div>
      </div>

      <footer className="border-t bg-muted/40 mt-auto">
        <div className="mx-auto max-w-5xl px-6 py-8 text-sm text-muted-foreground text-center">
          errorwrapper-v3 — MIT License
        </div>
      </footer>
    </div>
  );
}
