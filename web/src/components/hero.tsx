import { LucideMail, MapIcon } from "lucide-react";

export const HeroSection = () => {
  return (
    <div id="hero" className="w-full h-full flex flex-col gap-4">
      <div className="w-full h-[600px] flex gap-6">
        <div className="flex flex-1 flex-col justify-center gap-4">
          <h1 className="flex flex-col gap-6 text-5xl font-bold">
            <span className="text-5xl">Hey, i'm</span>
            <span className="text-green-400 text-5xl">Victor Arruda</span>
          </h1>
          <ul className="flex flex-col gap-4">
            <li className="font-semibold text-xl italic text-green-400">
              {"{Software Engineer}"}
            </li>
            <li className="text-sm leading-6 text-zinc-400">
              It is a long established fact that a reader will be distracted by
              the readable content of a page when looking at its layout. The
              purpose (injected humour and the like).
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
