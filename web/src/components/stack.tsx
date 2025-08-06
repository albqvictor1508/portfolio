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
import { FaJava, FaDocker, FaAws } from "react-icons/fa";

export const StackSection = () => {
  const technologies = [
    { icon: <SiGo size={16} />, name: "Go" },
    { icon: <SiTypescript size={14} />, name: "Typescript" },
    { icon: <SiNodedotjs size={14} />, name: "Node.js" },
    { icon: <SiBun size={14} />, name: "Bun" },
    { icon: <SiGit size={14} />, name: "Git" },
    { icon: <FaJava size={14} />, name: "Java" },
    { icon: <SiSpringboot size={14} />, name: "Springboot" },
    { icon: <FaDocker size={14} />, name: "Docker" },
    { icon: <FaAws size={14} />, name: "AWS" },
    { icon: <SiReact size={14} />, name: "React Native" },
    { icon: <SiPostgresql size={14} />, name: "Postgres" },
    { icon: <SiMongodb size={14} />, name: "MongoDB" },
    { icon: <SiExpress size={14} />, name: "Express.js" },
    { icon: <SiFastify size={14} />, name: "Fastify" },
    { icon: <SiSocketdotio size={14} />, name: "Websocket" },
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
            className=" flex justify-center items-center gap-2"
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
