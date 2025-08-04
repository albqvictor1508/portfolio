import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const client = new QueryClient();

export function App() {
  return <QueryClientProvider client={client}></QueryClientProvider>;
}
