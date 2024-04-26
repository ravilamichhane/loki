import { createLazyFileRoute } from '@tanstack/react-router'

export const Route = createLazyFileRoute('/_dashboard/_layout/dashboard')({
  component: () => <div>Hello /_dashboard/_layout/dashboard!</div>
})