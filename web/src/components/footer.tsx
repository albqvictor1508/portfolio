import { Link } from "@tanstack/react-router";
import { FaInstagram, FaWhatsapp } from "react-icons/fa";
import { FiGithub } from "react-icons/fi";
import { SiThreads } from "react-icons/si";
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
            <a
              href="https://github.com/albqvictor1508"
              target="_blank"
              rel="noopener"
              className="text-sm font-light "
            >
              <FiGithub size={18} />
            </a>
            <a
              href="https://www.threads.com/@albqvxc"
              target="_blank"
              rel="noopener"
              className="text-sm font-light "
            >
              <SiThreads size={18} />
            </a>
            <a
              href="https://www.linkedin.com/in/albqvictor1508"
              target="_blank"
              rel="noopener"
              className="text-sm font-light "
            >
              <SlSocialLinkedin size={18} />
            </a>
            <a
              href="https://wa.me/+5583991303948"
              className="text-sm font-light "
              target="_blank"
              rel="noopener"
            >
              <FaWhatsapp size={18} />
            </a>
          </div>
          <p className="flex justify-center text-xs font-light text-zinc-600">
            © 2025 Victor Albuquerque. All rights reserved.
          </p>
        </div>
      </div>
    </div>
  );
};
