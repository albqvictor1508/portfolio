import { useState } from "react";
import "./App.css";

export function App() {
	const [count, setCount] = useState(0);

	return (
		<div>
			<p>counter: {count}</p>
			<p>salve</p>

			<button type="button" onClick={() => setCount(count + 1)}>
				Add to counter
			</button>
		</div>
	);
}
