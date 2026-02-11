import { Aside1 } from '@/components/home/aside1'
import { Aside2 } from '@/components/home/aside2'
import { Header } from '@/components/home/header'
import { Main } from '@/components/home/main'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/home')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <main className="min-h-screen text-white">

        <Header/>

      <div className=" w-full grid grid-cols-2 md:grid-cols-[0.5fr_2fr_0.5fr] gap-4">
        <Aside1/>

        <Main/>

        <Aside2/>
      </div>
    </main>
  )
}
