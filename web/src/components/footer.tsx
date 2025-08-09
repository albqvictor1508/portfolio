import { Link } from "@tanstack/react-router";
import { FaInstagram, FaWhatsapp } from "react-icons/fa";
import { FiGithub } from "react-icons/fi";
import { SlSocialLinkedin } from "react-icons/sl";

export const Footer = () => {
  return (
    <div
      id="footer"
      className="border-t-2 border-t-zinc-900 w-full flex justify-center "
    >
      <div className="w-2/3 h-[200px] flex p-8 justify-between">
        <div className="w-full flex flex-col gap-8">
          <div className="flex items-center justify-around">
            <Link to="." className="text-sm font-light text-zinc-200">
              Home
            </Link>
            <Link to="." className="text-sm font-light text-zinc-200">
              About
            </Link>
            <Link to="." className="text-sm font-light text-zinc-200">
              Contact
            </Link>
            <Link to="." className="text-sm font-light text-zinc-200">
              Projects
            </Link>
            <Link to="." className="text-sm font-light text-zinc-200">
              My journey
            </Link>
          </div>
          <div className="w-2/3 mx-auto flex items-center justify-around">
            <Link to="." className="text-sm font-light ">
              <FiGithub size={18} />
            </Link>
            <Link to="." className="text-sm font-light ">
              <FaInstagram size={18} />
            </Link>
            <Link to="." className="text-sm font-light ">
              <SlSocialLinkedin size={18} />
            </Link>
            <Link to="." className="text-sm font-light ">
              <FaWhatsapp size={18} />
            </Link>
          </div>
          <p className="flex justify-center text-xs font-light text-zinc-600">
            © 2025 Victor Albuquerque. All rights reserved.
          </p>
        </div>
      </div>
    </div>
  );
};
