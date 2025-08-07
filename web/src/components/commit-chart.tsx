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

  const chartConfig = {
    commits: {
      label: "Commits",
      color: "hsl(142, 71%, 45%)",
    },
  };

  useEffect(() => {
    console.log("useEffect in CommitChart is running.");
    const fetchCommits = async () => {
      try {
        setLoading(true);
        const today = new Date();
        const todayUTC = new Date(
          Date.UTC(
            today.getUTCFullYear(),
            today.getUTCMonth(),
            today.getUTCDate(),
          ),
        );
        const sinceUTC = new Date(todayUTC);
        sinceUTC.setUTCDate(todayUTC.getUTCDate() - 89);
        const sinceDate = sinceUTC.toISOString().split("T")[0];

        console.log("Attempting to fetch from API...");
        let response;
        try {
          response = await fetch(
            `http://localhost:8080/commits?since=${sinceDate}`,
          );
        } catch (fetchError) {
          console.error("Error during fetch:", fetchError);
          throw fetchError;
        }

        let commitsByDay: { [key: string]: number };
        try {
          commitsByDay = await response.json();
          console.log("API Response (commitsByDay):", commitsByDay);
        } catch (jsonError) {
          console.error("Error parsing JSON response:", jsonError);
          throw jsonError;
        }

        try {
          const last90Days = [];
          for (let i = 0; i < 90; i++) {
            const pastDate = new Date(todayUTC);
            pastDate.setUTCDate(todayUTC.getUTCDate() - i);
            last90Days.push(pastDate);
          }
          last90Days.reverse();

          const finalChartData = last90Days.map((date) => {
            const isoDate = date.toISOString().split("T")[0];
            const commitCount = commitsByDay[isoDate] || 0;
            return {
              date: date.toLocaleDateString("en-US", {
                month: "short",
                day: "numeric",
                timeZone: "UTC",
              }),
              commits: commitCount,
            };
          });

          console.log(
            "Processed Data for Chart (finalChartData):",
            finalChartData,
          );
          setData(finalChartData);
        } catch (processingError) {
          console.error("Error during data processing:", processingError);
          throw processingError;
        }
      } catch (overallError) {
        console.error("Overall error in fetchCommits:", overallError);
      } finally {
        setLoading(false);
      }
    };

    fetchCommits();
  }, []);

  console.log("Data in state (for render):", data);
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
