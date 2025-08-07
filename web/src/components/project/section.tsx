import { Project } from "./project";

const projects = [
	{
		name: "Lorem Ipsum Project 1",
		startDate: "15 Jan, 2023",
		endDate: "Present",
		description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		imageUrl: "https://source.unsplash.com/random/800x600?tech,project",
		repoUrl: "#",
		demoUrl: "#",
	},
	{
		name: "Dolor Sit Amet Project 2",
		startDate: "01 Jun, 2021",
		endDate: "31 Dec, 2022",
		description: "Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
		imageUrl: "https://source.unsplash.com/random/800x600?web,development",
		repoUrl: "#",
		demoUrl: "#",
	},
	{
		name: "Consectetur Adipiscing Project 3",
		startDate: "20 Aug, 2019",
		endDate: "25 May, 2021",
		description: "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		imageUrl: "https://source.unsplash.com/random/800x600?code,software",
		repoUrl: "#",
		demoUrl: "#",
	},
];

export const ProjectSection = () => {
	return (
		<div className="w-full h-full flex flex-col gap-8">
			<div className="w-full h-full flex flex-col gap-2">
				<h2 className="text-2xl font-semibold"> Projects</h2>
				<p className="w-2/3 text-justify text-sm">
					Lorem Ipsum is simply dummy text of the printing and typesetting five
					centuries.
				</p>
			</div>

			{projects.map((project, index) => (
				<Project key={index} {...project} />
			))}
		</div>
	);
};
