import { Outlet, createFileRoute, redirect } from '@tanstack/react-router'
import { authClient } from '@/lib/auth'

export const Route = createFileRoute('/(private)')({
  component: RouteComponent,
  beforeLoad: async () => {
    const session = await authClient.getSession();

    if (!session.data) {
      throw redirect({ to: "/signin" });
    }
  }
})

function RouteComponent() {
  return <Outlet />
}
