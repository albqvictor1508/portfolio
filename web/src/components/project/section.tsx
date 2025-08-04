import type { ProjectSchema } from "@/types/project";
import { Project } from "./project.tsx";

export const ProjectSection = ({ projects }: { projects: ProjectSchema[] }) => {
	return (
		<div className="w-full h-full flex flex-col gap-10">
			<div className="w-full h-full flex flex-col gap-2">
				<h2 className="text-2xl font-semibold">Projects</h2>
				<p className="w-2/3 text-justify">
					Lorem Ipsum is simply dummy text of the printing and typesetting
					scrambled it to make a type specimen book. It has survived not only
					five centuries.
				</p>
			</div>

			{/*

      {projects.map(p => (
      ))}
    */}
			<Project />
			<Project />
			<Project />
		</div>
	);
};
