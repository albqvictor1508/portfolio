import { useLanguage } from "../../context/LanguageContext";
import { Project } from "./project";

const projects = [
	{
		name: "Zentto Chatbot",
		startDate: "01 Jan, 2017",
		endDate: "Present",
		description:
			"Zentto Chatbot é um chatbot integrado ao IXCsoft que criei com o objetivo de atender todos os provedores de internet da minha região que trabalham com essa ferramenta",
		imageUrl: "https://flameshot.org/media/images/preview.png",
		repoUrl: "",
		demoUrl: "",
	},
	{
		name: "Salve",
		startDate: "01 Jan, 2007",
		endDate: "Present",
		description:
			"ShareX is a free and open source program that lets you capture or record any area of your screen and share it with a single press of a key. It also allows uploading images, text or other types of files to over 80 supported destinations you can choose from.",
		imageUrl: "https://getsharex.com/img/hero_image.png",
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
