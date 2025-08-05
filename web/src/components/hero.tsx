import { LucideMail, MapIcon } from "lucide-react";

export const HeroSection = () => {
  return (
    <div className="w-full h-full flex flex-col gap-4">
      <div className="w-full h-full flex gap-8">
        <div className="flex flex-1 flex-col gap-4">
          <h1 className="text-3xl font-bold">
            Hey, I'm <span className="text-green-400">Victor Arruda</span>
          </h1>
          <ul className="flex flex-col gap-2">
            <li>Back-end Engineer</li>
            <li className="text-sm text-justify">
              Founder{" "}
              <span className="font-bold text-green-400">Zentto Chatbot</span>,
              a software development chatbot dedicated to talk with customer .
            </li>
          </ul>
        </div>
        <div className="flex flex-1 justify-center items-center w-max rounded-md border-2 border-zinc-800">
          salveeeeee hero
        </div>
      </div>
      <div className="w-full flex justify-between items-center">
        <div className="flex items-center gap-2">
          <span className="text-green-200">
            <MapIcon size={18} />
          </span>
          <span className="text-sm text-green-200 font-semibold">
            João Pessoa, Brazil
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
