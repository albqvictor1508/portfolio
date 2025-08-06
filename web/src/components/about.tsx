import { Code, Contact, GithubIcon, Layers } from "lucide-react";
import { CommitChart } from "./commit-chart";
import { Button } from "./ui/button";

export const AboutSection = () => {
  return (
    <div className="w-full h-full flex gap-10">
      <div className="w-full flex-1 flex flex-col">
        <div className="w-full h-full flex flex-col pt-6 gap-4">
          <h2 className="font-semibold text-2xl">About me.</h2>
          <p className="text-sm font-light w-full leading-6 text-zinc-400">
            Lorem Ipsum is simply dummy text of the printing and typesetting and
            scrambled it to make a type specimen book. It has survived not only
            five centuries.
          </p>
          <p className="text-sm w-full leading-6 text-zinc-400">
            Lorem Ipsum is simply dummy text of the printing and typesetting and
            scrambled it to make a type specimen book. It has survived not only
            five centuries.
          </p>

          <p className="text-sm w-full leading-6 text-zinc-400">
            with desktop publishing software like Aldus PageMaker including
            versions of Lorem Ipsum.
          </p>
          <div className="flex gap-2">
            <Button
              className="flex gap-2 items-center cursor-pointer"
              variant={"default"}
            >
              <span>
                <GithubIcon />
              </span>
              View on Github
            </Button>
            <Button
              className="flex gap-2 items-center cursor-pointer"
              variant={"secondary"}
            >
              <span>
                <Contact />
              </span>
              Contact me
            </Button>
          </div>
        </div>
      </div>
      <div className="w-full flex flex-col gap-4 flex-1 rounded-md border-2 border-r-zinc-800 p-4">
        <div className="flex-1 w-full">
          <CommitChart />
        </div>
        <Button variant={"outline"}>Download CV</Button>
        <div className="flex-1 w-full flex gap-6">
          <div className="flex-1 flex flex-col p-4 gap-2 h-[150px] rounded-md border-2 border-zinc-800">
            <div className="flex flex-col gap-2">
              <Layers size={24} className="text-green-400" />
              <p className="text-xs font-semibold">Solution Architecture</p>
            </div>
            <p className=" text-zinc-400 text-[10px] leading-4 w-full">
              Designing robust, scalable, and maintainable software systems to
              meet complex business requirements.
            </p>
          </div>
          <div className="flex-1 flex flex-col p-4 gap-2 h-[150px] rounded-md border-2 border-zinc-800">
            <div className="flex flex-col gap-2">
              <Code size={24} className="text-green-400" />
              <p className="text-xs font-semibold">Full-Stack Engineering</p>
            </div>
            <p className=" text-zinc-400 text-[10px] leading-4 w-full">
              Building complete, end-to-end applications with a focus on clean
              code, performance, and user experience.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};
