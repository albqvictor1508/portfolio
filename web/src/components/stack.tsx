import { Badge } from "./ui/badge";

export const StackSection = () => {
	const technologies = [
		{ icon: "", name: "Go" },
		{ icon: "", name: "Typescript" },
		{ icon: "", name: "Java" },
		{ icon: "", name: "Docker" },
		{ icon: "", name: "AWS" },
		{ icon: "", name: "React Native" },
		{ icon: "", name: "Postgres" },
		{ icon: "", name: "MongoDB" },
		{ icon: "", name: "Websocket" },
	];
	return (
		<div className="w-full h-full flex flex-col gap-4">
			<h2 className="text-2xl font-semibold">My Technologies.</h2>
			<p className="text-sm text-zinc-400">
				Building high-performance websites with clean code and strong SEO
				fundamentals.
			</p>
			<div className="flex m-auto  gap-4 items-center flex-wrap justify-center ">
				{technologies.map((tech: { icon: string; name: string }) => (
					<Badge
						variant={"outline"}
						className="p-2 flex justify-center items-center gap-2"
						key={tech.name}
					>
						<span className="w-6 h-6">
							<img src={tech.icon} alt="tech icon" />
						</span>
						{tech.name}
					</Badge>
				))}
			</div>
		</div>
	);
};
