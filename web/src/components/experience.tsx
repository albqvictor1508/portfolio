function formatDate(date: Date): string {
	const day = date.getDay().toFixed(2);
	const month = date.getMonth();

	return `${day}/${month}`;
}

export const ExperienceSection = () => {
	const experiences = [
		{
			companyName: "salve",
			startDate: new Date(),
			endDate: null,
			role: "backend",
			nsei: "salve",
			description: "salve",
		},
		{
			companyName: "salve",
			startDate: new Date(),
			endDate: new Date(),
			role: "backend",
			description: "salve",
		},
		{
			companyName: "salve",
			startDate: new Date(),
			endDate: new Date(),
			role: "backend",
			description: "salve",
		},
		{
			companyName: "salve",
			startDate: new Date(),
			endDate: new Date(),
			role: "backend",
			description: "salve",
		},
	];
	return (
		<div className="w-full h-full flex flex-col gap-4 ">
			<h2>My Experience </h2>
			<p className="text-sm text-zinc-400">lorem</p>

			{experiences.map((xp, index) => (
				<div key={new Date().toString() + index.toString()}>
					<h2 className="text-sm text-zinc-400">{xp.companyName}</h2>
					<p>
						{formatDate(xp.startDate)} -{" "}
						{xp.endDate ? formatDate(xp.endDate) : "Presente"}
					</p>
					<p>{xp.description}</p>
				</div>
			))}
		</div>
	);
};
