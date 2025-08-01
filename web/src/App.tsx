import { useState } from "react";
import "./App.css";

export function App() {
	const [count, setCount] = useState(0);

	return (
		<div className="w-full h-full">
			<p>
				counter: <span>{count}</span>
			</p>
			<p>salve</p>

			<button type="button" onClick={(): void => setCount(count + 1)}>
				Add to counter
			</button>
		</div>
	);
}
