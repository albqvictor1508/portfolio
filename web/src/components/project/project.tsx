import { Button } from "../ui/button";

interface ProjectProps {
	name: string;
	startDate: string;
	endDate: string;
	description: string;
	imageUrl: string;
	repoUrl: string;
	demoUrl: string;
}

export const Project = ({
	name,
	startDate,
	endDate,
	description,
	imageUrl,
	repoUrl,
	demoUrl,
}: ProjectProps) => {
	return (
		<div className="w-full h-full flex flex-col gap-4">
			<div className="flex flex-col gap-2">
				<h3 className="font-semibold text-xl">{name}</h3>
				<span className="text-sm text-zinc-500">
					{startDate} - {endDate}
				</span>
			</div>
			<p>{description}</p>

			<img
				src={imageUrl}
				alt={name}
				className="w-full h-68 rounded-md bg-zinc-800 object-cover"
			/>
			<div className="flex justify-end items-center gap-4">
				<a href={repoUrl} target="_blank" rel="noopener noreferrer">
					<Button variant={"secondary"}>View Repo</Button>
				</a>
				<a href={demoUrl} target="_blank" rel="noopener noreferrer">
					<Button variant={"default"}>View Demo</Button>
				</a>
			</div>
		</div>
	);
};
