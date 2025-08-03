import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useRef } from "react";

const client = new QueryClient();

export function App() {
	const salve = useRef("salve");
	salve.current = "mudei o valor de salve";

	return (
		<QueryClientProvider client={client}>
			<div className="w-full h-full">
				<p>{salve.current}</p>
			</div>
		</QueryClientProvider>
	);
}
