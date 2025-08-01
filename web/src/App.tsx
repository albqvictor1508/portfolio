import { useState } from "react";
import "./App.css";

function App() {
	const [count, setCount] = useState(0);

	return (
		<div>
			<p>counter: {count}</p>
			<p>salve</p>
		</div>
	);
}

export default App;
