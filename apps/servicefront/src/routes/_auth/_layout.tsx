import { Outlet, createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_auth/_layout")({
  component: LayoutComponent,
});

function LayoutComponent() {
  return (
    <div className="w-screen h-screen overflow-y-scroll flex items-center justify-center">
      <Outlet />
    </div>
  );
}
