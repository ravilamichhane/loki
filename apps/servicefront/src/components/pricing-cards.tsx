import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";

export function PricingCards() {
  return (
    <section className="w-full py-12 md:py-24 lg:py-32">
      <div className="container px-4 md:px-6">
        <div className="grid gap-4 md:gap-8">
          <div className="text-center">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl">Pricing</h2>
            <p className="mx-auto max-w-[700px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400">
              Choose the plan that fits your needs.
            </p>
          </div>
          <div className="grid md:grid-cols-3 gap-6">
            <Card className="w-full flex flex-col justify-between">
              <CardHeader className="flex flex-col items-center space-y-2">
                <CardTitle className="text-2xl">Free</CardTitle>
                <CardDescription className="text-center">Get started for free</CardDescription>
                <div className="text-3xl font-semibold">€0</div>
              </CardHeader>
              <CardContent className="grid gap-4">
                <ul className="grid gap-4">
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    Basic email support
                  </li>
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    Access to free content
                  </li>
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    Community forums
                  </li>
                </ul>
              </CardContent>
              <CardFooter className="flex justify-end ">
                <Button className="w-full">Get Started</Button>
              </CardFooter>
            </Card>
            <Card className="w-full flex flex-col justify-between">
              <CardHeader className="flex flex-col items-center space-y-2">
                <CardTitle className="text-2xl">Business</CardTitle>
                <CardDescription className="text-center">Perfect for small businesses</CardDescription>
                <div className="text-3xl font-semibold">€49</div>
              </CardHeader>
              <CardContent className="grid gap-4">
                <ul className="grid gap-4">
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    5-day email support
                  </li>
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    Access to premium content
                  </li>
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    Exclusive webinars
                  </li>
                </ul>
              </CardContent>
              <CardFooter className="flex justify-end">
                <Button className="w-full">Buy Now</Button>
              </CardFooter>
            </Card>
            <Card className="w-full flex flex-col justify-between">
              <CardHeader className="flex flex-col items-center space-y-2">
                <CardTitle className="text-2xl">Enterprise</CardTitle>
                <CardDescription className="text-center">For large businesses and teams</CardDescription>
                <div className="text-3xl font-semibold">€499</div>
              </CardHeader>
              <CardContent className="grid gap-4">
                <ul className="grid gap-4">
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    24/7 priority email support
                  </li>
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    Unlimited access to premium content
                  </li>
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    Dedicated account manager
                  </li>
                  <li className="flex items-center space-x-2">
                    <CheckIcon className="h-4 w-4" />
                    Custom integrations and features
                  </li>
                </ul>
              </CardContent>
              <CardFooter className="flex justify-end">
                <Button className="w-full">Contact Sales</Button>
              </CardFooter>
            </Card>
          </div>
        </div>
      </div>
    </section>
  );
}

function CheckIcon(props: any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <polyline points="20 6 9 17 4 12" />
    </svg>
  );
}
