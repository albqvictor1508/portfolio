import { Project } from "./project";

export const ProjectSection = () => {
  return (
    <div className="w-full h-full flex flex-col gap-10">
      <div className="w-full h-full flex flex-col gap-2">
        <h2 className="text-2xl font-semibold">Projects</h2>
        <p className="w-2/3 text-justify text-base">
          Lorem Ipsum is simply dummy text of the printing and typesetting five
          centuries.
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
