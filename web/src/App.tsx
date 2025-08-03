import { useRef } from "react";

export function App() {
  const salve = useRef("salve");

  return (
    <div className="w-full h-full">
      <p>{salve.current}</p>
    </div>
  );
}
