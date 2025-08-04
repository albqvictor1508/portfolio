import { Bar, BarChart } from "recharts";
import { ChartContainer } from "./ui/chart";

export const CommitChart = () => {
  return (
    <div>
      <ChartContainer>
        <BarChart data={[{ salve: "salve" }]}>
          <Bar dataKey={"value"}></Bar>
        </BarChart>
      </ChartContainer>
    </div>
  );
};
