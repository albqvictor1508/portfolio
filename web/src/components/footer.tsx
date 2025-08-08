import { FaInstagram, FaWhatsapp } from "react-icons/fa";
import { FiGithub } from "react-icons/fi";
import { SlSocialLinkedin } from "react-icons/sl";
import { useLanguage } from "../context/LanguageContext";

export const Footer = () => {
  const { t } = useLanguage();
  return (
    <div
      id="footer"
      className="border-t-2 border-t-zinc-900 w-full flex justify-center "
    >
      <div className="w-2/3 h-[200px] flex p-8 justify-between">
        <div className="w-full flex flex-col  gap-6">
          <div className="flex items-center justify-around">
            <span className="text-sm font-light text-zinc-200">Home</span>
            <span className="text-sm font-light text-zinc-200">About</span>
            <span className="text-sm font-light text-zinc-200">Contact</span>
            <span className="text-sm font-light text-zinc-200">Projects</span>
            <span className="text-sm font-light text-zinc-200">My journey</span>
          </div>
          <div className="w-2/3 mx-auto flex items-center justify-around">
            <span className="text-sm font-light ">
              <FiGithub size={16} />
            </span>
            <span className="text-sm font-light ">
              <FaInstagram size={16} />
            </span>
            <span className="text-sm font-light ">
              <SlSocialLinkedin size={16} />
            </span>
            <span className="text-sm font-light ">
              <FaWhatsapp size={16} />
            </span>
          </div>
          <p className="flex justify-center text-xs font-light text-zinc-600">
            © 2025 Victor Albuquerque. All rights reserved.
          </p>
        </div>
      </div>
    </div>
  );
};
