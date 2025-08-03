import { HeroSection } from "@/components/hero-section";
import { Menu } from "@/components/menu";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
	component: RouteComponent,
});

function RouteComponent() {
	return (
		<div className="w-[1000px] h-full flex flex-col gap-12 m-auto justify-center items-center">
			<Menu />
			<HeroSection />
		</div>
	);
}
