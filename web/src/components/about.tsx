import {
  Contact,
  GitBranchPlus,
  Github,
  GithubIcon,
  Icon,
  Mail,
  MapIcon,
  Phone,
} from "lucide-react";
import { CommitChart } from "./commit-chart";
import { Button } from "./ui/button";

export const AboutSection = () => {
  return (
    <div className="w-full h-full flex gap-10">
      <div className="w-full flex-1 flex flex-col">
        <div className="w-full h-full flex flex-col pt-8 gap-4">
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
          <div className="flex-1 flex justify-center items-center gap-2 h-[150px] rounded-md border-2 border-zinc-800">
            <span className="font-bold text-2xl green-gradient bg-clip-text text-transparent">
              +49
            </span>{" "}
            projects
          </div>
          <div className="flex-1 flex flex-col gap-2 h-[150px] rounded-md border-2 border-zinc-800">
            <div className="flex-1">
              <div className=" p-2 bg-green-500 rounded-md">
                <MapIcon />
              </div>
            </div>
            <p className="flex-1 text-zinc-400 text-xs leading-4 w-2/3">
              Building high-performance websites with clean code and strong SEO
              fundamentals.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};
