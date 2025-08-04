export const HeroSection = () => {
	return (
		<div className="w-full h-full flex flex-col gap-4">
			<div className="w-full h-full flex gap-8">
				<div className="flex flex-1 flex-col gap-4">
					<h1 className="text-3xl font-bold">
						Hey, I'm <span className="text-blue-400">Victor Arruda</span>
					</h1>
					<ul className="flex flex-col gap-2">
						<li>Back-end Engineer</li>
						<li className="text-sm text-justify">
							Founder{" "}
							<span className="font-bold text-orange-500">Zentto Chatbot</span>,
							a software development chatbot dedicated to talk with customer .
							We specialize in building modern SaaS platforms, websites, mobile
							apps, and internal tools tailored for performance, scalability,
							and great user experience.
						</li>
					</ul>
				</div>
				<div className="flex flex-1 justify-center items-center w-max rounded-md border-2 border-zinc-800">
					salveeeeee hero
				</div>
			</div>
			<div className="w-full flex justify-between">
				<div className="">location</div>
				<div className="">email</div>
			</div>
			<div className="w-full flex justify-between">
				<span className="text-sm text-muted">salve</span>
				<span className="text-sm text-muted">salve</span>
				<span className="text-sm text-muted">salve</span>
				<span className="text-sm text-muted">salve</span>
			</div>
		</div>
	);
};
