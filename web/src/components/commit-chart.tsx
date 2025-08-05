"use client";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { ChartContainer, ChartTooltip, ChartTooltipContent } from "./ui/chart";
import { Bar, BarChart, CartesianGrid, XAxis } from "recharts";

export const CommitChart = () => {
  const today = new Date();
  const data = Array.from({ length: 30 }, (_, i) => {
    const date = new Date();
    date.setDate(today.getDate() - (29 - i));
    return {
      date: date.toLocaleDateString("en-US", {
        month: "short",
        day: "numeric",
      }),
      commits: Math.floor(Math.random() * 15),
    };
  });

  const chartConfig = {
    commits: {
      label: "Commits",
      color: "hsl(142, 71%, 45%)",
    },
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>Commit History</CardTitle>
        <CardDescription>Last 30 days</CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="h-[150px] w-full">
          <BarChart accessibilityLayer data={data}>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="date"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.split(" ")[1]}
              interval={4}
            />
            <ChartTooltip content={<ChartTooltipContent />} />
            <Bar dataKey="commits" fill="var(--color-commits)" radius={4} barSize={10} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
};
