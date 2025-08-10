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
		const fetchCommits = async () => {
			try {
				setLoading(true);
				const today = new Date();
				const sinceDate = new Date(today);
				sinceDate.setDate(today.getDate() - 89);
				const formattedSinceDate = sinceDate.toISOString().split("T")[0]; // YYYY-MM-DD

				const response = await fetch(
					`${import.meta.env.VITE_API_BASE_URL}/commits?since=${formattedSinceDate}`,
				);
				const commitsByDay: { [key: string]: number } = await response.json();

				const last90DaysData: ChartData[] = [];
				for (let i = 0; i < 90; i++) {
					const date = new Date(today);
					date.setDate(today.getDate() - (89 - i)); // Iterate from 89 days ago to today
					const isoDate = date.toISOString().split("T")[0]; // YYYY-MM-DD for lookup
					last90DaysData.push({
						date: date.toLocaleDateString("en-US", {
							month: "short",
							day: "numeric",
						}),
						commits: commitsByDay[isoDate] || 0,
					});
				}

				setData(last90DaysData);
			} catch (error) {
				console.error("Failed to fetch commits:", error);
			} finally {
				setLoading(false);
			}
		};

		fetchCommits();
	}, []);

	return (
		<Card>
			<CardHeader>
				<CardTitle>Commit History</CardTitle>
				<CardDescription>Last 90 days</CardDescription>
			</CardHeader>
			<CardContent>
				{loading ? (
					<Skeleton className="h-[150px] w-full" />
				) : data.length === 0 ? (
					<div className="h-[150px] w-full flex items-center justify-center text-zinc-500">
						No commit data available.
					</div>
				) : (
					<ChartContainer config={chartConfig} className="h-[150px] w-full">
						<BarChart accessibilityLayer data={data}>
							<CartesianGrid vertical={false} />
							<XAxis
								dataKey="date"
								tickLine={false}
								tickMargin={2}
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
								barSize={20}
							/>
						</BarChart>
					</ChartContainer>
				)}
			</CardContent>
		</Card>
	);
};
