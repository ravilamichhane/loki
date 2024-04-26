import { StoreIcon } from "@/components/header";
import { Link, Outlet, createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_auth/_layout")({
  component: LayoutComponent,
});

function LayoutComponent() {
  return (
    <>
      <div className="min-h-screen overflow-y-scroll w-full flex items-center">
        <Link className="flex items-center fixed top-6 left-12 " to="/">
          <StoreIcon className="h-6 w-6" />
          <span className="sr-only">Ecommerce App</span>
        </Link>
        <Outlet />
      </div>
    </>
  );
}
