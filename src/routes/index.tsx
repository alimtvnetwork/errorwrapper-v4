import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: Index,
});

function Index() {
  return (
    <div
      data-lovable-blank-page-placeholder
      className="flex min-h-screen items-center justify-center bg-background"
    >
      <p className="text-sm text-muted-foreground">Blank page</p>
    </div>
  );
}
