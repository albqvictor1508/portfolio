import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea.tsx";

export const ContactSection = () => {
	return (
		<div className="w-full h-full flex flex-col gap-4">
			<h2 className="text-lg font-semibold">Contact me.</h2>
			<p className="text-sm text-zinc-400">lorem ipsum.</p>

			<form className="w-full h-full flex flex-col p-4 gap-8 rounded-md bg-zinc-900">
				<div
					className="w-full flex gap-6
          "
				>
					<Input className="" placeholder="Your name." />
					<Input placeholder="Your email." />
				</div>
				<Textarea
					className="w-full h-[150px]"
					maxLength={500}
					minLength={3}
					placeholder="Your message"
				/>
			</form>
		</div>
	);
};
