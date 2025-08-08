import { useLanguage } from "../../context/LanguageContext";
import { Project } from "./project";

const projects = [
  {
    name: "Zentto Chatbot",
    startDate: "04 Jun, 2025",
    endDate: "Present",
    description:
      "Zentto Chatbot é um chatbot integrado ao IXCsoft que criei com o objetivo de atender todos os provedores de internet da minha região que trabalham com essa ferramenta",
    imageUrl: "https://flameshot.org/media/images/preview.png",
    repoUrl: "",
    demoUrl: "",
  },
  {
    name: "Plush",
    startDate: "01 Mar, 2025",
    endDate: "Present",
    description:
      "Uma plataforma de chat em tempo real inspirada no whatsapp, com funcionalidades completas como: listagem de chats, envio, atualização e deleção de imagens, criação de grupos, entre outras funcionalidades",
    imageUrl: "",
    repoUrl: "",
    demoUrl: "https://getsharex.com/",
  },
];

export const ProjectSection = () => {
  const { t } = useLanguage();
  return (
    <div className="w-full h-full flex flex-col gap-8">
      <div className="w-full h-full flex flex-col gap-2">
        <h2 className="text-2xl font-semibold">{t("project_section.title")}</h2>
        <p className="w-2/3 text-justify text-sm">
          {t("project_section.description")}
        </p>
      </div>

      <div className="flex flex-col gap-12">
        {projects.map((project, index) => (
          <Project key={index.toString()} {...project} />
        ))}
      </div>
    </div>
  );
};
