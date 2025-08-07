import { Project } from "./project";

const projects = [
	{
		name: "Flameshot",
		startDate: "01 Jan, 2017",
		endDate: "Present",
		description:
			"Powerful yet simple to use screenshot software. Flameshot has a huge range of tools to use, it is also highly customizable and it is available for Linux, Mac and Windows.",
		imageUrl: "https://flameshot.org/media/images/preview.png",
		repoUrl: "https://github.com/flameshot-org/flameshot",
		demoUrl: "https://flameshot.org/",
	},
	{
		name: "ShareX",
		startDate: "01 Jan, 2007",
		endDate: "Present",
		description:
			"ShareX is a free and open source program that lets you capture or record any area of your screen and share it with a single press of a key. It also allows uploading images, text or other types of files to over 80 supported destinations you can choose from.",
		imageUrl: "https://getsharex.com/img/hero_image.png",
		repoUrl: "https://github.com/ShareX/ShareX",
		demoUrl: "https://getsharex.com/",
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
