import { formatDate } from "../../utils/format-date.ts";
import { Button } from "../ui/button";

interface ProjectProps {
	name: string;
	startDate: string;
	endDate: string;
	role: string;
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
	role,
	repoUrl,
	demoUrl,
}: ProjectProps) => {
	return (
		<div className="w-full h-full flex flex-col gap-4">
			<div className="flex flex-col gap-1">
				<h3>
					<span className="text-xl font-semibold">{name}</span>{" "}
					<span className="text-sm text-zinc-400">
						({startDate} - {endDate})
					</span>
				</h3>
				{role && <span>{role}</span>}
			</div>
			<p className="text-sm">{description}</p>
			{/*
			<img
				src={imageUrl}
				alt={name}
				className="w-full h-68 rounded-md bg-zinc-800 object-cover"
			/>
			*/}
			<div className="flex justify-end items-center gap-4">
				{repoUrl && (
					<a href={repoUrl} target="_blank" rel="noopener noreferrer">
						<Button variant={"secondary"}>View Repo</Button>
					</a>
				)}
				{demoUrl && (
					<a href={demoUrl} target="_blank" rel="noopener noreferrer">
						<Button variant={"default"}>View Demo</Button>
					</a>
				)}
			</div>
		</div>
	);
};
