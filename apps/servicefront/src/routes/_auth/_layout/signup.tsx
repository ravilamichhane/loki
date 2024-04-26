import { SignUpForm } from "@/forms/auth/signup-form";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_auth/_layout/signup")({
  component: SignUpForm,
});
