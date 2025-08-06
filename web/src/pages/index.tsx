import { AboutSection } from "@/components/about";
import { Footer } from "@/components/footer";
import { HeroSection } from "@/components/hero";
import { Menu } from "@/components/menu";
import { ProjectSection } from "@/components/project/section";
import { StackSection } from "@/components/stack";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="w-[1100px] h-full flex flex-col gap-12 m-auto justify-center items-center">
      <Menu />
      <HeroSection />
      <AboutSection />
      <ProjectSection />
      <StackSection />
      <Footer isLight={true} />
    </div>
  );
}
