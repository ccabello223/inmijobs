import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/(private)/')({
  component: App,
})

function App() {
  return (
    <div>Inmijobs</div>
  )
}
