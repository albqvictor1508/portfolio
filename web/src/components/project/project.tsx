import { Button } from "../ui/button";

export const Project = () => {
  return (
    <div className="w-full h-full flex flex-col gap-2">
      <h2 className="font-semibold text-2xl">Zentto</h2>
      <span className="text-sm text-zinc-500">25 Jun, 2025 - Presente</span>
      <p>lorem</p>

      <div className="w-full h-68 rounded-md bg-zinc-800"></div>
      <div className="flex justify-end items-center gap-4">
        <Button variant={"secondary"}>View Repo</Button>
        <Button variant={"default"}>View Demo</Button>
      </div>
    </div>
  );
};
