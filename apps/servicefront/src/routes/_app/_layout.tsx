import { Footer } from "@/components/footer";
import { Header } from "@/components/header";
import { createFileRoute, Outlet } from "@tanstack/react-router";

export const Route = createFileRoute("/_app/_layout")({
  component: HomePage,
});

export default function HomePage() {
  return (
    <div className="flex flex-col min-h-[100dvh]">
      <Header />
      <Outlet />
      <Footer />
    </div>
  );
}
