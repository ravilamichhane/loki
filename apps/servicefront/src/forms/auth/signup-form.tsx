import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Form } from "@/components/ui/form";
import { Link } from "@tanstack/react-router";
import { TextInput } from "../inputs/text-input";
import useSignUpForm from "./use-signup-form";

export function SignUpForm() {
  const [form, submit] = useSignUpForm();

  return (
    <Card className="mx-auto max-w-sm flex-1">
      <CardHeader>
        <CardTitle className="text-xl">Sign Up</CardTitle>
        <CardDescription>Enter your information to create an account</CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <Form {...form}>
          <form className="grid gap-4" onSubmit={submit}>
            <TextInput control={form.control} name="first_name" label="First Name" placeholder="Ravi" />
            <TextInput control={form.control} name="last_name" label="Last Name" placeholder="Lamichhane" />
            <TextInput control={form.control} name="email" label="Email" placeholder="ravi@gmail.com" type="email" />

            <TextInput
              control={form.control}
              name="password"
              label="Password"
              placeholder="******************"
              type="password"
            />
            <Button type="submit" className="w-full">
              Create an account
            </Button>
          </form>
          <Button variant="outline" className="w-full">
            Sign up with GitHub
          </Button>
        </Form>
        <div className="mt-4 text-center text-sm">
          Already have an account?{" "}
          <Link to="/signin" className="underline">
            Sign in
          </Link>
        </div>
      </CardContent>
    </Card>
  );
}
