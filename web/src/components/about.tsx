import { Button } from "./ui/button";

export const AboutSection = () => {
  return (
    <div className="w-full h-full flex gap-10">
      <div className="w-full flex-1 flex flex-col gap-4">
        <div className="w-full flex flex-col gap-2.5">
          <h2 className="font-semibold text-2xl">About me.</h2>
          <p className="text-sm text-justify">
            Lorem Ipsum is simply dummy text of the printing and typesetting and
            scrambled it to make a type specimen book. It has survived not only
            five centuries.
          </p>
          <p className="text-sm text-justify">
            with desktop publishing software like Aldus PageMaker including
            versions of Lorem Ipsum.
          </p>
        </div>
        <div className="flex gap-2">
          <Button className="cursor-pointer" variant={"default"}>
            View on Github
          </Button>
          <Button className="cursor-pointer" variant={"secondary"}>
            Contact me
          </Button>
        </div>
      </div>
      <div className="w-full flex flex-col gap-4 flex-1 rounded-md border-2 border-r-zinc-800 p-4">
        <div className="flex-1">
          <h2>salve</h2>
        </div>
        <div className="flex-1 w-full h-[150px] flex gap-10">
          <div className="flex-1  rounded-md border-2 border-zinc-800"></div>
          <div className="flex-1  rounded-md border-2 border-zinc-800"></div>
        </div>
      </div>
    </div>
  );
};
