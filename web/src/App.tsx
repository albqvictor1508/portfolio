import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useRef } from "react";
import { Home } from "./components/home";

const client = new QueryClient();

export function App() {
  const salve = useRef("salve");
  salve.current = "mudei o valor de salve";

  return (
    <QueryClientProvider client={client}>
      <Home />
    </QueryClientProvider>
  );
}
