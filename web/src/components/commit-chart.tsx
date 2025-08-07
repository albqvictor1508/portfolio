"use client";
import { Bar, BarChart, CartesianGrid, XAxis } from "recharts";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { ChartContainer, ChartTooltip, ChartTooltipContent } from "./ui/chart";
import { useEffect, useState } from "react";
import { Skeleton } from "./ui/skeleton";

type ChartData = {
  date: string;
  commits: number;
};

export const CommitChart = () => {
  const [data, setData] = useState<ChartData[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCommits = async () => {
      try {
        const today = new Date();
        const since = new Date();
        since.setDate(today.getDate() - 89);
        const sinceDate = since.toISOString().split("T")[0];

        const response = await fetch(
          `http://localhost:8080/commits?since=${sinceDate}`,
        );
        const commitsByDay: { [key: string]: number } = await response.json();

        const allDaysData = Array.from({ length: 90 }, (_, i) => {
          const date = new Date();
          date.setUTCDate(date.getUTCDate() - i);
          const isoDate = date.toISOString().split("T")[0];
          return {
            date: isoDate,
            commits: 0,
          };
        });

        for (const day of allDaysData) {
          if (commitsByDay[day.date]) {
            day.commits = commitsByDay[day.date];
          }
        }

        const formattedData = allDaysData
          .map((d) => ({
            ...d,
            date: new Date(`${d.date}T00:00:00Z`).toLocaleDateString("en-US", {
              month: "short",
              day: "numeric",
              timeZone: "UTC",
            }),
          }))
          .reverse();

        setData(formattedData);
      } catch (error) {
        console.error("Failed to fetch commits:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchCommits();
  }, []);

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
        <CardDescription>Last 90 days</CardDescription>
      </CardHeader>
      <CardContent>
        {loading ? (
          <Skeleton className="h-[150px] w-full" />
        ) : (
          <ChartContainer config={chartConfig} className="h-[150px] w-full">
            <BarChart accessibilityLayer data={data}>
              <CartesianGrid vertical={false} />
              <XAxis
                dataKey="date"
                tickLine={false}
                tickMargin={10}
                axisLine={false}
                tickFormatter={(value) => value.split(" ")[1]}
                interval={10}
              />
              <ChartTooltip
                cursor={false}
                content={<ChartTooltipContent hideLabel />}
              />
              <Bar
                dataKey="commits"
                fill="var(--color-commits)"
                radius={4}
                barSize={10}
              />
            </BarChart>
          </ChartContainer>
        )}
      </CardContent>
    </Card>
  );
};
