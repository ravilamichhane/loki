import { Toaster } from "@/components/ui/sonner";
import { TooltipProvider } from "@/components/ui/tooltip";
import { QueryProvider } from "@/providers/query-client-provider";
import { createRootRoute, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";

export const Route = createRootRoute({
  component: () => (
    <>
      <Toaster />
      <TooltipProvider>
        <QueryProvider>
          <Outlet />
          <TanStackRouterDevtools />
        </QueryProvider>
      </TooltipProvider>
    </>
  ),
});
