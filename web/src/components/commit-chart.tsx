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
        setLoading(true);
        // Determine the date range for the API call
        const today = new Date();
        const todayUTC = new Date(Date.UTC(today.getUTCFullYear(), today.getUTCMonth(), today.getUTCDate()));
        const sinceUTC = new Date(todayUTC);
        sinceUTC.setUTCDate(todayUTC.getUTCDate() - 89);
        const sinceDate = sinceUTC.toISOString().split("T")[0];

        // Fetch commit data
        const response = await fetch(
          `http://localhost:8080/commits?since=${sinceDate}`,
        );
        const commitsByDay: { [key:string]: number } = await response.json();

        // Generate a list of the last 90 days in UTC
        const last90Days = [];
        for (let i = 0; i < 90; i++) {
          const pastDate = new Date(todayUTC);
          pastDate.setUTCDate(todayUTC.getUTCDate() - i);
          last90Days.push(pastDate);
        }
        last90Days.reverse(); // Put in chronological order

        // Map the generated dates to the chart data format, populating with fetched commit counts
        const finalChartData = last90Days.map(date => {
            const isoDate = date.toISOString().split('T')[0]; // YYYY-MM-DD format for lookup
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

        setData(finalChartData);
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
