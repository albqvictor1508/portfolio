import { useLanguage } from "../../context/LanguageContext";
import { Project } from "./project";

const projects = [
  {
    name: "Zentto Chatbot",
    startDate: new Date("2025-06-04"),
    endDate: null,
    description:
      "Desenvolvimento de um chatbot inteligente e integrado ao sistema IXCsoft, projetado para otimizar o atendimento ao cliente de provedores de internet na região. A ferramenta automatiza interações, agilizando o suporte e melhorando a experiência do usuário final.",
    imageUrl: "https://flameshot.org/media/images/preview.png",
    repoUrl: "",
    demoUrl: "",
  },
  {
    name: "Plush",
    startDate: new Date("2025-03-01"),
    endDate: null,
    description:
      "Plataforma de comunicação em tempo real, 'Plush', inspirada no WhatsApp, oferecendo uma experiência de chat completa. Inclui funcionalidades robustas como listagem e gerenciamento de conversas, envio e manipulação de mídias (imagens), e criação de grupos, proporcionando uma interação fluida e rica.",
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
