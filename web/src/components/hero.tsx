import { LucideMail, MapIcon } from "lucide-react";
import { FaLinkedinIn, FaWhatsapp } from "react-icons/fa";
import { SiGithub, SiThreads } from "react-icons/si";
import { useLanguage } from "../context/LanguageContext";
import TypingEffect from "./typing-effect";

export const HeroSection = () => {
  const { t } = useLanguage();
  return (
    <div id="hero" className="w-full h-full flex flex-col">
      <div className="w-full h-full flex">
        <div className="flex flex-col gap-12">
          <h1 className="flex flex-col gap-7 text-5xl font-extrabold">
            <span className="text-7xl font-montserrat text-zinc-200">
              {t("hero_section.hey_im")}
            </span>
            <span className="text-green-gradient text-7xl font-montserrat">
              <TypingEffect text="Victor Albuquerque" />
            </span>
          </h1>
          <ul className="flex flex-col gap-6">
            <li className="font-medium text-3xl italic  text-green-200">
              {"{Software Engineer}"}
            </li>
            <li className="leading-6 text-zinc-400">
              {t("hero_section.description")}
            </li>
          </ul>

          <div className="w-full flex flex-col gap-10">
            <div className="flex justify-between items-center">
              <div className="flex items-center gap-2">
                <span className="text-green-200">
                  <MapIcon size={18} />
                </span>
                <span className="text-base text-green-200 font-semibold">
                  {t("hero_section.location")}
                </span>
              </div>
              <div className="flex items-center gap-2">
                <span className="text-green-200">
                  <LucideMail size={18} />
                </span>
                <span className="text-base text-green-200 font-semibold">
                  albq.victor@gmail.com
                </span>
              </div>
            </div>

            <div className="w-full flex justify-between">
              <p className="flex items-center gap-2">
                <span className="flex justify-center items-center p-2 bg-green-gradient rounded-md">
                  <SiGithub className="font-bold text-zinc-100" size={20} />
                </span>
                <span className="">Github</span>
              </p>

              <p className="flex items-center gap-2">
                <span className="p-2 bg-green-gradient rounded-md">
                  <FaLinkedinIn className="font-bold text-zinc-100" size={20} />
                </span>
                <span className="">LinkedIn</span>
              </p>

              <p className="flex items-center gap-2">
                <span className="p-2 bg-green-gradient rounded-md">
                  <SiThreads className="font-bold text-zinc-100" size={20} />
                </span>
                <span className="">Threads</span>
              </p>

              <p className="flex items-center gap-2">
                <span className="p-2 bg-green-gradient rounded-md">
                  <FaWhatsapp className="font-bold text-zinc-100" size={20} />
                </span>
                <span className="">Whatsapp</span>
              </p>
            </div>
          </div>
        </div>
        {/*
				<div className="flex flex-1 justify-center items-center w-max rounded-md border-2 border-zinc-800">
					salveeeeee hero
				</div>
        */}
      </div>
    </div>
  );
};
