import { Code, Contact, GithubIcon, Layers } from "lucide-react";
import { useLanguage } from "../context/LanguageContext";
import { CommitChart } from "./commit-chart";
import { Button } from "./ui/button";

export const AboutSection = () => {
  const { t } = useLanguage();

  return (
    <div className="w-full h-full flex gap-10">
      <div className="w-full flex-1 flex flex-col">
        <div className="w-full h-full flex flex-col pt-6 gap-4">
          <h2 className="font-semibold text-2xl">{t("about_section.title")}</h2>
          <p className="text-sm font-light w-full leading-6 text-zinc-400">
            {t("about_section.description_p1")}
          </p>
          <p className="text-sm w-full leading-6 text-zinc-400">
            {t("about_section.description_p2")}
          </p>

          <p className="text-sm w-full leading-6 text-zinc-400">
            {t("about_section.description_p3")}
          </p>
          <div className="flex gap-2">
            <Button
              className="flex gap-2 items-center cursor-pointer"
              variant={"default"}
            >
              <span>
                <GithubIcon />
              </span>
              {t("about_section.view_on_github")}
            </Button>
            <Button
              className="flex gap-2 items-center cursor-pointer"
              variant={"secondary"}
            >
              <span>
                <Contact />
              </span>
              {t("about_section.contact_me")}
            </Button>
          </div>
        </div>
      </div>
      <div className="w-full flex flex-col gap-4 flex-1 rounded-md border-2 border-r-zinc-800 p-4">
        <div className="flex-1 w-full">
          <CommitChart />
        </div>
        <Button variant={"outline"}>{t("about_section.download_cv")}</Button>
        <div className="flex-1 w-full flex gap-6">
          <div className="flex-1 flex flex-col p-4 gap-2 h-[150px] rounded-md border-2 border-zinc-800">
            <div className="flex flex-col gap-2">
              <Layers size={24} className="text-green-400" />
              <p className="text-xs font-semibold">
                {t("about_section.solution_architecture_title")}
              </p>
            </div>
            <p className=" text-zinc-400 text-[10px] leading-4 w-full">
              {t("about_section.solution_architecture_description")}
            </p>
          </div>
          <div className="flex-1 flex flex-col p-4 gap-2 h-[150px] rounded-md border-2 border-zinc-800">
            <div className="flex flex-col gap-2">
              <Code size={24} className="text-green-400" />
              <p className="text-xs font-semibold">
                {t("about_section.full_stack_engineering_title")}
              </p>
            </div>
            <p className=" text-zinc-400 text-[10px] leading-4 w-full">
              {t("about_section.full_stack_engineering_description")}
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};
