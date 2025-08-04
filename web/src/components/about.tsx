import { Button } from "./ui/button";

export const AboutSection = () => {
	return (
		<div className="w-full h-full flex gap-8">
			<div className="w-full flex-1 flex flex-col gap-4">
				<div className="w-full flex flex-col gap-2">
					<h2 className="font-semibold text-2xl">About me.</h2>
					<p className="text-sm text-justify">
						Lorem Ipsum is simply dummy text of the printing and typesetting
						industry. Lorem Ipsum has been the industry's standard dummy text
						ever since the 1500s, when an unknown printer took a galley of type
						and scrambled it to make a type specimen book. It has survived not
						only five centuries.
					</p>
					<p className="text-sm text-justify">
						But also the leap into electronic typesetting, remaining essentially
						unchanged. It was popularised in the 1960s with the release of
						Letraset sheets containing Lorem Ipsum passages, and more recently
						with desktop publishing software like Aldus PageMaker including
						versions of Lorem Ipsum.
					</p>
				</div>
				<div className="flex gap-2">
					<Button variant={"default"}>View on Github</Button>
					<Button variant={"secondary"}>Contact me</Button>
				</div>
			</div>
			<div className="w-full flex justify-center items-center flex-1 rounded-md border-2 border-r-zinc-800 p-8">
				salve
			</div>
		</div>
	);
};
