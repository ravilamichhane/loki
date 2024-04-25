import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/_app/_layout/about")({
  component: () => <div>Hello /about!</div>,
});
