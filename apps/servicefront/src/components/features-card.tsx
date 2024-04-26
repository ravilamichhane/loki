/**
 * v0 by Vercel.
 * @see https://v0.dev/t/VIJEyUAh0Aq
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
export function Features() {
  return (
    <section className="w-full py-12 md:py-24 lg:py-32">
      <div className="container space-y-12 px-4 md:px-6">
        <div className="flex flex-col items-center justify-center space-y-4 text-center">
          <div className="space-y-2">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-5xl">Empower Your Business with Ecommerce</h2>
            <p className="max-w-[900px] text-gray-500 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed dark:text-gray-400">
              Discover how Ecommerce's powerful features can transform your online business and drive success.
            </p>
          </div>
        </div>
        <div className="container px-4 md:px-6 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-x-10 gap-20">
          <div className="flex flex-col items-center text-center space-y-4">
            <BotIcon className="h-12 w-12 text-gray-900 dark:text-gray-50" />
            <div>
              <h3 className="text-lg font-semibold">Automation</h3>
              <p className="text-gray-500 dark:text-gray-400">Automate your workflows and save time.</p>
            </div>
          </div>
          <div className="flex flex-col items-center text-center space-y-4">
            <PieChartIcon className="h-12 w-12 text-gray-900 dark:text-gray-50" />
            <div>
              <h3 className="text-lg font-semibold">Analytics</h3>
              <p className="text-gray-500 dark:text-gray-400">Gain insights into your business data.</p>
            </div>
          </div>
          <div className="flex flex-col items-center text-center space-y-4">
            <ScalingIcon className="h-12 w-12 text-gray-900 dark:text-gray-50" />
            <div>
              <h3 className="text-lg font-semibold">Growth</h3>
              <p className="text-gray-500 dark:text-gray-400">Grow your Ecommerce store with our tools.</p>
            </div>
          </div>
          <div className="flex flex-col items-center text-center space-y-4">
            <BotIcon className="h-12 w-12 text-gray-900 dark:text-gray-50" />
            <div>
              <h3 className="text-lg font-semibold">Automation</h3>
              <p className="text-gray-500 dark:text-gray-400">Automate your workflows and save time.</p>
            </div>
          </div>
          <div className="flex flex-col items-center text-center space-y-4">
            <PieChartIcon className="h-12 w-12 text-gray-900 dark:text-gray-50" />
            <div>
              <h3 className="text-lg font-semibold">Analytics</h3>
              <p className="text-gray-500 dark:text-gray-400">Gain insights into your business data.</p>
            </div>
          </div>
          <div className="flex flex-col items-center text-center space-y-4">
            <ScalingIcon className="h-12 w-12 text-gray-900 dark:text-gray-50" />
            <div>
              <h3 className="text-lg font-semibold">Growth</h3>
              <p className="text-gray-500 dark:text-gray-400">Grow your Ecommerce store with our tools.</p>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

function BotIcon(props: any) {
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
      <path d="M12 8V4H8" />
      <rect width="16" height="12" x="4" y="8" rx="2" />
      <path d="M2 14h2" />
      <path d="M20 14h2" />
      <path d="M15 13v2" />
      <path d="M9 13v2" />
    </svg>
  );
}

function PieChartIcon(props: any) {
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
      <path d="M21.21 15.89A10 10 0 1 1 8 2.83" />
      <path d="M22 12A10 10 0 0 0 12 2v10z" />
    </svg>
  );
}

function ScalingIcon(props: any) {
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
      <path d="M21 3 9 15" />
      <path d="M12 3H3v18h18v-9" />
      <path d="M16 3h5v5" />
      <path d="M14 15H9v-5" />
    </svg>
  );
}
