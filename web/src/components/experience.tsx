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
        "Iniciei minha carreira na Agility Telecom no setor de tributos, onde identifiquei uma oportunidade de otimizar o atendimento ao cliente. Desenvolvi o Zentto Chatbot, uma solução inovadora que integrei ao sistema interno da empresa, resultando em melhorias significativas na eficiência e qualidade do atendimento. Este projeto me permitiu aprofundar meus conhecimentos em desenvolvimento de software e demonstrar o valor da tecnologia para a área de negócios.",
      technologies: [
        { icon: <SiTypescript size={24} />, name: "Typescript" },
        { icon: <SiMongodb size={24} />, name: "MongoDB" },
        { icon: <FaDocker size={24} />, name: "Docker" },
        { icon: <SiFastify size={24} />, name: "Fastify" },
      ],
    },

    {
      companyName: "Dosefy",
      startDate: new Date("2025-06-01"),
      endDate: new Date("2025-06-25"),
      role: "Backend Developer",
      description:
        "Como Desenvolvedor Backend na Dosemed, fui responsável pela criação de um ecossistema completo para o gerenciamento de medicamentos, incluindo o controle de fórmulas, dosagens, e alertas de uso. Além disso, liderei a integração de sistemas de pagamento, implementando com sucesso o Stripe para transações com cartão de crédito e o AbacatePay para pagamentos via Pix, garantindo uma experiência de compra segura e eficiente para os usuários.",
      technologies: [
        { icon: <SiTypescript size={24} />, name: "Typescript" },
        { icon: <SiMongodb size={24} />, name: "MongoDB" },
        { icon: <FaDocker size={24} />, name: "Docker" },
        { icon: <SiFastify size={24} />, name: "Fastify" },
      ],
    },
    {
      companyName: "Utility",
      startDate: new Date("2025-07-15"),
      endDate: null,
      role: "Freelance",
      description:
        "Como desenvolvedor freelancer, estou criando o Utility, um aplicativo de fretes que conecta motoristas a clientes de forma rápida e segura. A plataforma oferece recursos como cálculo de rotas, estimativa de custos, e acompanhamento em tempo real, otimizando a logística e a comunicação entre as partes. O projeto está sendo desenvolvido com foco em usabilidade e desempenho, proporcionando uma solução completa e confiável para o mercado de fretes.",
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
            <p className="text-sm leading-6">{xp.description}</p>
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
