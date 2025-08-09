import { FaAws, FaDocker, FaJava } from "react-icons/fa";
import {
  SiArchlinux,
  SiBun,
  SiExpress,
  SiFastify,
  SiGit,
  SiGo,
  SiJest,
  SiJunit5,
  SiMongodb,
  SiNodedotjs,
  SiPostgresql,
  SiReact,
  SiRedis,
  SiSocketdotio,
  SiSpringboot,
  SiTailwindcss,
  SiTypescript,
  SiVite,
  SiVitest,
  SiNeovim,
} from "react-icons/si";
import { useLanguage } from "../context/LanguageContext";
import { Badge } from "./ui/badge";

export const StackSection = () => {
  const { t } = useLanguage();
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
    { icon: <SiRedis size={24} />, name: "Redis" },
    { icon: <SiVite size={24} />, name: "Vite" },
    { icon: <SiTailwindcss size={24} />, name: "TailwindCSS" },
    { icon: <SiJest size={24} />, name: "Jest" },
    { icon: <SiVitest size={24} />, name: "Vitest" },
    { icon: <SiJunit5 size={24} />, name: "JUnit" },
    { icon: <SiArchlinux size={24} />, name: "Arch Linux" },
    { icon: <SiNeovim size={24} />, name: "Neovim" },
  ];
  return (
    <div className="w-full h-full flex flex-col gap-4">
      <h2 className="text-2xl font-semibold">{t("stack_section.title")}</h2>
      <p className="text-sm text-zinc-400">{t("stack_section.description")}</p>
      <div className="flex m-auto gap-3 items-center flex-wrap ">
        {technologies.map((tech: { icon: React.ReactNode; name: string }) => (
          <Badge
            variant={"outline"}
            className="p-2 flex justify-center items-center gap-2"
            key={tech.name}
          >
            <span>{tech.icon}</span>
            {tech.name}
          </Badge>
        ))}
      </div>
    </div>
  );
};
