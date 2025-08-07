import { LucideMail, MapIcon } from "lucide-react";
import { useLanguage } from "../context/LanguageContext";

export const HeroSection = () => {
	const { t } = useLanguage();
	return (
		<div id="hero" className="w-full h-full flex flex-col gap-4">
			<div className="w-full h-[600px] flex">
				<div className="flex flex-1 flex-col gap-12">
					<h1 className="flex flex-col gap-8 text-5xl font-bold font-montserrat">
						<span className="text-6xl ">{t("hero_section.hey_im")}</span>
						<span className="text-green-400 text-6xl">Victor Arruda</span>
					</h1>
					<ul className="flex flex-col gap-4">
						<li className="font-semibold text-2xl italic text-green-400">
							{"{Software Engineer}"}
						</li>
						<li className="text-sm leading-6 text-zinc-400">
							{t("hero_section.description")}
						</li>
					</ul>
				</div>
				{/*
				<div className="flex flex-1 justify-center items-center w-max rounded-md border-2 border-zinc-800">
					salveeeeee hero
				</div>
        */}
			</div>
			<div className="w-full flex justify-between items-center">
				<div className="flex items-center gap-2">
					<span className="text-green-200">
						<MapIcon size={18} />
					</span>
					<span className="text-sm text-green-200 font-semibold">
						{t("hero_section.location")}
					</span>
				</div>
				<div className="flex items-center gap-2">
					<span className="text-green-200">
						<LucideMail size={18} />
					</span>
					<span className="text-sm text-green-200 font-semibold">
						albq.victor@gmail.com
					</span>
				</div>
			</div>
			<div className="w-full flex justify-between">
				<span className="text-sm text-muted">salve</span>
				<span className="text-sm text-muted">salve</span>
				<span className="text-sm text-muted">salve</span>
				<span className="text-sm text-muted">salve</span>
			</div>
		</div>
	);
};
