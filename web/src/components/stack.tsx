import { FaAws, FaDocker, FaJava } from "react-icons/fa";
import {
  SiBun,
  SiExpress,
  SiFastify,
  SiGit,
  SiGo,
  SiMongodb,
  SiNodedotjs,
  SiPostgresql,
  SiReact,
  SiSocketdotio,
  SiSpringboot,
  SiTypescript,
} from "react-icons/si";
import { Badge } from "./ui/badge";

export const StackSection = () => {
  const technologies = [
    { icon: <SiGo size={24} />, name: "Go" },
    { icon: <SiTypescript size={24} />, name: "Typescript" },
    { icon: <SiNodedotjs size={24} />, name: "Node.js" },
    { icon: <SiBun size={24} />, name: "Bun" },
    { icon: <SiGit size={24} />, name: "Git" },
    { icon: <FaJava size={24} />, name: "Java" },
    { icon: <SiSpringboot size={24} />, name: "Springboot" },
    { icon: <FaDocker size={24} />, name: "Docker" },
    { icon: <FaAws size={24} />, name: "AWS" },
    { icon: <SiReact size={24} />, name: "React Native" },
    { icon: <SiPostgresql size={24} />, name: "Postgres" },
    { icon: <SiMongodb size={24} />, name: "MongoDB" },
    { icon: <SiExpress size={24} />, name: "Express.js" },
    { icon: <SiFastify size={24} />, name: "Fastify" },
    { icon: <SiSocketdotio size={24} />, name: "Websocket" },
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
            className="py-1 px-2 flex justify-center items-center gap-2"
            key={tech.name}
          >
            {tech.icon}
            {tech.name}
          </Badge>
        ))}
      </div>
    </div>
  );
};
