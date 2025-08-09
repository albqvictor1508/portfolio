import React, { useEffect, useState } from "react";

interface TypingEffectProps {
	text: string;
	speed?: number;
}

const TypingEffect: React.FC<TypingEffectProps> = ({ text, speed = 150 }) => {
	const [displayedText, setDisplayedText] = useState("");
	const [index, setIndex] = useState(0);
	const [showCursor, setShowCursor] = useState(true);

	useEffect(() => {
		if (index < text.length) {
			const timeoutId = setTimeout(() => {
				setDisplayedText((prev) => prev + text.charAt(index));
				setIndex((prev) => prev + 1);
			}, speed);
			return (): void => clearTimeout(timeoutId);
		}

		const cursorInterval = setInterval(() => {
			setShowCursor((prev) => !prev);
		}, 500);
		return (): void => clearInterval(cursorInterval);
	}, [text, index, speed]);

	return (
		<span>
			{displayedText}
			<span className={showCursor ? "typing-cursor" : ""}>&nbsp;</span>
		</span>
	);
};

export default TypingEffect;
