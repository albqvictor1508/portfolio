import type { ProjectSchema } from "@/types/project";
import { Project } from "./project.tsx";

export const ProjectSection = ({ projects }: { projects: ProjectSchema[] }) => {
	return (
		<div className="w-full h-full flex flex-col gap-2">
			<h2>Projects</h2>
			<p>lorem</p>

			{/*
      {projects.map(p => (

      ))}
    */}
			<Project />
		</div>
	);
};
