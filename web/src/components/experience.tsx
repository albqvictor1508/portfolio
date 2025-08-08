import { FaDocker } from "react-icons/fa";
import { SiFastify, SiMongodb, SiTypescript } from "react-icons/si";
import { useLanguage } from "../context/LanguageContext";
import { formatDate } from "../utils/format-date";
import { Badge } from "./ui/badge.tsx";

export const ExperienceSection = () => {
  const { t } = useLanguage();
  const experiences = [
    {
      companyName: "Agility Telecom",
      startDate: new Date("2023-01-15"),
      endDate: null,
      role: "Software Engineer",
      description:
        "Entrei na empresa no setor de tributos, não como dev, e após meses trabalhando na parte fiscal, criei o Zentto Chatbot e o integrei ao serviço interno da empresa, trazendo grandes melhorias para o atendimento, e por isso, migrei para trabalhar como dev, trazendo inovações para a empresa.",
      technologies: [
        { icon: <SiTypescript size={24} />, name: "Typescript" },
        { icon: <SiMongodb size={24} />, name: "MongoDB" },
        { icon: <FaDocker size={24} />, name: "Docker" },
        { icon: <SiFastify size={24} />, name: "Fastify" },
      ],
    },

    {
      companyName: "Dosemed",
      startDate: new Date("2021-06-01"),
      endDate: new Date("2022-12-31"),
      role: "Backend Developer",
      description:
        "Fui responsável pela criação de um ecossistema para gerenciamento de medicamentos, com suas medidas, fórmulas, avisos, erros, entre outros. E também fui responsável por integrar o sistema com pagamentos, utilizando Stripe para integração de Cartão de Crédito e AbacatePay para pagamentos via Pix",
      technologies: [
        { icon: <SiTypescript size={24} />, name: "Typescript" },
        { icon: <SiMongodb size={24} />, name: "MongoDB" },
        { icon: <FaDocker size={24} />, name: "Docker" },
        { icon: <SiFastify size={24} />, name: "Fastify" },
      ],
    },
    {
      companyName: "Utility",
      startDate: new Date("2019-08-20"),
      endDate: new Date("2021-05-25"),
      role: "Freelance",
      description:
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
      technologies: [
        { icon: <SiTypescript size={24} />, name: "Typescript" },
        { icon: <SiMongodb size={24} />, name: "MongoDB" },
        { icon: <FaDocker size={24} />, name: "Docker" },
        { icon: <SiFastify size={24} />, name: "Fastify" },
      ],
    },
    {
      companyName: "Tempor Incididunt LLC",
      startDate: new Date("2018-03-10"),
      endDate: new Date("2019-07-15"),
      role: "Junior Developer",
      description:
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
      technologies: [
        { icon: <SiTypescript size={24} />, name: "Typescript" },
        { icon: <SiMongodb size={24} />, name: "MongoDB" },
        { icon: <FaDocker size={24} />, name: "Docker" },
        { icon: <SiFastify size={24} />, name: "Fastify" },
      ],
    },
  ];
  return (
    <div className="w-full h-full flex flex-col gap-6">
      <div className="flex flex-col gap-2">
        <h2 className="text-2xl font-semibold">
          {t("experience_section.title")}
        </h2>
        <p className="text-sm font-light w-3/4 leading-6 text-zinc-400">
          {t("experience_section.description")}
        </p>
      </div>

      <div className="flex flex-col gap-12">
        {experiences.map((xp, index) => (
          <div
            className="flex flex-col gap-4"
            key={new Date().toString() + index.toString()}
          >
            <div className="flex flex-col gap-1">
              <h3>
                <span className="text-xl font-semibold">{xp.companyName}</span>{" "}
                <span className="text-sm text-zinc-400">
                  ({formatDate(xp.startDate)} -{" "}
                  {xp.endDate
                    ? formatDate(xp.endDate)
                    : t("experience_section.present")}
                  )
                </span>
              </h3>
              <p className="text-sm text-zinc-400">
                <span className="italic text-sm text-zinc-400">{xp.role}</span>
              </p>
            </div>
            <p className="text-sm">{xp.description}</p>
            {/*
            <div>
              {xp.technologies &&
                xp.technologies?.map((tech) => {
                  <Badge
                    variant={"outline"}
                    className="p-2 flex justify-center items-center gap-2"
                    key={tech.name}
                  >
                    <span>{tech.icon}</span>
                    {tech.name}
                  </Badge>;
                })}
            </div>
            */}
          </div>
        ))}
      </div>
    </div>
  );
};
