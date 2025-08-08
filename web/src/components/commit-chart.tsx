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
				const response = await fetch(
					`http://localhost:3333/commits?since=2025-06-01`,
				);
				const data = await response.json();
				console.log(data);
				setData(data);
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
