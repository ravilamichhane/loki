import { PricingCards } from "@/components/pricing-cards";
import { createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/_app/_layout/pricing")({
  component: () => <PricingCards />,
});
