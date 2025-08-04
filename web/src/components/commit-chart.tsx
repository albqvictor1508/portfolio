'use client';
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from './ui/tooltip';

export const CommitChart = () => {
  const today = new Date();
  const daysInPast = 90;
  const days = Array.from({ length: daysInPast }, (_, i) => {
    const date = new Date();
    date.setDate(today.getDate() - (daysInPast - 1 - i));
    return {
      date,
      commits: Math.floor(Math.random() * 10),
    };
  });

  const getTooltipText = (day: { date: Date; commits: number } | undefined) => {
    if (!day) return '';
    const date = day.date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });
    return `${day.commits} commits on ${date}`;
  };

  const getColor = (commits: number) => {
    if (commits === 0) return 'fill-zinc-800';
    if (commits < 2) return 'fill-green-400';
    if (commits < 5) return 'fill-green-600';
    if (commits < 8) return 'fill-green-800';
    return 'fill-green-950';
  };

  const squareSize = 40;
  const gap = 10;
  const firstDayOfWeek = days[0]?.date.getDay() ?? 0;
  const weekCount = Math.ceil((daysInPast + firstDayOfWeek) / 7);

  const svgWidth = weekCount * (squareSize + gap) - gap;
  const svgHeight = 7 * (squareSize + gap) - gap;

  return (
    <div className="w-full">
      <svg
        width="100%"
        viewBox={`0 0 ${svgWidth} ${svgHeight}`}
        preserveAspectRatio="xMidYMid meet"
      >
        <g>
          {days.map((day, index) => {
            const totalIndex = index + firstDayOfWeek;
            const weekIndex = Math.floor(totalIndex / 7);
            const dayIndex = totalIndex % 7;

            const x = weekIndex * (squareSize + gap);
            const y = dayIndex * (squareSize + gap);

            return (
              <TooltipProvider key={index}>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <rect
                      x={x}
                      y={y}
                      width={squareSize}
                      height={squareSize}
                      className={getColor(day.commits)}
                      rx="0"
                      ry="0"
                    />
                  </TooltipTrigger>
                  <TooltipContent>
                    <p>{getTooltipText(day)}</p>
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>
            );
          })}
        </g>
      </svg>
    </div>
  );
};