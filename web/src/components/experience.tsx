import { formatDate } from "../utils/format-date";

export const ExperienceSection = () => {
	const experiences = [
		{
			companyName: "Lorem Ipsum Company",
			startDate: new Date("2023-01-15"),
			endDate: null,
			role: "Software Engineer",
			description:
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		},
		{
			companyName: "Dolor Sit Amet Inc.",
			startDate: new Date("2021-06-01"),
			endDate: new Date("2022-12-31"),
			role: "Frontend Developer",
			description:
				"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
		},
		{
			companyName: "Consectetur Adipiscing Corp.",
			startDate: new Date("2019-08-20"),
			endDate: new Date("2021-05-25"),
			role: "Backend Developer",
			description:
				"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		},
		{
			companyName: "Tempor Incididunt LLC",
			startDate: new Date("2018-03-10"),
			endDate: new Date("2019-07-15"),
			role: "Junior Developer",
			description:
				"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		},
	];
	return (
		<div className="w-full h-full flex flex-col gap-4">
			<h2 className="text-2xl font-semibold">My Experience. </h2>
			<p className="text-sm font-light w-3/4 leading-6 text-zinc-400">
				A brief overview of my professional journey and roles.
			</p>

			{experiences.map((xp, index) => (
				<div
					className="flex flex-col gap-4"
					key={new Date().toString() + index.toString()}
				>
					<div className="flex flex-col gap-1">
						<h3 className="text-lg font-semibold">{xp.companyName}</h3>
						<p className="text-sm text-zinc-400">
							{formatDate(xp.startDate)} -{" "}
							{xp.endDate ? formatDate(xp.endDate) : "Present"}
						</p>
					</div>
					<p className="text-sm">{xp.description}</p>
				</div>
			))}
		</div>
	);
};
