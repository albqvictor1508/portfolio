"use client";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "./ui/tooltip";

export const CommitChart = () => {
  const today = new Date();
  const currentDayOfWeek = today.getDay();
  const days = Array.from({ length: 365 }, (_, i) => {
    const date = new Date(today);
    date.setDate(today.getDate() - 365 + i);
    return {
      date,
      commits: Math.floor(Math.random() * 10),
    };
  });

  const getTooltipText = (day: { date: Date; commits: number }) => {
    const date = day.date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "long",
      day: "numeric",
    });
    return `${day.commits} commits on ${date}`;
  };

  const getColor = (commits: number) => {
    if (commits === 0) return "bg-zinc-800";
    if (commits < 2) return "bg-green-800";
    if (commits < 5) return "bg-green-600";
    if (commits < 8) return "bg-green-400";
    return "bg-green-200";
  };

  return (
    <div className="flex flex-col gap-2 w-full">
      <div className="grid grid-cols-53 grid-rows-7 gap-1 w-full">
        {Array.from({ length: currentDayOfWeek }).map((_, i) => (
          <div key={`empty-${i}`} className="w-3 h-3 rounded-sm" />
        ))}
        {days.map((day, i) => (
          <TooltipProvider key={i}>
            <Tooltip>
              <TooltipTrigger asChild>
                <div
                  className={`w-3 h-3 rounded-sm ${getColor(day.commits)}`}
                />
              </TooltipTrigger>
              <TooltipContent>
                <p>{getTooltipText(day)}</p>
              </TooltipContent>
            </Tooltip>
          </TooltipProvider>
        ))}
      </div>
    </div>
  );
};

