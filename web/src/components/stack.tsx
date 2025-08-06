import { Badge } from "./ui/badge";
import {
  SiGo,
  SiTypescript,
  SiReact,
  SiPostgresql,
  SiMongodb,
  SiSocketdotio,
  SiNodedotjs,
  SiBun,
  SiGit,
  SiSpringboot,
  SiExpress,
  SiFastify,
} from "react-icons/si";

export const StackSection = () => {
  const technologies = [
    { icon: <SiGo className="w-full h-full" />, name: "Go" },
    { icon: <SiTypescript className="w-full h-full" />, name: "Typescript" },
    { icon: <SiNodedotjs className="w-full h-full" />, name: "Node.js" },
    { icon: <SiBun className="w-full h-full" />, name: "Bun" },
    { icon: <SiGit className="w-full h-full" />, name: "Git" },
    { icon: "", name: "Java" },
    { icon: <SiSpringboot className="w-full h-full" />, name: "Springboot" },
    { icon: "", name: "Docker" },
    { icon: "", name: "AWS" },
    { icon: <SiReact className="w-full h-full" />, name: "React Native" },
    { icon: <SiPostgresql className="w-full h-full" />, name: "Postgres" },
    { icon: <SiMongodb className="w-full h-full" />, name: "MongoDB" },
    { icon: <SiExpress className="w-full h-full" size={ } />, name: "Express.js" },
    { icon: <SiFastify className="w-full h-full" />, name: "Fastify" },
    { icon: <SiSocketdotio className="w-full h-full" />, name: "Websocket" },
  ];
  return (
    <div className="w-full h-full flex flex-col gap-4">
      <h2 className="text-2xl font-semibold">My Technologies.</h2>
      <p className="text-sm text-zinc-400">
        I work with a modern stack, focusing on performance, scalability, and
        robust solutions.
      </p>
      <div className="flex m-auto gap-3 items-center flex-wrap ">
        {technologies.map((tech: { icon: React.ReactNode; name: string }) => (
          <Badge
            variant={"outline"}
            className="p-2 flex justify-center items-center gap-2"
            key={tech.name}
          >
            <span className="w-6 h-6">{tech.icon}</span>
            {tech.name}
          </Badge>
        ))}
      </div>
    </div>
  );
};
